package main

import (
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

func init() {
	// Open the database connection using the provided connection string.
	var err error
	db, err = sql.Open("postgres", "postgres://amarjit:amarjit@localhost/sendgriddb?sslmode=disable")
	if err != nil {
		panic(err)
	}
}
