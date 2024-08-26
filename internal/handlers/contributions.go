package handlers

import (
    "fmt"
    "net/http"
    "github-streak-api/internal/utils"
)

func ContributionHandler(writer http.ResponseWriter, request *http.Request, username string, token string) {
    // Attempt to obtain the contribution count and any potential errors
    contributions, err := utils.GetTotalContributions(username, token)
	
    // Exit on presence of an error
    if err != nil {
        http.Error(writer, "Failed to get contributions", http.StatusInternalServerError)
        return
    }

    // Output the contribution count to the page
    fmt.Fprintf(writer, "User %s has made %d contributions this year.", username, contributions)
}
