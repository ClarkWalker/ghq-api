// Package ghq ...
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Route a basic form for writing routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes slices Route ...?
type Routes []Route

// 'dem routes
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}

// NewRouter ...
func NewRouter() *mux.Router {
	// sanity check
	fmt.Println("routes.go has run")

	// declarations
	var router = mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(route.HandlerFunc)
	}

	return router
}
