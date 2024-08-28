package handlers

import (
    "fmt"
    "net/http"
)

// Define a basic request handler to print a welcome message
func HomeHandler(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Welcome to the GitHub Streak API")
}