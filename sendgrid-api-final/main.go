package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
