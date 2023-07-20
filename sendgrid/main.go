package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	_ "github.com/lib/pq"
)

func saveResponseToDatabase(statusCode int, responseBody string, headers map[string][]string) error {
	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", "postgres://amarjit:amarjit@localhost/sendgriddb?sslmode=disable")
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

func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	from := mail.NewEmail("Amarjit", "amarjitkrxs@gmail.com")
	subject := "Sending with Twilio SendGrid is Fun"
	to := mail.NewEmail("Amarjit Kumar", "amarjitkr0511@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
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
		http.Error(w, "Error saving response to database", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Email sent successfully!")
}

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading config.env file:", err)
	}

	http.HandleFunc("/send-email", sendEmailHandler)

	// You can also add a separate endpoint to save the response to the database, for example:
	//http.HandleFunc("/save-response", saveResponseHandler)

	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
