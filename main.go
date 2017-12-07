package main

import (
	"fmt"
	"log"
	"net/http"
	// "encoding/json"
	// "fmt"
	// "html"
	// "log"
	// "net/http"
	// "reflect"
	//
	// "github.com/gorilla/mux" // currently I prefer this over chi
	// "github.com/go-chi/chi"  // this is different http server that I didn't use
)

func main() {
	// sanity check
	fmt.Println("main.go has run")

	// declarations
	var port = ":3000"
	var router = NewRouter()

	// run migrations and seeds
	// Migrations()

	// because it's a serve!
	log.Fatal(http.ListenAndServe(port, router))
}
