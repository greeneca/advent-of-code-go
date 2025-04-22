package aoc2017

import (
	"fmt"
	"strconv"
)

func day1Part1(data []string) {
	input := data[0]
	prev := string(input[len(input)-1])
	sum := 0
	for i := range len(input) {
		current := string(input[i])
		if current == prev {
			value, _ := strconv.Atoi(current)
			sum += value
		}
		prev = current
	}
	fmt.Println("Sum of matching digits:", sum)
}

func day1Part2(data []string) {
	input := data[0]
	length := len(input)
	halfway := length / 2
	sum := 0

	for i := range length {
		current := string(input[i])
		next := string(input[(i+halfway)%length])
		if current == next {
			value, _ := strconv.Atoi(current)
			sum += value
		}
	}
	fmt.Println("Sum of matching digits:", sum)
}
