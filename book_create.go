package main

import (
	"net/http"
)

func createBook(w http.ResponseWriter, r *http.Request) {
	var name string
	if name = r.PostForm.Get(formBookNameField); name == "" {
		writeErrorWithMessage(w, r, http.StatusBadRequest, "invalid book name")
		return
	}

	m, err := NewModel(connString)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError)
		return
	}
	defer func() { _ = m.close() }()

	b, err := m.InsertBook(name)
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}
	writeResponse(w, r, b)
}
