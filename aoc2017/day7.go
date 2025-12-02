package aoc2017

import (
	"maps"
	"slices"
	"strconv"
	"strings"
)

func day7Part1(data []string) string {
	programs := []string{}
	holding := []string{}
	for _, line := range data {
		parts := strings.Split(line, " ")
		programs = append(programs, parts[0])
		if len(parts) > 2 {
			for _, held := range parts[3:] {
				holding = append(holding, strings.ReplaceAll(held, ",", ""))
			}
		}
	}
	for _, program := range programs {
		if !slices.Contains(holding, program) {
			return program
		}
	}
	return ""
}

type Node struct {
	Weight int
	Children []string
	ChildrenWeight int
}

func day7Part2(data []string) string {
	tree := map[string]Node{}
	tower := []string{}
	for _, line := range data {
		parts := strings.Split(line, " ")
		name := parts[0]
		weight := 0
		children := []string{}
		if len(parts) > 1 {
			w, err := strconv.Atoi(strings.Trim(parts[1], "()"))
			if err != nil {
				println("Error parsing weight:", parts[1], "in line:", line)
				continue
			}
			weight = w
			if len(parts) > 2 {
				for _, held := range parts[3:] {
					held = strings.TrimSuffix(held, ",")
					children = append(children, held)
					tower = append(tower, held)
				}
			}
		}
		if name != "" {
			tree[name] = Node{Weight: weight, Children: children}
		}
	}
	for program := range maps.Keys(tree) {
		if !slices.Contains(tower, program) {
			println("Root program for balancing:", program)
			// root program found
			weightDiff := 0
			for true {
				node := tree[program]
				childWeights := map[int][]string{}
				for _, child := range node.Children {
					weight := calculatedWeight(child, tree)
					childWeights[weight] = append(childWeights[weight], child)
				}
				if (len(childWeights) == 1 && weightDiff != 0) {
					correctedWeight := tree[program].Weight + weightDiff
					return strconv.Itoa(correctedWeight)
				} else {
					var correctWeight, incorrectWeight int
					for weight, children := range childWeights {
						if len(children) == 1 {
							incorrectWeight = weight
							program = children[0]
						} else {
							correctWeight = weight
						}
					}
					weightDiff = correctWeight - incorrectWeight
				}
			}
		}
	}
	return ""
}

func calculatedWeight(program string, tree map[string]Node) int {
	node := tree[program]
	totalWeight := node.Weight
	if node.ChildrenWeight != 0 {
		return totalWeight + node.ChildrenWeight
	}
	childrenWeight := 0
	for _, child := range node.Children {
		childWeight := calculatedWeight(child, tree)
		childrenWeight += childWeight
		totalWeight += childWeight
	}
	node.ChildrenWeight = childrenWeight
	tree[program] = node

	return totalWeight
}
