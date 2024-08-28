package utils

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "github-streak-api/types"
)

// Set the GitHub GraphQL API route to utilise in fetching
const githubApiUrl string = "https://api.github.com/graphql"

/*
 * Obtains the contribution data of a given GitHub user.
 *
 * Defines a query fetching the contribution calendar of the parsed user and sends a POST request to GitHub's GraphQL API.
 * The request is parsed with the token environment variable for authorisation and a response is received.
 * A ContributionCalendar object is populated from the response and returned, alongside an error object to indicate success.
 *
 * Params:
 * - username (string): The username to fetch contribution data for.
 * - token (string): The GitHub token to use to authenticate access to the API. 
 *
 * Returns:
 * - (types.ContributionCalendar): A populated ContributionCalendar object containing the user's contribution data.
 * - (error): An error object that indicates if something went wrong during the request or parsing process (nil if successful). 
*/
func GetContributionData(username string, token string) (types.ContributionCalendar, error) {
	// Define the GraphQL query to fetch the user's contribution calendar
    query := fmt.Sprintf(`
        {
            user(login: "%s") {
                contributionsCollection {
                    contributionCalendar {
                        totalContributions
                        weeks {
                            contributionDays {
                                date
                                contributionCount
                            }
                        }
                    }
                }
            }
        }
    `, username)

	// Define the request body using the query
    reqBody := map[string]string{
        "query": query,
    }

	// Parse the request body to JSON format
    jsonBody, err := json.Marshal(reqBody)

    if err != nil {
        return types.ContributionCalendar{}, err
    }

	// Initialise a POST request to the GitHub API with the JSON request body
    req, err := http.NewRequest("POST", githubApiUrl, bytes.NewBuffer(jsonBody))

    if err != nil {
        return types.ContributionCalendar{}, err
    }

	// Set the request headers, authorising with the fetched GitHub token
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Content-Type", "application/json")

	// Send the request and await a response
    client := &http.Client{}
    resp, err := client.Do(req)

    if err != nil {
        return types.ContributionCalendar{}, err
    }

    defer resp.Body.Close()

	// Read the response body returned from the request
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        return types.ContributionCalendar{}, err
    }

	// Initialise and populate a result object typed for the query
    var result types.ContributionCalendar

    if err := json.Unmarshal(body, &result); err != nil {
        return types.ContributionCalendar{}, err
    }

	// Return the fetched contribution calendar object of the user
    return result, nil
}