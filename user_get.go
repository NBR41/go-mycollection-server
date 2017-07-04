package main

import (
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	ids, err := getURLIDs(r, urlUserIDField)
	if err != nil {
		writeError(w, r, http.StatusBadRequest)
		return
	}

	m, err := NewModel(connString)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError)
		return
	}

	u, err := m.GetUserByID(ids[0])
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}
	if u == nil {
		writeError(w, r, http.StatusNotFound)
		return
	}
	writeResponse(w, r, u)
}
