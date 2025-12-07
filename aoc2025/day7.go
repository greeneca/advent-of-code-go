package aoc2025

import (
	"strconv"
	"strings"
)

func day7Part1(data []string) string {
	count := 0
	startLine := data[0]
	beams := map[int]bool{
		strings.Index(startLine, "S"): true,
	}
	data = data[1:]
	for _, line := range data {
		if line == "" { continue }
		newBeams := map[int]bool{}
		for beam := range beams {
			if line[beam] == '^' {
				if newBeams[beam] != true {
					count++
				}
				newBeams[beam-1] = true
				newBeams[beam+1] = true
			} else if line[beam] == '.' {
				newBeams[beam] = true
			}
		}
		beams = newBeams
	}
	return strconv.Itoa(count)
}

func day7Part2(data []string) string {
	startLine := data[0]
	beams := map[int]int{
		strings.Index(startLine, "S"): 1,
	}
	data = data[1:]
	for _, line := range data {
		if line == "" { continue }
		newBeams := map[int]int{}
		for beam, count := range beams {
			if line[beam] == '^' {
				newBeams[beam-1] += count
				newBeams[beam+1] += count
			} else if line[beam] == '.' {
				newBeams[beam] += count
			}
		}
		beams = newBeams
	}
	count := 0
	for _, c := range beams {
		count += c
	}
	return strconv.Itoa(count)
}
