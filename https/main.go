package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from config.env file
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database.
	err = connectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Setup Gin router.
	router := gin.Default()
	setupRoutes(router)

	serverAddr := "localhost:8082"
	log.Println("Server running at:", serverAddr)

	// Replace these paths with the actual paths to your SSL certificate and key files.
	sslCertPath := "/Users/amarjitkumar/Desktop/key/ssl_cert.pem"
	sslKeyPath := "/Users/amarjitkumar/Desktop/key/ssl_key.pem"

	// Run the server with HTTPS
	err = router.RunTLS(serverAddr, sslCertPath, sslKeyPath)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
