package aoc2025

import (
	"slices"
	"sort"
	"strconv"
	"strings"
)

type IngredentRange struct {
	Start int
	End   int
}

func (r *IngredentRange) Count() int {
	return r.End - r.Start + 1
}

func (r *IngredentRange) Includes(id int) bool {
	return id >= r.Start && id <= r.End
}

func (r *IngredentRange) Overlap(r2 IngredentRange) bool {
	return r.Includes(r2.Start) || r.Includes(r2.End) || r2.Includes(r.Start) || r2.Includes(r.End)
}

func MergeRanges(r1, r2 IngredentRange) IngredentRange {
	start := r1.Start
	if r2.Start < start {
		start = r2.Start
	}
	end := r1.End
	if r2.End > end {
		end = r2.End
	}
	return IngredentRange{
		Start: start,
		End:   end,
	}
}

func day5Part1(data []string) string {
	freshRanges := []IngredentRange{}
	in_ranges := true
	count := 0
	for _, line := range data {
		if in_ranges && line == "" {
			in_ranges = false
			continue
		}
		if in_ranges {
			r := strings.Split(line, "-")
			start_range, _ := strconv.Atoi(r[0])
			end_range, _ := strconv.Atoi(r[1])
			freshRanges = append(freshRanges, IngredentRange{
				Start: start_range,
				End: end_range,
			})
		} else {
			id, _ := strconv.Atoi(line)
			for _, r := range freshRanges {
				if r.Includes(id) {
					count++
					break
				}
			}
		}
	}
	return strconv.Itoa(count)
}

func day5Part2(data []string) string {
	freshRanges := []IngredentRange{}
	count := 0
	for _, line := range data {
		if line == "" {
			break
		}
		r := strings.Split(line, "-")
		start_range, _ := strconv.Atoi(r[0])
		end_range, _ := strconv.Atoi(r[1])
		freshRanges = append(freshRanges, IngredentRange{
			Start: start_range,
			End: end_range,
		})
	}
	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i].Start < freshRanges[j].Start
	})
	i := 0
	for i < len(freshRanges)-1 {
		current := freshRanges[i]
		next := freshRanges[i+1]
		if current.Overlap(next) {
			merged := MergeRanges(current, next)
			freshRanges[i] = merged
			freshRanges = slices.Delete(freshRanges, i+1, i+2)
		} else {
			i++
		}
	}
	count = 0
	for _, r := range freshRanges {
		count += r.Count()
	}
	return strconv.Itoa(count)
}
