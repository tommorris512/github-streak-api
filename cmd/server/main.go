package main

import (
    "fmt"
    "os"
    "net/http"
    "github.com/joho/godotenv"
    "github-streak-api/internal/handlers"
)

func main() {
	// Define the port number to run the server on
	const port int = 8080

    // Attempt to load the .env file
    err := godotenv.Load()
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }
    
    // Obtain the GitHub token
    githubToken := os.Getenv("GITHUB_TOKEN")
    fmt.Println("GitHub token found: %s", githubToken)

	// Assign the relevant handlers to paths
    http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/contributions", func(w http.ResponseWriter, r *http.Request) {
        handlers.ContributionHandler(w, r, "tommorris512", githubToken)
    })

	// Log the server is running
    fmt.Println("Server started on port: ", port)

	// Start the server and listen on the port, printing an error if encountered
    if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
        fmt.Printf("Error: %s\n", err)#
        return
    }
}
