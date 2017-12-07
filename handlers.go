package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// Index gets all records
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	var users = Users{
		Students{Name: "Clark Walker"},
		Students{Name: "Lizz Robbins"},
		Students{Name: "Cole Carol"},
		Students{Name: "Dakota Pfeifer"},
	}

	json.NewEncoder(w).Encode(users)
}
