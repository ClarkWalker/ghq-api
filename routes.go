// Package ghq ...
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	// "encoding/json"
	// "fmt"
	// "html"
	// "log"
	// "net/http"
	// "reflect"
	//
	// "github.com/gorilla/mux" // currently I prefer this over chi
	// "github.com/go-chi/chi"
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

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	// Route{
	// 	"TodoIndex",
	// 	"GET",
	// 	"/todos",
	// 	TodoIndex,
	// },
	// Route{
	// 	"TodoShow",
	// 	"GET",
	// 	"/todos/{todoId}",
	// 	TodoShow,
	// },
}
