package aoc2025

import (
	"regexp"
	"strconv"
	"strings"
)

type Problem struct {
	Index  int
	Values []int
	Op string
}

func (p *Problem) Solve() int {
	switch p.Op {
	case "*":
		prod := 1
		for _, v := range p.Values {
			prod *= v
		}
		return prod
	case "+":
		sum := 0
		for _, v := range p.Values {
			sum += v
		}
		return sum
	default:
		return 0
	}
}

func day6Part1(data []string) string {
	problems := []*Problem{}
	re := regexp.MustCompile(`\s+`)
	sum := 0
	for _, line := range data {
		inputs := re.Split(line, -1)
		if len(problems) == 0 {
			for range inputs {
				problems = append(problems, &Problem{})
			}
		}
		index := -1
		for _, input := range inputs {
			if input == "" { continue }
			index++
			p := problems[index]
			val, err := strconv.Atoi(input)
			if err != nil {
				p.Op = input
				sum += p.Solve()
			} else {
				p.Values = append(p.Values, val)
			}
		}
	}
	return strconv.Itoa(sum)
}

func day6Part2(data []string) string {
	problems := []*Problem{}
	sum := 0
	opLine := data[len(data)-2]
	data = data[:len(data)-2]
	for idx, runeVal := range opLine {
		if runeVal == ' ' { continue }
		p := &Problem{Index: idx}
		p.Op = string(runeVal)
		problems = append(problems, p)
	}
	for idx, p := range problems {
		end := len(data[0])-1
		if len(problems) >= idx+2 {
			end = problems[idx+1].Index-1
		}
		i := p.Index - 1
		for i < end {
			i++
			valStr := ""
			for _, line := range data {
				valStr += string(line[i])
			}
			val, err := strconv.Atoi(strings.Trim(valStr, " "))
			if err != nil {
				continue
			}
			p.Values = append(p.Values, val)
		}
		sum += p.Solve()
	}
	return strconv.Itoa(sum)
}
