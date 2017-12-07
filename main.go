package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// sanity check
	fmt.Println("main.go has run")

	// declarations
	var port = ":3000"
	var router = NewRouter()

	// run migrations and seeds
	Queries()

	// because it's a serve!
	log.Fatal(http.ListenAndServe(port, router))
}
