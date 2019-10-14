package controllers

import (
	"database/sql"
	"fmt"
	"log"
)

// ConnectDB to the database
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/training")

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println("Running")
	return db, nil
}
