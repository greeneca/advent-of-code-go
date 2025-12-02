package aoc2017

import (
	"slices"
	"strconv"
	"strings"
)

func day4Part1(data []string) string {
	valid := 0
	for _, line := range data {
		if len(line) > 0 && isValidPassphrase(line) {
			valid++
		}
	}
	return strconv.Itoa(valid)
}

func isValidPassphrase(line string) bool {
	wordMap := map[string]bool{}
	words := strings.Fields(line)
	for _, word := range words {
		if _, exists := wordMap[word]; exists {
			return false
		}
		wordMap[word] = true
	}
	return true
}

func day4Part2(data []string) string {
	valid := 0
	for _, line := range data {
		if len(line) > 0 && isValidPassphraseAnagram(line) {
			valid++
		}
	}
	return strconv.Itoa(valid)
}

func isValidPassphraseAnagram(line string) bool {
	wordMap := map[string]bool{}
	words := strings.Fields(line)
	for _, word := range words {
		chars := strings.Split(word, "")
		slices.SortFunc(chars, func(a, b string) int {
			return strings.Compare(a, b)
		})
		word = strings.Join(chars, "")
		if _, exists := wordMap[word]; exists {
			return false
		}
		wordMap[word] = true
	}
	return true
}
