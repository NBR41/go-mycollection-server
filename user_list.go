package main

import (
	"net/http"
)

func listUsers(w http.ResponseWriter, r *http.Request) {
	m, err := NewModel(connString)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError)
		return
	}

	l, err := m.GetUserList()
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}
	writeResponse(w, r, l)
}
