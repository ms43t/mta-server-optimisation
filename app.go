package main

import (
	"log"
	"net/http"
)

// main is the entry point of the program.
func main() {
	err := loadMockData() // Load mock data for testing.
	if err != nil {
		log.Fatal("Failed to load mock data: ", err)
	}

	http.HandleFunc("/mta-hosting-optimizer", getInstanceName) // Handle API requests.
	log.Fatal(http.ListenAndServe(":8080", nil))               // Start the HTTP server.
}
