package aoc2017

import (
	"fmt"
	"strconv"
	"strings"
)

func day6Part1(data []string) {
	steps := 0
	cache := map[string]bool{}
	memmory := []int{}
	for bank := range strings.SplitSeq(data[0], "\t") {
		value, _ := strconv.Atoi(bank)
		memmory = append(memmory, value)
	}
	for {
		steps++
		maxIndex := getMaxMemmoryIndex(memmory)
		value := memmory[maxIndex]
		memmory[maxIndex] = 0
		for i := 1; i <= value; i++ {
			memmory[(maxIndex+i)%len(memmory)]++
		}
		cacheKey := fmt.Sprint(memmory)
		if _, exists := cache[cacheKey]; exists {
			break
		}
		cache[cacheKey] = true
	}
	println("Steps to loop:", steps)
}

func day6Part2(data []string) {
	steps := 0
	cache := map[string]int{}
	memmory := []int{}
	for bank := range strings.SplitSeq(data[0], "\t") {
		value, _ := strconv.Atoi(bank)
		memmory = append(memmory, value)
	}
	for {
		steps++
		maxIndex := getMaxMemmoryIndex(memmory)
		value := memmory[maxIndex]
		memmory[maxIndex] = 0
		for i := 1; i <= value; i++ {
			memmory[(maxIndex+i)%len(memmory)]++
		}
		cacheKey := fmt.Sprint(memmory)
		if _, exists := cache[cacheKey]; exists {
			steps = steps - cache[cacheKey]
			break
		}
		cache[cacheKey] = steps
	}
	println("Size of loop:", steps)
}

func getMaxMemmoryIndex(memmory []int) int {
	maxIndex := 0
	maxValue := memmory[0]
	for i, value := range memmory {
		if value > maxValue {
			maxValue = value
			maxIndex = i
		}
	}
	return maxIndex
}
