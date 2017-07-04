package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

var connString string

const (
	secretSalt = "secret salt"

	formEmailField    = "email"
	formPasswordField = "password"
	formNicknameField = "nickname"

	urlUserIDField = "user_id"
	urlBookIDField = "book_id"
)

func init() {
	connString := os.Getenv("DB_CONN_STR")
	if connString == "" {
		panic("no connexion string")
	}
}
func main() {
	r := NewRouter()
	http.ListenAndServe(":1123", handlers.RecoveryHandler()(r))
}
