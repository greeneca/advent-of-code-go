package aoc2017

import (
	"strconv"
	"strings"
)

func day12Part1(data []string) string {
	houses := make(map[int][]int)
	for _, line := range data {
		if line == "" { continue }
		parts := strings.SplitN(line, " <-> ", 2)
		id, _ := strconv.Atoi(parts[0])
		connections := strings.SplitSeq(parts[1], ", ")
		for conn := range connections {
			connID, _ := strconv.Atoi(conn)
			houses[id] = append(houses[id], connID)
		}
	}
	to_visit := []int{0}
	seen := make(map[int]bool)
	for len(to_visit) > 0 {
		current := to_visit[0]
		to_visit = to_visit[1:]
		seen[current] = true
		for _, neighbor := range houses[current] {
			if !seen[neighbor] {
				to_visit = append(to_visit, neighbor)
			}
		}
	}
	return strconv.Itoa(len(seen))
}

func day12Part2(data []string) string {
	houses := make(map[int][]int)
	for _, line := range data {
		if line == "" { continue }
		parts := strings.SplitN(line, " <-> ", 2)
		id, _ := strconv.Atoi(parts[0])
		connections := strings.SplitSeq(parts[1], ", ")
		for conn := range connections {
			connID, _ := strconv.Atoi(conn)
			houses[id] = append(houses[id], connID)
		}
	}
	groups := []map[int]bool{}
	for len(houses) > 0 {
		for start := range houses {
			newHouses, group := getGroup(houses, start)
			groups = append(groups, group)
			houses = newHouses
			break
		}
	}
	return strconv.Itoa(len(groups))
}

func getGroup(houses map[int][]int, start int) (map[int][]int, map[int]bool) {
	to_visit := []int{start}
	seen := make(map[int]bool)
	for len(to_visit) > 0 {
		current := to_visit[0]
		to_visit = to_visit[1:]
		seen[current] = true
		for _, neighbor := range houses[current] {
			if !seen[neighbor] {
				to_visit = append(to_visit, neighbor)
			}
		}
		delete(houses, current)
	}
	return houses, seen
}
