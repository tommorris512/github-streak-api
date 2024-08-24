package main

import (
    "fmt"
    "net/http"
    "github-streak-api/internal/handlers"
)

func main() {
	// Define the port number to run the server on
	const port int = 8080

	// Assign the root path to the relevant handler
    http.HandleFunc("/", handlers.HomeHandler)

	// Log the server is running
    fmt.Println("Server started on port %d", port)

	// Start the server and listen on the port, printing an error if encountered
    if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
        fmt.Printf("Error: %s\n", err)
    }
}
