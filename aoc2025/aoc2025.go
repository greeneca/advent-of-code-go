package aoc2025

func GetProblems() map[string]func([]string)string {
	return map[string]func([]string)string{
		"day1Part1": day1Part1,
		"day1Part2": day1Part2,
		"day2Part1": day2Part1,
		"day2Part2": day2Part2,
		"day3Part1": day3Part1,
		"day3Part2": day3Part2,
		// Add other problems here
	}
}
