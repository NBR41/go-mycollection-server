package main

import (
	"net/http"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	var email, password, nickname string
	if email = r.PostForm.Get(formEmailField); email == "" {
		writeErrorWithMessage(w, r, http.StatusBadRequest, "invalid email")
		return
	}
	if password = r.PostForm.Get(formPasswordField); password == "" {
		writeErrorWithMessage(w, r, http.StatusBadRequest, "invalid password")
		return
	}
	if nickname = r.PostForm.Get(formNicknameField); nickname == "" {
		writeErrorWithMessage(w, r, http.StatusBadRequest, "invalid nickname")
		return
	}

	m, err := newModel(connString)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError)
		return
	}
	defer func() { _ = m.close() }()

	u, err := m.getUserByEmailOrNickname(email, nickname)
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}

	if u != nil {
		writeErrorWithMessage(w, r, http.StatusBadRequest, "user already exists")
		return
	}

	u, err = m.insertUser(nickname, email, password)
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}
	u.initURL()
	writeResponse(w, r, u)
}
