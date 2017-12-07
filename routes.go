// Package ghq ...
package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
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

// main
func main() {
	// sainity check
	fmt.Println("router.go has run")
	var port = ":3000"
	Queries()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(port, router))

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	var users = Users{
		Students{Name: "Clark Walker"},
		Students{Name: "Lizz Robbins"},
		Students{Name: "Cole Carol"},
		Students{Name: "Dakota Pfeifer"},
	}
	json.NewEncoder(w).Encode(users)
}
