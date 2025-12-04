package aoc2025

import (
	"strconv"

	"github.com/greeneca/advent-of-code-go/vector"
)

func day4Part1(data []string) string {
	paper := make(map[vector.Vector]bool)
	for y, line := range data {
		for x, char := range line {
			if char == '@' {
				paper[vector.Vector{X: x, Y: y}] = true
			}
		}
	}
	count := 0
	for vec := range paper {
		neighbours := vec.GetMatchingNeighbours(func(v vector.Vector) bool {
			_, exists := paper[v]
			return exists
		})
		if len(neighbours) < 4 {
			count++
		}
	}
	return strconv.Itoa(count)
}

func day4Part2(data []string) string {
	paper := make(map[vector.Vector]bool)
	for y, line := range data {
		for x, char := range line {
			if char == '@' {
				paper[vector.Vector{X: x, Y: y}] = true
			}
		}
	}
	to_remove := make([]vector.Vector, 0)
	for vec := range paper {
		neighbours := vec.GetMatchingNeighbours(func(v vector.Vector) bool {
			_, exists := paper[v]
			return exists
		})
		if len(neighbours) < 4 {
			to_remove = append(to_remove, vec)
		}
	}
	count := 0
	for len(to_remove) > 0 {
		vec := to_remove[0]
		to_remove = to_remove[1:]
		if _, exists := paper[vec]; exists {
			delete(paper, vec)
			count++
		}
		vec.ForEachNeighbour(func(v vector.Vector) {
			if _, exists := paper[v]; exists {
				neighbors := v.GetMatchingNeighbours(func(v2 vector.Vector) bool {
					_, exists2 := paper[v2]
					return exists2
				})
				if len(neighbors) < 4 {
					to_remove = append(to_remove, v)
				}
			}
		})
	}
	return strconv.Itoa(count)
}
