package main

import (
	"net/http"
)

func deleteBook(w http.ResponseWriter, r *http.Request) {
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
	defer func() { _ = m.close() }()

	err = m.DeleteBook(ids[0])
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}
	writeResponse(w, r, nil)
}
