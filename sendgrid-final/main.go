package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

// Constants for environment variable keys
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

func main() {
	err := loadEnvVariables()
	if err != nil {
		log.Fatal("Error loading environment variables:", err)
	}

	r := gin.Default()
	r.POST("/send-email", sendEmailHandler)

	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}