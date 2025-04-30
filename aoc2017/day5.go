package aoc2017

import "strconv"

func day5Part1(data []string){
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
	println("Steps to exit:", steps)
}

func day5Part2(data []string){
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
	println("Steps to exit:", steps)
}
