package aoc2017

import (
	"strconv"

	"github.com/greeneca/advent-of-code-go/vector"
)

const input = 277678

func day3Part1(data []string) string {
	level := 1
	for {
		maxValue := level * level
		add := false
		if maxValue >= input {
			jump := (level-1)/2
			for maxValue > input {
				maxValue -= jump
				add = !add
			}
			shellDistance := input - maxValue
			if !add {
				shellDistance = shellDistance *-1
			}
			distance := jump+shellDistance
			return strconv.Itoa(distance)
		}
		level += 2
	}
}
func day3Part2(data []string) string {
	grid := map[vector.Vector]int{}
	position := vector.New(0, 0)
	direction := 0
	grid[position] = 1
	for {
		position, direction = spiralGridMove(position, &grid, direction)
		sum := 0
		for _, v := range vector.GetAllDirections() {
			if val, ok := grid[position.Add(v)]; ok {
				sum += val
			}
		}
		grid[position] = sum
		if sum > input {
			return strconv.Itoa(sum)
		}
	}
}

func spiralGridMove(position vector.Vector, grid *map[vector.Vector]int, direction int) (vector.Vector, int){
	dirs := vector.GetDirections()
	if position.IsAt(0, 0) {
		position = position.Add(dirs[direction])
		return position, direction
	}
	leftPos := position.Add(dirs[(direction+3)%4])
	if _, ok := (*grid)[leftPos]; !ok {
		direction = (direction + 3) % 4
	}
	position = position.Add(dirs[direction])
	return position, direction
}
