package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	_ "github.com/lib/pq"
)

func saveResponseToDatabase(statusCode int, responseBody string, headers map[string][]string) error {
	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		return err
	}
	defer db.Close()

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

func sendEmailHandler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	from := mail.NewEmail("Amarjit", "amarjitkrxs@gmail.com")
	subject := "Sending with Twilio SendGrid is Fun"
	to := mail.NewEmail("Amarjit Kumar", "amarjitkr0511@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Error sending email", StatusCode: 500}, err
	}

	// Extract the necessary data from the response
	statusCode := response.StatusCode
	responseBody := plainTextContent
	headers := response.Headers

	// Convert the response body to a string
	body := string(responseBody)

	// Save the response in the PostgreSQL database
	err = saveResponseToDatabase(statusCode, body, headers)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Error saving response to database", StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{Body: "Email sent successfully!", StatusCode: 200}, nil
}

func main() {
	lambda.Start(sendEmailHandler)
}
