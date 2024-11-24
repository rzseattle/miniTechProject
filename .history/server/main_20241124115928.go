package main

import "net/http"

func main() {
	// Create a new server
	server := http.Server{
		Addr:    ":4000",
		Handler: nil,
	}

	// Start the server
	server.ListenAndServe()
}
