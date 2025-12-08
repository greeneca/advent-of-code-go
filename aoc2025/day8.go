package aoc2025

import (
	"fmt"
	"math"
	"sort"

	"gonum.org/v1/gonum/stat/combin"
)

type Circuit struct {
	Indexes []int
}

func day8Part1(data []string) string {
	maxTimes := 1000
	boxes := [][3]int{}
	for _, line := range data {
		if line == "" { continue }
		var l, w, h int
		fmt.Sscanf(line, "%d,%d,%d", &l, &w, &h)
		boxes = append(boxes, [3]int{l, w, h})
	}
	pairs := combin.Combinations(len(boxes), 2)
	sort.Slice(pairs, func(i, j int) bool {
		a, b := pairs[i], pairs[j]
		distA := getPairDistance(boxes[a[0]], boxes[a[1]])
		distB := getPairDistance(boxes[b[0]], boxes[b[1]])
		return distA < distB
	})

	circuits := map[int]*Circuit{}
	circuitSet := map[*Circuit]bool{}
	times := 0
	for _, pair := range pairs {
		if times >= maxTimes { break }
		a, b := pair[0], pair[1]
		aCircuit, aExists := circuits[a]
		bCircuit, bExists := circuits[b]
		if !aExists && !bExists {
			newCircuit := &Circuit{Indexes: []int{a, b}}
			circuitSet[newCircuit] = true
			circuits[a] = newCircuit
			circuits[b] = newCircuit
		} else if aExists && !bExists {
			aCircuit.Indexes = append(aCircuit.Indexes, b)
			circuits[b] = aCircuit
		} else if !aExists && bExists {
			bCircuit.Indexes = append(bCircuit.Indexes, a)
			circuits[a] = bCircuit
		} else if aCircuit != bCircuit {
			// Merge circuits
			aCircuit.Indexes = append(aCircuit.Indexes, bCircuit.Indexes...)
			for _, index := range bCircuit.Indexes {
				circuits[index] = aCircuit
			}
			delete(circuitSet, bCircuit)
		}
		times++
	}
	circuitsList := make([]*Circuit, len(circuitSet))
	i := 0
	for circuit := range circuitSet {
		circuitsList[i] = circuit
		i++
	}
	sort.Slice(circuitsList, func(i, j int) bool {
		return len(circuitsList[i].Indexes) > len(circuitsList[j].Indexes)
	})
	prod := 1
	for i := 0; i < 3 && i < len(circuitsList); i++ {
		prod *= len(circuitsList[i].Indexes)
	}
	return fmt.Sprintf("%d", prod)
}

func getPairDistance(a, b [3]int) float64 {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	dz := a[2] - b[2]
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

func day8Part2(data []string) string {
	boxes := [][3]int{}
	for _, line := range data {
		if line == "" { continue }
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		boxes = append(boxes, [3]int{x, y, z})
	}
	pairs := combin.Combinations(len(boxes), 2)
	sort.Slice(pairs, func(i, j int) bool {
		a, b := pairs[i], pairs[j]
		distA := getPairDistance(boxes[a[0]], boxes[a[1]])
		distB := getPairDistance(boxes[b[0]], boxes[b[1]])
		return distA < distB
	})

	circuits := map[int]*Circuit{}
	circuitSet := map[*Circuit]bool{}
	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		aCircuit, aExists := circuits[a]
		bCircuit, bExists := circuits[b]
		if !aExists && !bExists {
			newCircuit := &Circuit{Indexes: []int{a, b}}
			circuitSet[newCircuit] = true
			circuits[a] = newCircuit
			circuits[b] = newCircuit
		} else if aExists && !bExists {
			aCircuit.Indexes = append(aCircuit.Indexes, b)
			circuits[b] = aCircuit
		} else if !aExists && bExists {
			bCircuit.Indexes = append(bCircuit.Indexes, a)
			circuits[a] = bCircuit
		} else if aCircuit != bCircuit {
			// Merge circuits
			aCircuit.Indexes = append(aCircuit.Indexes, bCircuit.Indexes...)
			for _, index := range bCircuit.Indexes {
				circuits[index] = aCircuit
			}
			delete(circuitSet, bCircuit)
		}
		if len(circuitSet) == 1 && len(circuits) == len(boxes) {
			prod := boxes[a][0] * boxes[b][0]
			return fmt.Sprintf("%d", prod)
		}
	}
	println("No full circuit found", len(circuitSet), len(circuits), len(boxes))
	return "Failed"
}
