package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"

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
	err = router.Run(serverAddr)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
