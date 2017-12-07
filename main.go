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
	Migrations()
	// var port = os.Getenv("PORT")
	var port = os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	fmt.Println("Port =", port)

	var router = NewRouter()

	// run migrations and seeds

	// because it's a serve!
	log.Fatal(http.ListenAndServe(port, router))
}
