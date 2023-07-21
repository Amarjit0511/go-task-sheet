// email.go
package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"fmt"
	"os"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// EmailContent struct for email data
type EmailContent struct {
	FromName         string `json:"from_name"`
	FromEmail        string `json:"from_email"`
	ToName           string `json:"to_name"`
	ToEmail          string `json:"to_email"`
	Subject          string `json:"subject"`
	PlainTextContent string `json:"plain_text_content"`
	HTMLContent      string `json:"html_content"`
}

func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body into the EmailContent struct
	var emailContent EmailContent
	err := json.NewDecoder(r.Body).Decode(&emailContent)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	// Validate email addresses
	if !isEmailValid(emailContent.FromEmail) || !isEmailValid(emailContent.ToEmail) {
		http.Error(w, "Invalid email address", http.StatusBadRequest)
		return
	}

	// Create the SendGrid email message
	message := mail.NewSingleEmail(
		mail.NewEmail(emailContent.FromName, emailContent.FromEmail),
		emailContent.Subject,
		mail.NewEmail(emailContent.ToName, emailContent.ToEmail),
		emailContent.PlainTextContent,
		emailContent.HTMLContent,
	)

	// Create the SendGrid client and send the email
	client := sendgrid.NewSendClient(os.Getenv(SendGridAPIKeyEnv))
	response, err := client.Send(message)
	if err != nil {
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}

	// Check the SendGrid API response for email success status
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		// Success: Extract the necessary data from the response
		statusCode := response.StatusCode
		responseBody := emailContent.PlainTextContent
		headers := response.Headers

		// Get the database connection
		db, err := initDB()
		if err != nil {
			http.Error(w, "Error initializing database", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Save the response in the PostgreSQL database
		err = saveResponseToDatabase(db, statusCode, responseBody, headers)
		if err != nil {
			http.Error(w, "Error saving response to database", http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "Email sent successfully!")
	} else {
		// Error: SendGrid API response indicates an error
		log.Printf("Error sending email. Response status code: %d", response.StatusCode)
		log.Printf("Response body: %s", response.Body)
		http.Error(w, "Error sending email", http.StatusInternalServerError)
	}
}


func isEmailValid(email string) bool {
	// ... (same as before)
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
