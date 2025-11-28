package aoc2017

import (
	"fmt"
	"strings"
)

func day9Part1(data []string) {
	input := strings.Split(data[0], "")
	level := 0
	garbage := false
	score := 0
	for i := 0; i < len(input); i++ {
		char := input[i]
		if char == "!" {
			i++
			continue
		}
		if garbage {
			if char == ">" {
				garbage = false
			}
			continue
		}
		if char == "<" {
			garbage = true
			continue
		}
		if char == "{" {
			level++
			continue
		}
		if char == "}" {
			score += level
			level--
			continue
		}
	}
	fmt.Println("Total score:", score)
}

func day9Part2(data []string) {
	input := strings.Split(data[0], "")
	garbage := false
	charCount := 0
	for i := 0; i < len(input); i++ {
		char := input[i]
		if char == "!" {
			i++
			continue
		}
		if garbage {
			if char == ">" {
				garbage = false
			} else {
				charCount++
			}
			continue
		}
		if char == "<" {
			garbage = true
			continue
		}
	}
	fmt.Println("Total char count:", charCount)
}
