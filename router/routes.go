package router

import (
	"net/http"

	handler "github/Tanawat_DEVz/example-rest-api2/handlers"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{"Get", "GET", "/people", handler.Get},
	Route{"Find", "GET", "/people/{id}", handler.Find},
	Route{"Create", "POST", "/people", handler.Create},
}
