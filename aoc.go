package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/greeneca/advent-of-code-go/aoc2017"
	"github.com/greeneca/advent-of-code-go/aoc2025"
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
	args := os.Args[1:]
	start := time.Now()
	year, day, part, err := getProblemValues(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := getInput(args)
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
func getInput(args []string) ([]string, error) {
	path, err := getInputFilePath(args)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := fetchInputFile(path)
		if err != nil {
			return nil, err
		}
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	return lines, nil
}
func getInputFilePath(args []string) (string, error) {
	folder := fmt.Sprintf("inputs/%s", args[0])
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			return "", err
		}
	}
	if len(args) > 2 {
		path := fmt.Sprintf("%s/%s.txt", folder, args[2])
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("Custom input file does not exist. Falling back to default input file.")
		} else {
			return path, nil
		}
	}
	problem := strings.Split(args[1], "-")
	path := fmt.Sprintf("%s/%s.txt", folder, problem[0])
	return path, nil
}
func fetchInputFile(path string) error {
	fmt.Printf("Fetching input file: %s\n", path)
	token, err := getSessionToken()
	if err != nil {
		return err
	}
	year := strings.Split(path, "/")[1]
	day := strings.Split(strings.Split(path, "/")[2], ".")[0]
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)
	body, err := doRequestWithSession(url, token)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, body, 0644)
	return nil
}
func getSessionToken() (string, error) {
	token, err := os.ReadFile(".session_token")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(token)), nil
}
func doRequestWithSession(url, token string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: token})
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
func getProblem(yearProblems map[string]func([]string)string, day, part int) (func([]string)string, error) {
	key := fmt.Sprintf("day%dPart%d", day, part)
	problem, exists := yearProblems[key]
	if !exists {
		return nil, errors.New("Problem not yet implemented")
	}
	return problem, nil
}

