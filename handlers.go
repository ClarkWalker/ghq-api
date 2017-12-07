package main

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

// Index gets all records
func Index(w http.ResponseWriter, r *http.Request) {
	// I wanna break this out to a connection.go file
	db, err := gorm.Open("postgres", "host=localhost dbname=go_getters_g_portal sslmode=disable")
	// on connection err
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// this is a working get to / it returns all records
	var students []Students
	var getAll = db.Find(&students)
	json.NewEncoder(w).Encode(getAll)
}

// Reset runs migrations
func Reset(w http.ResponseWriter, r *http.Request) {
	// I wanna break this out to a connection.go file
	db, err := gorm.Open("postgres", "host=localhost dbname=go_getters_g_portal sslmode=disable")
	// on connection err
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	Migrations()
}

// // Read
// db.Find(&student)
// var getAll = db.Find(&student)
// fmt.Println(getAll)

// // Update
// db.Model(&student).Update("Role", "vary snarky student")

// // Delete
// db.Delete(&student)
