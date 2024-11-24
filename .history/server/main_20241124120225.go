package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new server
	server := http.Server{
		Addr:    "0.0.0.0",
		Handler: nil,
	}

	http.HandleFunc("/hello", helloHandler)
	// Start the server
	fmt.Println("Starting server on port 4001")
	server.ListenAndServe(":3333", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
