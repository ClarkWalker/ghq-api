package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Links ...
// The Names Need To Start With An UpperCase Letter
// PascalCase PascalCase PascalCase PascalCase PascalCase PascalCase
type Links struct {
	gorm.Model
	FlexboxURL string
	SqlzooURL  string
	MdnURL     string
	KnexjsURL  string
	ExpressURL string
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
	// these are seeds ;)
	// The Names Need To Start With An UpperCase Letter
	// PascalCase PascalCase PascalCase PascalCase PascalCase PascalCase
	db.Create(&Links{ // Clark Walker
		// ID:              1,                                  // uint
		FlexboxURL: "http://flexboxfroggy.com/",            // string
		SqlzooURL:  "http://sqlzoo.net/",                   // string
		MdnURL:     "https://developer.mozilla.org/en-US/", // string
		KnexjsURL:  "http://knexjs.org/",                   // string
		ExpressURL: "https://expressjs.com/"})              // string

} // end main
