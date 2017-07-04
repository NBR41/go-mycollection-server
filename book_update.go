package main

import (
	"net/http"
)

func updateBook(w http.ResponseWriter, r *http.Request) {
	var name string
	if name = r.PostForm.Get(formBookNameField); name == "" {
		writeErrorWithMessage(w, r, http.StatusBadRequest, "invalid book name")
		return
	}

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

	err = m.UpdateBook(ids[0], name)
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}
	b := Book{ID: ids[0], Name: name}
	b.initURL()
	writeResponse(w, r, b)
}
