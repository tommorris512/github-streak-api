package handlers

import (
	"fmt"
	"github-streak-api/utils"
	"net/http"
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

	fmt.Fprintf(writer, "User %s has made %d contributions this year.\n", username, totalContributions)
	fmt.Fprintf(writer, "Most contributions in a day: %d on %s\n", maxDailyContributions, maxDailyContributionsDate)
	fmt.Fprintf(writer, "Current streak: %d days\n", currentStreak)
	fmt.Fprintf(writer, "Longest streak: %d days\n", longestStreak)
}
