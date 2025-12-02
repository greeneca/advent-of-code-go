package aoc2017

import "strconv"

func day5Part1(data []string) string {
	maze := []int{}
	for _, line := range data {
		if len(line) > 0 {
			value, _ := strconv.Atoi(line)
			maze = append(maze, value)
		}
	}
	pos := 0
	steps := 0
	for pos < len(maze) {
		jump := maze[pos]
		maze[pos]++
		pos += jump
		steps++
	}
	return strconv.Itoa(steps)
}

func day5Part2(data []string) string {
	maze := []int{}
	for _, line := range data {
		if len(line) > 0 {
			value, _ := strconv.Atoi(line)
			maze = append(maze, value)
		}
	}
	pos := 0
	steps := 0
	for pos < len(maze) {
		jump := maze[pos]
		if jump >= 3 {
			maze[pos]--
		} else {
			maze[pos]++
		}
		pos += jump
		steps++
	}
	return strconv.Itoa(steps)
}
