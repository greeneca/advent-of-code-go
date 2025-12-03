package aoc2025

import (
	"math"
	"strconv"
	"strings"
)

func day3Part1(data []string) string {
	sum := 0
	for _, bank := range data {
		if bank == "" { continue }
		batteries := strings.Split(bank, "")
		max := 0
		maxIndex := -1
		for i, bat := range batteries[:len(batteries)-1] {
			val, _ := strconv.Atoi(bat)
			if val > max {
				max = val
				maxIndex = i
			}
		}
		sum += max * 10
		max = 0
		for _, bat := range batteries[maxIndex+1:] {
			val, _ := strconv.Atoi(bat)
			if val > max {
				max = val
			}
		}
		sum += max
	}
	return strconv.Itoa(sum)
}

func day3Part2(data []string) string {
	sum := int64(0)
	for _, bank := range data {
		if bank == "" { continue }
		batteries := strings.Split(bank, "")
		start := 0
		for i := range 12 {
			max := 0
			newStart := start
			for j, bat := range batteries[start:len(batteries)-(11 - i)] {
				val, _ := strconv.Atoi(bat)
				if val > max {
					max = val
					newStart = start + j + 1
				}
			}
			jolts := int64(max) * int64(math.Pow(10, float64(11 - i)))
			sum += jolts
			start = newStart
		}
	}
	return strconv.FormatInt(sum, 10)
}
