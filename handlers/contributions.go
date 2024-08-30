package handlers

import (
    "fmt"
    "net/http"
    "github-streak-api/utils"
)

func ContributionHandler(writer http.ResponseWriter, request *http.Request, username string, token string) {
    data, err := utils.GetContributionData(username, token)
    if err != nil {
        http.Error(writer, "Failed to get contributions", http.StatusInternalServerError)
        return
    }

    totalContributions := utils.CalculateTotalContributions(data)
    mostContribDate, mostContribCount := utils.CalculateMostDailyContributions(data)

    fmt.Fprintf(writer, "User %s has made %d contributions this year.\n", username, totalContributions)
    fmt.Fprintf(writer, "Most contributions in a day: %d on %s\n", mostContribCount, mostContribDate)
}
