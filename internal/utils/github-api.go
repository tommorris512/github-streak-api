package utils

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// Define a response type for the API response
type GitHubResponse struct {
    Data struct {
        User struct {
            ContributionsCollection struct {
                ContributionCalendar struct {
                    TotalContributions int `json:"totalContributions"`
                } `json:"contributionCalendar"`
            } `json:"contributionsCollection"`
        } `json:"user"`
    } `json:"data"`
}

func GetTotalContributions(username string, token string) (int, error) {
    // Define the GraphQL query
    query := fmt.Sprintf(`
        {
            user(login: "%s") {
                contributionsCollection {
                    contributionCalendar {
                        totalContributions
                    }
                }
            }
        }
    `, username)

    // Define the request body using the query
    reqBody := map[string]string{
        "query": query,
    }

    // Convert the request body to JSON
    jsonBody, err := json.Marshal(reqBody)
    if err != nil {
        return 0, err
    }

    // Initialise a new POST request to the GitHub api with the request body created
    req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonBody))
    if err != nil {
        return 0, err
    }

    // Set the request headers with token authorisation
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Content-Type", "application/json")

    // Initialise the client and execute the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    // Read the response body returned from the request
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return 0, err
    }

    // Initialise and populate a result object typed for the query
    var result GitHubResponse
    if err := json.Unmarshal(body, &result); err != nil {
        return 0, err
    }

    // Return the total contributions
    return result.Data.User.ContributionsCollection.ContributionCalendar.TotalContributions, nil
}
