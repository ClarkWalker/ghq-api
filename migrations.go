package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Links ...
type Links struct {
	gorm.Model
	flexboxURL string
	sqlzooURL  string
	mdnURL     string
	knexjsURL  string
	expressURL string
}

// Migrations ...
func Migrations() {
	//
	fmt.Println("Migration")

	// connection
	db, err := gorm.Open("postgres", "host=localhost dbname=go_getters_g_portal sslmode=disable")
	// on connection err
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Drops the table to rest the database
	db.DropTable(&Links{})
	// Migrate taht ol schema
	db.AutoMigrate(&Links{})
	// these are seeds ;) // Create
	db.Create(&Links{ // Clark Walker
		// ID:              1,                                  // uint
		flexboxURL: "http://flexboxfroggy.com/",            // string
		sqlzooURL:  "http://sqlzoo.net/",                   // string
		mdnURL:     "https://developer.mozilla.org/en-US/", // string
		knexjsURL:  "http://knexjs.org/",                   // string
		expressURL: "https://expressjs.com/"})              // string

} // end main
