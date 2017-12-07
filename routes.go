// Package ghq ...
package main

import (
	"net/http"
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
	Route{
		"Index",
		"GET",
		"/index",
		Index,
	},
	Route{
		"Reset",
		"GET",
		"/reset",
		Index,
	},
}
