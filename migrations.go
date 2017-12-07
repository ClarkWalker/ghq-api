package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Students ...
type Students struct {
	gorm.Model
	// ID              uint
	Name            string
	Email           string
	Cohort          int
	GithubURL       string
	LinkedinURL     string
	SlackHandle     string
	Password        string
	Role            string
	SlackPassword   string
	AbsentExcused   int
	AbsentUnexcused int
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
	db.DropTable(&Students{})

	// Migrate taht ol schema
	db.AutoMigrate(&Students{})

	// these are seeds ;) // Create
	db.Create(&Students{ // Clark Walker
		// ID:              1,                                  // uint
		Name:            "Clark Walker",                     // string
		Email:           "303walker@gmail.com",              // string
		Cohort:          64,                                 // int
		GithubURL:       "https://github.com/ClarkWalker",   // string
		LinkedinURL:     "www.linkedin.com/in/clark-walker", // string
		SlackHandle:     "cMonster",                         // string
		Password:        "password123",                      // string
		Role:            "snarky student",                   // string
		SlackPassword:   "123password",                      // string
		AbsentExcused:   2,                                  // int
		AbsentUnexcused: 1})                                 // int

	db.Create(&Students{ // Dakota Pfeifer
		// ID:              2,                                  // uint
		Name:            "Dakota Pfeifer",             // string
		Email:           "dpfeif@outlook.com",         // string
		Cohort:          64,                           // int
		GithubURL:       "https://github.com/dpfeif",  // string
		LinkedinURL:     "www.linkedin.com/in/dpfeif", // string
		SlackHandle:     "dpfeif",                     // string
		Password:        "password123",                // string
		Role:            "fancy lad",                  // string
		SlackPassword:   "123password",                // string
		AbsentExcused:   2,                            // int
		AbsentUnexcused: 1})                           // int

} // end main
