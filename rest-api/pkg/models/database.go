package models

import (
	"database/sql"
	"fmt"
)

// Initializes a database connection
func InitConnection() *sql.DB {
	db, err := sql.Open("postgres", "dbname=development port=4323 sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}
