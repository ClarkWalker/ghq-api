package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// sanity check
	fmt.Println("main.go has run")

	// declarations
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	var router = NewRouter()

	// run migrations and seeds
	Migrations()

	// because it's a serve!
	log.Fatal(http.ListenAndServe(port, router))
}
