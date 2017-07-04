package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{Name: "UserLogin", Method: "POST", Pattern: "/user/authenticate", HandlerFunc: authenticateUser},
	Route{Name: "UserReset", Method: "POST", Pattern: "/user/reset", HandlerFunc: mock},

	//users
	Route{Name: "UserList", Method: "GET", Pattern: "/users", HandlerFunc: listUsers},
	Route{Name: "UserGet", Method: "GET", Pattern: "/users/{userid:[0-9]+}", HandlerFunc: getUser},
	Route{Name: "UserCreate", Method: "POST", Pattern: "/users", HandlerFunc: parseForm(createUser)},
	Route{Name: "UserUpdate", Method: "PUT", Pattern: "/users/{userid:[0-9]+}", HandlerFunc: checkUserRight(parseForm(updateUser))},
	Route{Name: "UserDelete", Method: "DELETE", Pattern: "/users/{userid:[0-9]+}", HandlerFunc: checkUserRight(parseForm(deleteUser))},

	// barcode
	Route{Name: "Barcode", Method: "POST", Pattern: "/barcode", HandlerFunc: checkUserRight(parseForm(mock))},

	// books
	Route{Name: "BookList", Method: "GET", Pattern: "/books", HandlerFunc: mock},
	Route{Name: "BookGet", Method: "GET", Pattern: "/books/{id:[0-9]+}", HandlerFunc: mock},
	Route{Name: "BookCreate", Method: "POST", Pattern: "/books", HandlerFunc: checkUserRight(parseForm(mock))},
	Route{Name: "BookUpdate", Method: "PUT", Pattern: "/books/{bookidid:[0-9]+}", HandlerFunc: checkUserRight(parseForm(mock))},
	Route{Name: "BookDelete", Method: "DELETE", Pattern: "/books/{bookidid:[0-9]+}", HandlerFunc: checkUserRight(parseForm(mock))},

	// user books
	Route{Name: "UserBookList", Method: "GET", Pattern: "/users/{userid:[0-9]+}/books", HandlerFunc: mock},
	Route{Name: "UserBookGet", Method: "GET", Pattern: "/users/{userid:[0-9]+}/books/{bookid:[0-9]+}", HandlerFunc: mock},
	Route{Name: "UserBookCreate", Method: "POST", Pattern: "/users/{userid:[0-9]+}/books", HandlerFunc: checkUserRight(parseForm(mock))},
	Route{Name: "UserBookDelete", Method: "DELETE", Pattern: "/users/{userid:[0-9]+}/books/{bookid:[0-9]+}", HandlerFunc: checkUserRight(parseForm(mock))},
}

func mock(arg http.ResponseWriter, arg2 *http.Request) {

}
