package main

import (
	"net/http"
	"regexp"
	"os"

	"github.com/gin-gonic/gin"
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

func sendEmailHandler(c *gin.Context) {
	// Parse the JSON request body into the EmailContent struct
	var emailContent EmailContent
	if err := c.ShouldBindJSON(&emailContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}

	// Validate email addresses
	if !isEmailValid(emailContent.FromEmail) || !isEmailValid(emailContent.ToEmail) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email"})
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
			return
		}
		defer db.Close()

		// Save the response in the PostgreSQL database
		err = saveResponseToDatabase(db, statusCode, responseBody, headers)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving response to database"})
			return
		}

		c.String(http.StatusOK, "Email sent successfully!")
	} else {
		// Error: SendGrid API response indicates an error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email"})
	}
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
