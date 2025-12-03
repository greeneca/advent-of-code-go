package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/greeneca/advent-of-code-go/aoc2017"
	"github.com/greeneca/advent-of-code-go/aoc2025"
	api "github.com/greeneca/advent-of-code-go/aocApi"
)

func getYear(year int) (map[string]func([]string)string, error) {
	switch year {
	case 2017:
		return aoc2017.GetProblems(), nil
	case 2025:
		return aoc2025.GetProblems(), nil
	// Add more years here as needed
	default:
		return nil, errors.New("Year not yet implemented")
	}
}
func main() {
	api, err := api.NewAOCAPI()
	if err != nil {
		fmt.Printf("Error initializing AOC API: %v\n", err)
		return
	}
	args := os.Args[1:]
	start := time.Now()
	year, day, part, err := getProblemValues(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := api.GetInput(year, day, args)
	if err != nil {
		fmt.Printf("Error reading/fetching input file: %v\n", err)
		return
	}
	yearProblems, err := getYear(year)
	if err != nil {
		fmt.Printf("Error getting problems: %v\n", err)
		return
	}
	problem, err := getProblem(yearProblems, day, part)
	if err != nil {
		fmt.Printf("Error getting problem: %v\n", err)
		return
	}
	result := problem(data)
	fmt.Printf("Result: %v\n", result)
	duration := time.Since(start)
	fmt.Printf("Execution time: %s\n", duration)
}
func getProblemValues(args []string) (int, int, int, error) {
	if len(args) < 2 {
		return 0, 0, 0, errors.New("Invalid Usage: go run aoc.go <year> <problem> (input file)")
	}
	year, _ := strconv.Atoi(args[0])
	problem := strings.Split(args[1], "-")
	day, _ := strconv.Atoi(problem[0])
	part, _ := strconv.Atoi(problem[1])
	return year, day, part, nil
}
func getProblem(yearProblems map[string]func([]string)string, day, part int) (func([]string)string, error) {
	key := fmt.Sprintf("day%dPart%d", day, part)
	problem, exists := yearProblems[key]
	if !exists {
		return nil, errors.New("Problem not yet implemented")
	}
	return problem, nil
}

