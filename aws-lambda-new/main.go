package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"
	"regexp"

	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Data structure for email data
type EmailContent struct {
	FromName         string `json:"from_name"`
	FromEmail        string `json:"from_email"`
	ToName           string `json:"to_name"`
	ToEmail          string `json:"to_email"`
	Subject          string `json:"subject"`
	PlainTextContent string `json:"plain_text_content"`
	HTMLContent      string `json:"html_content"`
}

// Can be accessible by any part of the code
const (
	SendGridAPIKeyEnv = "SENDGRID_API_KEY"
	DBConnStringEnv   = "DB_CONNECTION_STRING"
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
	// Converting headers(map) to a JSON string
	headersJSON, err := json.Marshal(headers)
	if err != nil {
		return err
	}

	// Creating a statement that can be executed later any number of time 
	statement, err := db.Prepare("INSERT INTO sendgrid_response (status_code, body, headers) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer statement.Close()

	// Executing the insert statement
	_, err = statement.Exec(statusCode, responseBody, headersJSON)
	if err != nil {
		return err
	}

	return nil
}

func sendEmailHandler(ctx context.Context, emailContent EmailContent) (string, error) {
	// Validating the email address of the sender and receiver
	if !isEmailValid(emailContent.FromEmail) || !isEmailValid(emailContent.ToEmail) {
		return "Invalid email address", nil
	}

	// Creating an email that will be sent by SendGrid
	message := mail.NewSingleEmail(
		mail.NewEmail(emailContent.FromName, emailContent.FromEmail),
		emailContent.Subject,
		mail.NewEmail(emailContent.ToName, emailContent.ToEmail),
		emailContent.PlainTextContent,
		emailContent.HTMLContent,
	)

	// Creating a new client after verifying the API key who will be sending the email
	client := sendgrid.NewSendClient(os.Getenv(SendGridAPIKeyEnv))
	response, err := client.Send(message)
	if err != nil {
		return "Error sending email", err
	}

	// Taking a note of the response only when the email is successfully sent to the recipient
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		// Extracting different data from the response
		statusCode := response.StatusCode
		responseBody := response.Body
		headers := response.Headers

		// Starting the DB connection
		db, err := initDB()
		if err != nil {
			return "Error initializing database", err
		}
		defer db.Close()

		// Saving the response to the db
		err = saveResponseToDatabase(db, statusCode, responseBody, headers)
		if err != nil {
			return "Error saving response to database", err
		}

		return "Email sent successfully to the recipient!", nil
	} else {
		return "Error sending email, please recheck the email", nil
	}
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func main() {
	// Starting the Lambda function
	lambda.Start(sendEmailHandler)
}
