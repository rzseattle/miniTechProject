package main

import "net/http"

func main() {
	// Create a new server
	server := http.Server{
		Addr:    "",
		Handler: nil,
	}

	// Start the server
	server.ListenAndServe()
}
