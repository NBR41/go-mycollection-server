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

	m, err := NewModel(connString)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError)
		return
	}
	defer func() { _ = m.close() }()

	u, err := m.GetUserByEmailOrNickname(email, nickname)
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}

	if u != nil {
		writeErrorWithMessage(w, r, http.StatusBadRequest, "user already exists")
		return
	}

	u, err = m.InsertUser(nickname, email, password)
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}
	writeResponse(w, r, u)
}
