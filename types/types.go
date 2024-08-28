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