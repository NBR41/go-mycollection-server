package main

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func mock(arg http.ResponseWriter, arg2 *http.Request) {

}

func main() {
	r := NewRouter()
	http.ListenAndServe(":1123", handlers.RecoveryHandler()(r))
}
