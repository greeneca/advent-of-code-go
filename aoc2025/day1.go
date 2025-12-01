package aoc2025

import (
	"fmt"
	"strconv"
)

func day1Part1(data []string) {
	dial := 50
	count := 0
	for _, line := range data {
		if len(line) < 2 {
			continue
		}
		change, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}
		if line[0] == 'L' {
			dial -= change
		} else if line[0] == 'R' {
			dial += change
		}
		dial %= 100
		if dial == 0 {
			count++
		}
	}
	fmt.Println("Password is:", count)
}

func day1Part2(data []string) {
	dial := 50
	count := 0
	for _, line := range data {
		if line == "" { continue }
		change, err := strconv.Atoi(line[1:])
		spins := change /100
		count += spins
		next := dial
		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}
		if line[0] == 'L' {
			next -= change % 100
		} else if line[0] == 'R' {
			next += change % 100
		}
		if next < 0 {
			if dial > 0 {
				count += 1
			}
			dial = 100 + next
		} else if next > 99 {
			count += 1
			dial = next % 100
		} else {
			if next == 0 {
				count++
			}
			dial = next
		}
	}
	fmt.Println("Password is:", count)
}
