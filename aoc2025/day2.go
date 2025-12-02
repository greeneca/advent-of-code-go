package aoc2025

import (
	"fmt"
	"strconv"
	"strings"
)

func day2Part1(data []string) {
	sum := 0
	ranges := strings.SplitSeq(data[0], ",")
	for r := range ranges {
		bounds := strings.Split(r, "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])
		for i := start; i <= end; i++ {
			strNum := strconv.Itoa(i)
			length := len(strNum)
			if strNum[0:(length)/2] == strNum[(length)/2:] {
				sum += i
			}
		}
	}
	fmt.Println("Sum is:", sum)
}

func day2Part2(data []string) {
	sum := 0
	ranges := strings.SplitSeq(data[0], ",")
	for r := range ranges {
		bounds := strings.Split(r, "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])
		for i := start; i <= end; i++ {
			strNum := strconv.Itoa(i)
			length := len(strNum)
			for j := 1; j <= length/2; j++ {
				if length%(j) != 0 { continue }
				matching := true
				for k := 0; k <= (length-(2*j)); k+= j {
					a1 := k
					a2 := k + (j)
					b1 := k + j
					b2 := k + (2 * j)

					if strNum[a1:a2] != strNum[b1:b2] {
						matching = false
						break
					}
				}
				if matching {
					sum += i
					break
				}
			}
		}
	}
	fmt.Println("Sum is:", sum)
}
