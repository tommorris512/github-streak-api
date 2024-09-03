package main

import (
    "fmt"
    "os"
    "net/http"
    "github.com/joho/godotenv"
    "github-streak-api/handlers"
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
    if githubToken == "" {
        fmt.Println("Error: Unable to obtain GitHub access token")
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
    fmt.Println("Server started on port: ", port)

	// Start the server and listen on the port, printing an error if encountered
    if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }
}
