package aoc2017

import (
	"errors"
	"fmt"
)

func RunProblem(day int, part int, data []string) error {
	problems := map[string]func([]string){
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
		"day8Part1": day8Part1,
		"day8Part2": day8Part2,
		"day9Part1": day9Part1,
		"day9Part2": day9Part2,
		// Add other problems here
	}
	functionName := fmt.Sprintf("day%dPart%d", day, part)
	if problemFunc, exists := problems[functionName]; exists {
		problemFunc(data)
	} else {
		return errors.New("problem not implemented "+functionName)
	}

	return nil
}
