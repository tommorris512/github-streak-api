package types

// Define a calendar type for GitHub API responses
type ContributionCalendar struct {
    Data struct {
        User struct {
            ContributionsCollection struct {
                ContributionCalendar struct {
                    TotalContributions int `json:"totalContributions"`
                    Weeks []struct {
                        ContributionDays []struct {
                            Date string `json:"date"`
                            ContributionCount int `json:"contributionCount"`
                        } `json:"contributionDays"`
                    } `json:"weeks"`
                } `json:"contributionCalendar"`
            } `json:"contributionsCollection"`
        } `json:"user"`
    } `json:"data"`
}

// Define a response type of structured JSON to return to API requests 
type ContributionDataResponse struct {
    Username string `json:"username"`
    TotalContributions int `json:"total_contributions"`
    MaxDailyContributions int `json:"max_daily_contributions"`
    MaxDailyContributionsDate string `json:"max_daily_contributions_date"`
    CurrentStreak int `json:"current_streak"`
    LongestStreak int `json:"longest_streak"`
}