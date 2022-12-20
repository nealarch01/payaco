package models

import (
	"database/sql"
	"fmt"
)

var psqlDB *sql.DB = nil

// Initializes a database connection
func InitConnection() *sql.DB {
	if psqlDB != nil {
		return psqlDB // If a connection is already established, return it rather than creating a new one
	}
	db, err := sql.Open("postgres", "dbname=development port=4323 sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db.SetMaxOpenConns(10) // Set the maximum number of open connections to the database
	db.SetConnMaxIdleTime(1)
	psqlDB = db
	return psqlDB
}
