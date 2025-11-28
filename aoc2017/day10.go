package aoc2017

import (
	"fmt"
	"strconv"
	"strings"
)

func day10Part1(data []string) {
	pos := 0
	skip := 0
	loop := make([]int, 256)
	for i := range 256 {
		loop[i] = i
	}
	lengths := strings.SplitSeq(data[0], ",")
	for l := range lengths {
		length, err := strconv.Atoi(strings.TrimSpace(l))
		if err != nil {
			continue
		}
		pos, skip = hashRound(length, pos, skip, &loop)
	}
	result := loop[0] * loop[1]
	fmt.Println("Result:", result)
}

func hashRound(length int, pos int, skip int, loop *[]int) (int, int) {
	for i := range length/2 {
		a := (pos + i) % len(*loop)
		b := (pos + length - 1 - i) % len(*loop)
		(*loop)[a], (*loop)[b] = (*loop)[b], (*loop)[a]
	}
	pos = (pos + length + skip) % len(*loop)
	skip++
	return pos, skip
}

func day10Part2(data []string) {
	pos := 0
	skip := 0
	loop := make([]int, 256)
	for i := range 256 {
		loop[i] = i
	}
	bytes := []byte(data[0])
	suffix := []int{17, 31, 73, 47, 23}
	lengths := []int{}
	for _, b := range bytes {
		lengths = append(lengths, int(b))
	}
	lengths = append(lengths, suffix...)
	for range 64 {
		for _, length := range lengths {
			pos, skip = hashRound(length, pos, skip, &loop)
		}
	}
	denseHash := []int{}
	for i := range 16 {
		block := 0
		for j := range 16 {
			block ^= loop[i*16+j]
		}
		denseHash = append(denseHash, block)
	}
	hashString := ""
	for _, num := range denseHash {
		hashString += fmt.Sprintf("%02x", num)
	}
	fmt.Println("Knot Hash:", hashString)
}
