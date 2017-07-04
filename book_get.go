package main

import (
	"net/http"
)

func getBook(w http.ResponseWriter, r *http.Request) {
	ids, err := getURLIDs(r, urlBookIDField)
	if err != nil {
		writeError(w, r, http.StatusBadRequest)
		return
	}

	m, err := NewModel(connString)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError)
		return
	}

	b, err := m.GetBookByID(ids[0])
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}
	if b == nil {
		writeError(w, r, http.StatusNotFound)
		return
	}
	writeResponse(w, r, b)
}
