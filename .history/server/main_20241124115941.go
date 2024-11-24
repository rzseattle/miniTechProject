package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new server
	server := http.Server{
		Addr:    ":4000",
		Handler: nil,
	}

	// Start the server
	fmt.Println("Starting server on port 4000")
	server.ListenAndServe()
}
