package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}

var routes = []route{
	{Name: "UserLogin", Method: "POST", Pattern: "/user/authenticate", HandlerFunc: authenticateUser},
	{Name: "UserReset", Method: "POST", Pattern: "/user/reset", HandlerFunc: mock},

	//users
	{Name: "UserList", Method: "GET", Pattern: "/users", HandlerFunc: listUsers},
	{Name: "UserGet", Method: "GET", Pattern: "/users/{userid:[0-9]+}", HandlerFunc: getUser},
	{Name: "UserCreate", Method: "POST", Pattern: "/users", HandlerFunc: parseForm(createUser)},
	{Name: "UserUpdate", Method: "PUT", Pattern: "/users/{userid:[0-9]+}", HandlerFunc: checkUserRight(parseForm(updateUser))},
	{Name: "UserDelete", Method: "DELETE", Pattern: "/users/{userid:[0-9]+}", HandlerFunc: checkUserRight(deleteUser)},

	// barcode
	{Name: "Barcode", Method: "POST", Pattern: "/barcode", HandlerFunc: checkUserRight(parseForm(mock))},

	// books
	{Name: "BookList", Method: "GET", Pattern: "/books", HandlerFunc: listBooks},
	{Name: "BookGet", Method: "GET", Pattern: "/books/{id:[0-9]+}", HandlerFunc: getBook},
	{Name: "BookCreate", Method: "POST", Pattern: "/books", HandlerFunc: checkUserRight(parseForm(createBook))},
	{Name: "BookUpdate", Method: "PUT", Pattern: "/books/{bookidid:[0-9]+}", HandlerFunc: checkUserRight(parseForm(updateBook))},
	{Name: "BookDelete", Method: "DELETE", Pattern: "/books/{bookidid:[0-9]+}", HandlerFunc: checkUserRight(deleteBook)},

	// user books
	{Name: "UserBookList", Method: "GET", Pattern: "/users/{userid:[0-9]+}/books", HandlerFunc: listUserBooks},
	{Name: "UserBookGet", Method: "GET", Pattern: "/users/{userid:[0-9]+}/books/{bookid:[0-9]+}", HandlerFunc: getUserBook},
	{Name: "UserBookCreate", Method: "POST", Pattern: "/users/{userid:[0-9]+}/books", HandlerFunc: checkUserRight(parseForm(createUserBook))},
	{Name: "UserBookDelete", Method: "DELETE", Pattern: "/users/{userid:[0-9]+}/books/{bookid:[0-9]+}", HandlerFunc: checkUserRight(deleteUserBook)},
}

func mock(arg http.ResponseWriter, arg2 *http.Request) {

}
