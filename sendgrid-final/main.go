package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func loadEnvVariables() error {
	err := godotenv.Load("config.env")
	if err != nil {
		return err
	}
	return nil
}

func initialisingDB() (*sql.DB, error) {
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

func sendEmailHandler(c *gin.Context) {
	// Checking if the JSON request matches with the data structure of EmailContent
	var emailContent EmailContent
	if err := c.ShouldBindJSON(&emailContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Sorry": "Invalid JSON request"})
		return
	}

	// Validating the email address of the sender and receiver
	if !isEmailValid(emailContent.FromEmail) || !isEmailValid(emailContent.ToEmail) {
		c.JSON(http.StatusBadRequest, gin.H{"Sorry": "Please check the email id again"})
		return
	}

	// Creating an email that will be send by SendGrid
	message := mail.NewSingleEmail(
		mail.NewEmail(emailContent.FromName, emailContent.FromEmail),
		emailContent.Subject,
		mail.NewEmail(emailContent.ToName, emailContent.ToEmail),
		emailContent.PlainTextContent,
		emailContent.HTMLContent,
	)

	// Create a new client after verifying the API key who will be sending the email
	client := sendgrid.NewSendClient(os.Getenv(SendGridAPIKeyEnv))
	response, err := client.Send(message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Sorry": "Error sending email"})
		return
	}

	// Taking a note of the response only when the email is successfully send to the recipient
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		// Extracting different data from the response
		statusCode := response.StatusCode
		responseBody := response.Body
		headers := response.Headers

		// Starting the DB connection
		db, err := initialisingDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Sorry": "Error initializing database"})
			return
		}
		defer db.Close()

		// Saving the response to the db
		err = saveResponseToDatabase(db, statusCode, responseBody, headers)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Sorry": "Error saving response to database"})
			return
		}

		c.String(http.StatusOK, "Email sent successfully to the recipient!")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Sorry": "Error sending email, please recheck the email"})
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
	// Enable CORS for all routes
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	r.POST("/send-email", sendEmailHandler)

	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
