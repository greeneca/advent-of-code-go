package aoc2017

import (
	"math"
	"strconv"
	"strings"
)

func day11Part1(data []string) string {
	movements := map[string][3]int{
		"n":  {1, 0, -1},
		"ne": {0, 1, -1},
		"se": {-1, 1, 0},
		"s":  {-1, 0, 1},
		"sw": {0, -1, 1},
		"nw": {1, -1, 0},
	}
	s, q, r := 0, 0, 0
	for _, move := range strings.Split(data[0], ",") {
		ds, dq, dr := movements[move][0], movements[move][1], movements[move][2]
		s += ds
		q += dq
		r += dr
	}
	return strconv.Itoa(int(math.Max(math.Max(math.Abs(float64(s)), math.Abs(float64(q))), math.Abs(float64(r)))))
}


func day11Part2(data []string) string {
	movements := map[string][3]int{
		"n":  {1, 0, -1},
		"ne": {0, 1, -1},
		"se": {-1, 1, 0},
		"s":  {-1, 0, 1},
		"sw": {0, -1, 1},
		"nw": {1, -1, 0},
	}
	maxDistance := 0
	s, q, r := 0, 0, 0
	for _, move := range strings.Split(data[0], ",") {
		ds, dq, dr := movements[move][0], movements[move][1], movements[move][2]
		s += ds
		q += dq
		r += dr
		currentDistance := int(math.Max(math.Max(math.Abs(float64(s)), math.Abs(float64(q))), math.Abs(float64(r))))
		if currentDistance > maxDistance {
			maxDistance = currentDistance
		}
	}
	return strconv.Itoa(maxDistance)
}
