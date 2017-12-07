package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// Index gets all records
func Index(w http.ResponseWriter, r *http.Request) {
	// sanity check
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

	// declarations
	var users = Users{
		Students{Name: "Clark Walker"},
		Students{Name: "Lizz Robbins"},
		Students{Name: "Cole Carol"},
		Students{Name: "Dakota Pfeifer"},
	}

	json.NewEncoder(w).Encode(users)
}
