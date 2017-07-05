package main

import (
	"net/http"
)

func listOwnerships(w http.ResponseWriter, r *http.Request) {
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

	l, err := m.GetOwnershipList(ids[0])
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}
	writeResponse(w, r, l)
}
