package main

import (
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	m, err := newModel(connString)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError)
		return
	}

	ids, err := getURLIDs(r, urlUserIDField)
	if err != nil {
		writeError(w, r, http.StatusBadRequest)
		return
	}

	u, err := m.getUserByID(ids[0])
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}
	writeResponse(w, r, u)
}
