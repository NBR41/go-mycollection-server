package main

import (
	"net/http"
)

func authenticateUser(w http.ResponseWriter, r *http.Request) {
	var password, email, nickname string
	if password = r.PostForm.Get(formPasswordField); password == "" {
		writeErrorWithMessage(w, r, http.StatusBadRequest, "missing password")
		return
	}
	email = r.PostForm.Get(formEmailField)
	nickname = r.PostForm.Get(formNicknameField)
	if email == "" && nickname == "" {
		writeErrorWithMessage(w, r, http.StatusBadRequest, "missing email or nickname")
		return
	}

	m, err := NewModel(connString)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError)
		return
	}

	u, err := m.GetAuthenticatedUser(password, email, nickname)
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}

	if u == nil {
		writeErrorWithMessage(w, r, http.StatusUnauthorized, "invalid credentials")
		return
	}

	if !u.IsValidated {
		writeErrorWithMessage(w, r, http.StatusUnauthorized, "user not verified")
		return
	}

	tok, err := getToken(u)
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}

	writeResponse(w, r, struct {
		Token string `json:"token"`
	}{string(tok)})
}
