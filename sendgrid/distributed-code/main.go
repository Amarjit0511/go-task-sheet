// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := loadEnvVariables()
	if err != nil {
		log.Fatal("Error loading environment variables:", err)
	}

	http.HandleFunc("/send-email", sendEmailHandler)

	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
