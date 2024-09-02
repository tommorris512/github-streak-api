package handlers

import (
	"net/http"
    "encoding/json"
    "github-streak-api/types"
    "github-streak-api/utils"
)

func ContributionHandler(writer http.ResponseWriter, request *http.Request, username string, token string) {
	data, err := utils.GetContributionData(username, token)
	if err != nil {
		http.Error(writer, "Failed to get contributions", http.StatusInternalServerError)
		return
	}

	totalContributions := utils.CalculateTotalContributions(data)
	maxDailyContributions, maxDailyContributionsDate := utils.CalculateMostDailyContributions(data)
	currentStreak := utils.CalculateCurrentContributionStreak(data)
	longestStreak := utils.CalculateLongestContributionStreak(data)

    // Instantiate and populate a response object with the relevant data
	response := types.ContributionDataResponse{
		Username: username,
		TotalContributions: totalContributions,
		MaxDailyContributions: maxDailyContributions,
		MaxDailyContributionsDate: maxDailyContributionsDate,
		CurrentStreak: currentStreak,
		LongestStreak: longestStreak,
	}

	// Convert the response object to JSON format
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, "Error converting response to JSON", http.StatusInternalServerError)
		return
	}

	// Set the content type of the response to JSON and write the response
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonResponse)
}
