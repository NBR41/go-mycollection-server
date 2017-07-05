package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

var connString string
var secretSalt string

const (
	formEmailField    = "email"
	formPasswordField = "password"
	formNicknameField = "nickname"

	formBookNameField = "book_name"

	urlUserIDField = "user_id"
	urlBookIDField = "book_id"
)

func init() {
	connString = os.Getenv("DB_CONN_STR")
	if connString == "" {
		panic("no connexion string")
	}

	secretSalt = os.Getenv("WS_SECRET_SALT")
	if secretSalt == "" {
		panic("no connexion secret salt")
	}
}

func main() {
	r := newRouter()
	http.ListenAndServe(":80", handlers.RecoveryHandler()(r))
}
