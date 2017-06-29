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
	Route{Name: "UserLogin", Method: "POST", Pattern: "/user/login", HandlerFunc: mock},
	Route{Name: "UserReset", Method: "POST", Pattern: "/user/reset", HandlerFunc: mock},

	//user
	Route{Name: "UserList", Method: "GET", Pattern: "/users", HandlerFunc: mock},
	Route{Name: "UserCreate", Method: "POST", Pattern: "/user", HandlerFunc: mock},
	Route{Name: "UserGet", Method: "GET", Pattern: "/user/{id:[0-9]+}", HandlerFunc: mock},
	Route{Name: "UserUpdate", Method: "PUT", Pattern: "/user/{id:[0-9]+}", HandlerFunc: mock},
	Route{Name: "UserDelete", Method: "DELETE", Pattern: "/user/{id:[0-9]+}", HandlerFunc: mock},

	Route{Name: "UserLogin", Method: "POST", Pattern: "/user/login", HandlerFunc: mock},
	Route{Name: "UserReset", Method: "POST", Pattern: "/user/reset", HandlerFunc: mock},

	// barcode
	Route{Name: "Barcode", Method: "POST", Pattern: "/barcode", HandlerFunc: mock},

	// books
	Route{Name: "BookList", Method: "GET", Pattern: "/books", HandlerFunc: mock},
	Route{Name: "BookCreate", Method: "POST", Pattern: "/book", HandlerFunc: mock},
	Route{Name: "BookGet", Method: "GET", Pattern: "/book/{id:[0-9]+}", HandlerFunc: mock},
	Route{Name: "BookUpdate", Method: "PUT", Pattern: "/book/{id:[0-9]+}", HandlerFunc: mock},
	Route{Name: "BookDelete", Method: "DELETE", Pattern: "/book/{id:[0-9]+}", HandlerFunc: mock},
}
