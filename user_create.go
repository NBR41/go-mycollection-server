package main

import (
	"net/http"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		writeError(w, r, http.StatusBadRequest)
		return
	}

	var email, password, nickname string
	if email = r.PostForm.Get("email"); email == "" {
		writeErrorWithMessage(w, r, http.StatusBadRequest, "invalid email")
		return
	}
	if password = r.PostForm.Get("password"); password == "" {
		writeErrorWithMessage(w, r, http.StatusBadRequest, "invalid password")
		return
	}
	if nickname = r.PostForm.Get("nickname"); nickname == "" {
		writeErrorWithMessage(w, r, http.StatusBadRequest, "invalid nickname")
		return
	}

	m, err := newModel(connString)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError)
		return
	}
	defer func() { _ = m.Close() }()

	u, err := m.GetUserByEmail(email)
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
	u.initUrl()
	writeResponse(w, r, u)
}
