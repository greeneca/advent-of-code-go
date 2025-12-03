package aocapi

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type AOCAPI struct {
	sessionToken string
}

func NewAOCAPI() (*AOCAPI, error) {
	token, err := os.ReadFile(".session_token")
	if err != nil {
		return nil, err
	}
	tokenString := strings.TrimSpace(string(token))
	return &AOCAPI{
		sessionToken: tokenString,
	}, nil
}

func (api *AOCAPI) GetInput(year, day int, args []string) ([]string, error) {
	path, err := getInputFilePath(args)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := fetchInputFile(year, day, path, api.sessionToken)
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
func fetchInputFile(year, day int, path, token string) error {
	fmt.Printf("Fetching input file: %s\n", path)
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	body, err := doRequestWithSession(token, req)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, body, 0644)
	return nil
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

func (api *AOCAPI) SubmitAnswer(year, day, part int, answer string) error {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)
	data := fmt.Sprintf("level=%d&answer=%s", part, answer)
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	_, err = doRequestWithSession(api.sessionToken, req)
	return err
}
func doRequestWithSession(token string, req *http.Request) ([]byte, error) {
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

