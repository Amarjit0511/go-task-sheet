package main

import (
	"database/sql"
	"encoding/json"
	"os"

	_ "github.com/lib/pq"
)

func initDB() (*sql.DB, error) {
	dbConnString := os.Getenv(DBConnStringEnv)
	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func saveResponseToDatabase(db *sql.DB, statusCode int, responseBody string, headers map[string][]string) error {
	// Convert headers map to a JSON string
	headersJSON, err := json.Marshal(headers)
	if err != nil {
		return err
	}

	// Prepare the INSERT statement
	stmt, err := db.Prepare("INSERT INTO sendgrid_response (status_code, body, headers) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the INSERT statement
	_, err = stmt.Exec(statusCode, responseBody, headersJSON)
	if err != nil {
		return err
	}

	return nil
}
