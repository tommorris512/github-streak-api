package main

import (
    "fmt"
    "os"
    "net/http"
    "log"
    "github.com/joho/godotenv"
    "github-streak-api/handlers"
)

func main() {
	// Define the port number to run the server on
	const port int = 8080

    // Check whether the code is in development or production
    if envMode := os.Getenv("ENV_MODE"); envMode != "production" {
        // Attempt to load the .env file
        err := godotenv.Load()
        if err != nil {
            log.Fatalf("Error: %s\n", err)
            return
        }
    }

    githubToken := os.Getenv("GITHUB_TOKEN")
    if githubToken == "" {
        log.Fatal("Error: Unable to obtain GitHub access token")
        return
    }

    //Assign the handler function to the root path
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        // Extract the username parameter from the URL
        username := request.URL.Query().Get("username")
        if username == "" {
            http.Error(writer, "Username is required", http.StatusBadRequest)
            return
        }

        // Pass the username and GitHub token to the ContributionHandler
        handlers.ContributionHandler(writer, request, username, githubToken)
    })

	// Log the server is running
    log.Printf("Server started on port: %d", port)

	// Start the server and listen on the port, printing an error if encountered
    if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
        log.Fatalf("Error: %s\n", err)
        return
    }
}
