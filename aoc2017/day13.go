package aoc2017

import (
	"fmt"
	"strconv"
)

func day13Part1(data []string) string {
	layers := []int{}
	for _, line := range data {
		if line == "" { continue }
		var layer, size int
		fmt.Sscanf(line, "%d: %d", &layer, &size)
		for len(layers) <= layer {
			layers = append(layers, 0)
		}
		layers[layer] = size
	}
	severity := 0
	for time, layer := range layers {
		if layer == 0 {
			continue
		}
		if time%(2*(layer-1)) == 0 {
			severity += time * layer
		}
	}
	return strconv.Itoa(severity)
}

func day13Part2(data []string) string {
	layers := []int{}
	for _, line := range data {
		if line == "" { continue }
		var layer, size int
		fmt.Sscanf(line, "%d: %d", &layer, &size)
		for len(layers) <= layer {
			layers = append(layers, 0)
		}
		layers[layer] = size
	}
	delay := 0
	for true {
		passed := true
		for time, layer := range layers {
			if layer == 0 { continue }
			if (time+delay)%(2*(layer-1)) == 0 {
				delay++
				passed = false
				break
			}
		}
		if passed {
			return strconv.Itoa(delay)
		}
	}
	return ""
}
