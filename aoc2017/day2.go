package aoc2017

import (
	"strconv"
	"strings"
)

func day2Part1(data []string) string {
	sum := 0
	for i := range len(data) {
		line := string(data[i])
		min := 9999
		max := 0
		values := strings.Split(line, "\t")
		for j := range len(values) {
			value, _ := strconv.Atoi(values[j])
			if value < min {
				min = value
			}
			if value > max {
				max = value
			}
		}
		sum += max - min
	}
	return strconv.Itoa(sum)
}

func day2Part2(data []string) string {
	sum := 0
	for i := range len(data) {
		line := string(data[i])
		values := strings.Split(line, "\t")
		for j := range len(values) {
			toAdd := -1
			value, _ := strconv.Atoi(values[j])
			for k := j+1; k < len(values); k++ {
				value2, _ := strconv.Atoi(values[k])
				if value%value2 == 0 {
					toAdd = value / value2
					break
				}
				if value2%value == 0 {
					toAdd = value2 / value
					break
				}
			}
			if toAdd != -1 {
				sum += toAdd
				break
			}
		}
	}
	return strconv.Itoa(sum)
}
