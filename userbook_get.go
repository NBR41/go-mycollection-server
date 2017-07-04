package main

import (
	"net/http"
)

func getUserBook(w http.ResponseWriter, r *http.Request) {
	ids, err := getURLIDs(r, urlUserIDField, urlBookIDField)
	if err != nil {
		writeError(w, r, http.StatusBadRequest)
		return
	}

	m, err := NewModel(connString)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError)
		return
	}
	defer func() { _ = m.close() }()

	ub, err := m.GetUserBook(ids[0], ids[1])
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}
	if ub == nil {
		writeError(w, r, http.StatusNotFound)
		return
	}
	writeResponse(w, r, ub)
}
