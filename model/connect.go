package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

//Connect function
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to the database")
	con = db
	return db
}
