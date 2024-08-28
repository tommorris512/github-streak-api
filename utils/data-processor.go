package utils

import (
	"github-streak-api/types"
)

/*
 * Obtains the total contribution count from a given ContributionCalendar.
 *
 * Params:
 * - data (types.ContributionCalendar): The ContributionCalendar to obtain the contribution count from.
 *
 * Returns:
 * - (int): The total contributions made in the contribution calendar.
*/
func CalculateTotalContributions(data types.ContributionCalendar) int {
    return data.Data.User.ContributionsCollection.ContributionCalendar.TotalContributions
}
