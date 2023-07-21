// database.go
package main

import (
	"database/sql"
	"encoding/json"
	"os"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

const (
	SendGridAPIKeyEnv = "SENDGRID_API_KEY"
	DBConnStringEnv   = "DB_CONNECTION_STRING"
)

func loadEnvVariables() error {
	err := godotenv.Load("config.env")
	if err != nil {
		return err
	}
	// Load other necessary environment variables here
	return nil
}

func initDB() (*sql.DB, error) {
	dbConnString := os.Getenv(DBConnStringEnv)
	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func saveResponseToDatabase(db *sql.DB, statusCode int, responseBody string, headers map[string][]string) error {
	// ... (same as before)
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
