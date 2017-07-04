package main

import (
	"net/http"
)

func deleteUser(w http.ResponseWriter, r *http.Request) {
	ids, err := getURLIDs(r, urlUserIDField)
	if err != nil {
		writeError(w, r, http.StatusBadRequest)
		return
	}

	m, err := newModel(connString)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError)
		return
	}
	defer func() { _ = m.close() }()

	err = m.deleteUser(ids[0])
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}

	u, err := m.getUserByID(ids[0])
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}
	u.initURL()
	writeResponse(w, r, u)
}
