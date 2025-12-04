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
		neighbors := 0
		for _, dir := range vector.GetAllDirections() {
			neighbor := vec.Add(dir)
			if _, exists := paper[neighbor]; exists {
				neighbors++
				if neighbors >= 4 {
					break
				}
			}
		}
		if neighbors < 4 {
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
		neighbors := 0
		for _, dir := range vector.GetAllDirections() {
			neighbor := vec.Add(dir)
			if _, exists := paper[neighbor]; exists {
				neighbors++
				if neighbors >= 4 {
					break
				}
			}
		}
		if neighbors < 4 {
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
		for _, dir := range vector.GetAllDirections() {
			neighbor := vec.Add(dir)
			if _, exists := paper[neighbor]; exists {
				neighbors := 0
				for _, dir2 := range vector.GetAllDirections() {
					neighbor2 := neighbor.Add(dir2)
					if _, exists2 := paper[neighbor2]; exists2 {
						neighbors++
						if neighbors >= 4 {
							break
						}
					}
				}
				if neighbors < 4 {
					to_remove = append(to_remove, neighbor)
				}
			}
		}
	}
	return strconv.Itoa(count)
}
