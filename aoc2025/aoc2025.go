package aoc2025

func GetProblems() map[string]func([]string)string {
	return map[string]func([]string)string{
		"day1Part1": day1Part1,
		"day1Part2": day1Part2,
		"day2Part1": day2Part1,
		"day2Part2": day2Part2,
		"day3Part1": day3Part1,
		"day3Part2": day3Part2,
		"day4Part1": day4Part1,
		"day4Part2": day4Part2,
		"day5Part1": day5Part1,
		"day5Part2": day5Part2,
		"day6Part1": day6Part1,
		"day6Part2": day6Part2,
		"day7Part1": day7Part1,
		"day7Part2": day7Part2,
		// Add other problems here
	}
}
