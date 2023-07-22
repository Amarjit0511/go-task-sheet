package main

import (
	"crypto/tls"
	"net/http"

)

func main() {
	// Load SSL/TLS certificates
	cert, err := tls.LoadX509KeyPair("/Users/amarjitkumar/Desktop/key/ssl_cert.pem", "/Users/amarjitkumar/Desktop/key/ssl_key.pem")
	if err != nil {
		panic(err)
	}

	// Create TLS configuration with loaded certificates
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	// Create the Gin router and set up routes
	router := SetupRouter()

	// Create an HTTPS server with TLS configuration
	server := &http.Server{
		Addr:      ":8080", // Change the port to the desired HTTPS port
		Handler:   router,
		TLSConfig: tlsConfig,
	}

	// Run the HTTPS server
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}
