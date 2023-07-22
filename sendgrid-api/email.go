package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"regexp"
	"os"

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

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func parseEmailContent(r *http.Request) (*EmailContent, error) {
	var emailContent EmailContent
	err := json.NewDecoder(r.Body).Decode(&emailContent)
	if err != nil {
		return nil, err
	}

	// Validate email addresses
	if !isEmailValid(emailContent.FromEmail) || !isEmailValid(emailContent.ToEmail) {
		return nil, fmt.Errorf("invalid email address")
	}

	return &emailContent, nil
}

func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	emailContent, err := parseEmailContent(r)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
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
		log.Println("Error sending email:", err)
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
			log.Println("Error initializing database:", err)
			return
		}
		defer db.Close()

		// Save the response in the PostgreSQL database
		err = saveResponseToDatabase(db, statusCode, responseBody, headers)
		if err != nil {
			http.Error(w, "Error saving response to database", http.StatusInternalServerError)
			log.Println("Error saving response to database:", err)
			return
		}

		fmt.Fprintln(w, "Email sent successfully!") 
	} else {
		// Error: SendGrid API response indicates an error
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		log.Println("Error sending email: SendGrid API response status code:", response.StatusCode)
	}
}
