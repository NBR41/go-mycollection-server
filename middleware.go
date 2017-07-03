package main

import (
	"net/http"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/gorilla/mux"
)

func getToken(u *user) ([]byte, error) {
	claims := jws.Claims{}
	claims.SetIssuer("myinventory-server")
	claims.Set("user", u.ID)
	claims.Set("admin", u.IsAdmin)
	jwt := jws.NewJWT(claims, crypto.SigningMethodHS256)
	return jwt.Serialize(secretSalt)
}

func checkUserRight(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			jwt, err := jws.ParseJWTFromRequest(r)
			if err != nil {
				writeError(w, r, http.StatusUnauthorized)
				return
			}

			err = jwt.Validate(secretSalt, crypto.SigningMethodHS256)
			if err != nil {
				writeError(w, r, http.StatusUnauthorized)
				return
			}

			cl := jwt.Claims()
			vars := mux.Vars(r)
			if vars["user_id"] != cl.Get("user") && false == cl.Get("admin") {
				writeError(w, r, http.StatusUnauthorized)
				return
			}
			f.ServeHTTP(w, r)
		},
	)
}
