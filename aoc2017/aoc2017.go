package aoc2017

import (
	"errors"
	"fmt"
)

func RunProblem(day int, part int, data []string) error {
	problems := map[string]func([]string){
		"day1Part1": day1Part1,
		"day1Part2": day1Part2,
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
