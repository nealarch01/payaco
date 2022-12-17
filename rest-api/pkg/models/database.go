package models

import (
	"database/sql"
	"fmt"
	"context"
)

var psqlDB *sql.DB

// Initializes a database connection
func InitConnection() *sql.DB {
	db, err := sql.Open("postgres", "dbname=development port=4323 sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	psqlDB = db
	return db
}

func GetConnection() *sql.Conn {
	conn, err := psqlDB.Conn(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return conn
}
