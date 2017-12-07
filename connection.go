package main

import (
	"github.com/jinzhu/gorm"
)

// Db ...
func Db() {
	// connection
	Connection, err := gorm.Open("postgres", "host=localhost dbname=go_getters_g_portal sslmode=disable")

	// on Connect err
	if err != nil {
		panic("failed to connect database")
	}
	defer Connection.Close()
}
