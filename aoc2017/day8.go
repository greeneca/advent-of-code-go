package aoc2017

import (
	"maps"
	"slices"
	"strconv"
	"strings"
)

func day8Part1(data []string) {
	registers := make(map[string]int)
	for _, line := range data {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		reg := parts[0]
		op := parts[1]
		value, _ := strconv.Atoi(parts[2])
		condReg := parts[4]
		condOp := parts[5]
		condValue, _ := strconv.Atoi(parts[6])
		condRegValue := registers[condReg]
		if testCondition(condOp, condRegValue, condValue) {
			if op == "inc" {
				registers[reg] += value
			} else if op == "dec" {
				registers[reg] -= value
			}
		}
	}
	maxValue := slices.Max(slices.Collect(maps.Values(registers)))
	println("Max register value after completion:", maxValue)
}

func testCondition(op string, a int, b int) bool {
	conditionMet := false
	switch op {
	case ">":
		conditionMet = a > b
	case "<":
		conditionMet = a < b
	case ">=":
		conditionMet = a >= b
	case "<=":
		conditionMet = a <= b
	case "==":
		conditionMet = a == b
	case "!=":
		conditionMet = a != b
	}
	return conditionMet
}

func day8Part2(data []string) {
	registers := make(map[string]int)
	maxValue := 0
	for _, line := range data {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		reg := parts[0]
		op := parts[1]
		value, _ := strconv.Atoi(parts[2])
		condReg := parts[4]
		condOp := parts[5]
		condValue, _ := strconv.Atoi(parts[6])
		condRegValue := registers[condReg]
		if testCondition(condOp, condRegValue, condValue) {
			if op == "inc" {
				registers[reg] += value
			} else if op == "dec" {
				registers[reg] -= value
			}
			if registers[reg] > maxValue {
				maxValue = registers[reg]
			}
		}
	}
	println("Max register value during run:", maxValue)
}
