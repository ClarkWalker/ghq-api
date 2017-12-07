package main

import (
	"encoding/json"
	"net/http"
)

// Index gets all records
func Index(w http.ResponseWriter, r *http.Request) {
	var users = Users{
		Students{Name: "Clark Walker"},
		Students{Name: "Lizz Robbins"},
		Students{Name: "Cole Carol"},
		Students{Name: "Dakota Pfeifer"},
	}

	json.NewEncoder(w).Encode(users)
}
