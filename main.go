package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

var connString string

func mock(arg http.ResponseWriter, arg2 *http.Request) {

}
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
