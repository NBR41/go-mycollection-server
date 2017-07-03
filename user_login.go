package main

import (
	"net/http"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

func authenticateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		writeError(w, r, http.StatusBadRequest)
		return
	}

	var password, email, nickname string
	if password = r.PostForm.Get("password"); password == "" {
		writeErrorWithMessage(w, r, http.StatusBadRequest, "invalid password")
		return
	}
	if email = r.PostForm.Get("email"); email == "" {
		writeErrorWithMessage(w, r, http.StatusBadRequest, "invalid email")
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

	claims := jws.Claims{}
	jwt := jws.NewJWT(claims, crypto.SigningMethodHS256)
	tok, err := jwt.Serialize([]byte("secret salt"))
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable)
		return
	}
	writeResponse(w, r, struct {
		Token string `json:"token"`
	}{string(tok)})
}
