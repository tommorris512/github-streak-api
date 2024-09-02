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

/*
 * Obtains the contribution count and date of the highest contributions made in a single day.
 *
 * Iterates through each day in the parsed ContributionCalendar object and checks the contribution count for each.
 * Maximum pointers are used to track the count and date of the highest daily contributions encountered in the search.
 * Once all days have been searched, the count and date of the highest daily contributions are returned.
 *
 * Params:
 * - data (types.ContributionCalendar): The ContributionCalendar used to obtain the highest daily contributions from.
 *
 * Returns:
 * - (int): The daily contributions made on the day with the highest daily contributions in the contribution calendar.
 * - (string): The date of the highest daily contributions in the contribution calendar.
 */
func CalculateMostDailyContributions(data types.ContributionCalendar) (int, string) {
	maxDailyContributions := 0
	maxDailyContributionsDate := ""

	for _, week := range data.Data.User.ContributionsCollection.ContributionCalendar.Weeks {
		for _, day := range week.ContributionDays {
			if day.ContributionCount > maxDailyContributions {
				maxDailyContributions = day.ContributionCount
				maxDailyContributionsDate = day.Date
			}
		}
	}

	return maxDailyContributions, maxDailyContributionsDate
}

/*
 * Obtains the longest streak of contributions.
 *
 * Iterates over each day in the parsed contribution calendar, incrementing the current streak being calculated if a contribution was made, and resetting it otherwise.
 * A maximum pointer is maintained throughout the iteration to store the longest streak found at any given point.
 * Once the entire calendar has been iterated over, the longest streak is returned.
 *
 * Params:
 * - data (types.ContributionCalendar): The ContributionCalendar used to obtain the longest contribution streak from.
 *
 * Returns:
 * - (int): The longest contribution streak in the contribution calendar.
 */
func CalculateLongestContributionStreak(data types.ContributionCalendar) int {
	longestStreak := 0
	streak := 0

	for _, week := range data.Data.User.ContributionsCollection.ContributionCalendar.Weeks {
		for _, day := range week.ContributionDays {
			if day.ContributionCount > 0 {
				streak++
			} else {
				if streak > longestStreak {
					longestStreak = streak
				}

				streak = 0
			}
		}
	}

	if streak > longestStreak {
		longestStreak = streak
	}

	return longestStreak
}

/*
 * Obtains the current streak of contributions.
 *
 * Iterates over each day in the ContributionCalendar in reverse-chronological order.
 * The current streak counter is continually incremented until a day is reached that has no contributions, and the current streak is returned.
 *
 * Params:
 * - data (types.ContributionCalendar): The ContributionCalendar used to obtain the current contribution streak from.
 *
 * Returns:
 * - (int): The current contribution streak in the contribution calendar.
 */
func CalculateCurrentContributionStreak(data types.ContributionCalendar) int {
	currentStreak := 0

	for i := len(data.Data.User.ContributionsCollection.ContributionCalendar.Weeks) - 1; i >= 0; i-- {
		week := data.Data.User.ContributionsCollection.ContributionCalendar.Weeks[i]

		for j := len(week.ContributionDays) - 1; j >= 0; j-- {
			day := week.ContributionDays[j]

			if day.ContributionCount <= 0 {
				return currentStreak
			}

			currentStreak++
		}
	}

	return currentStreak
}
