package main

import (
	"net/http"
)

func updateUser(w http.ResponseWriter, r *http.Request) {
	var nickname string
	if nickname = r.PostForm.Get(formNicknameField); nickname == "" {
		writeErrorWithMessage(w, r, http.StatusBadRequest, "invalid nickname")
		return
	}

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

	err = m.updateUserNickname(ids[0], nickname)
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
