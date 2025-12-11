package aoc2025

import (
	"fmt"
	"sort"

	"github.com/greeneca/advent-of-code-go/vector"
	"gonum.org/v1/gonum/stat/combin"
)

func day9Part1(data []string) string {
	corners := []vector.Vector{}
	for _, points := range data {
		if points == "" { continue }
		var x, y int
		fmt.Sscanf(points, "%d,%d", &x, &y)
		corners = append(corners, vector.New(x, y))
	}
	pairs := combin.Combinations(len(corners), 2)
	sort.Slice(pairs, func(i, j int) bool {
		a, b := pairs[i], pairs[j]
		sizeA := getRectSize(corners[a[0]], corners[a[1]])
		sizeB := getRectSize(corners[b[0]], corners[b[1]])
		return sizeA > sizeB
	})
	largest := pairs[0]
	return fmt.Sprintf("%d", getRectSize(corners[largest[0]], corners[largest[1]]))
}

func getRectSize(a vector.Vector, b vector.Vector) int {
	x := []int{a.X, b.X}
	sort.Ints(x)
	y := []int{a.Y, b.Y}
	sort.Ints(y)
	return (x[1] - x[0] +1) * (y[1] - y[0]+1)
}

func day9Part2(data []string) string {
	corners := []vector.Vector{}
	cornerMap := map[vector.Vector]bool{}

	for _, points := range data {
		if points == "" { continue }
		var x, y int
		fmt.Sscanf(points, "%d,%d", &x, &y)
		corners = append(corners, vector.New(x, y))
		cornerMap[vector.New(x, y)] = true
	}

	pairs := combin.Combinations(len(corners), 2)
	sizeCache := map[[2]int]int{}
	sort.Slice(pairs, func(i, j int) bool {
		a, b := pairs[i], pairs[j]
		sizeA, exists := sizeCache[[2]int{a[0], a[1]}]
		if !exists {
			sizeA = getRectSizeInside(corners[a[0]], corners[a[1]], corners)
			sizeCache[[2]int{a[0], a[1]}] = sizeA
		}
		sizeB, exists := sizeCache[[2]int{b[0], b[1]}]
		if !exists {
			sizeB = getRectSizeInside(corners[b[0]], corners[b[1]], corners)
			sizeCache[[2]int{b[0], b[1]}] = sizeB
		}
		return sizeA > sizeB
	})
	largest := pairs[0]
	return fmt.Sprintf("%d", sizeCache[[2]int{largest[0], largest[1]}])
}

func getRectSizeInside(a vector.Vector, b vector.Vector, corners []vector.Vector) int {
	x := []int{a.X, b.X}
	sort.Ints(x)
	y := []int{a.Y, b.Y}
	sort.Ints(y)

	rect := []vector.Vector{}
	rect = append(rect, vector.New(x[0], y[0]))
	rect = append(rect, vector.New(x[1], y[0]))
	rect = append(rect, vector.New(x[1], y[1]))
	rect = append(rect, vector.New(x[0], y[1]))
	for _, point := range rect {
		if !isValidPoint(corners, point) {
			return 0
		}
	}
	for i, a := range rect {
		b := rect[(i+1)%len(rect)]
		if sampleEdge(a, b, corners) == false {
			return 0
		}
	}
	size := (x[1] - x[0] + 1) * (y[1] - y[0] + 1)
	return size
}

func sampleEdge(a, b vector.Vector, corners []vector.Vector) bool {
	samples := 75
	sampleSize := 0
	if vertical(a, b) {
		diff := a.Y - b.Y
		sampleSize = max(diff, (diff*-1))/samples
	} else if horizontal(a, b) {
		diff := a.X - b.X
		sampleSize = max(diff, (diff*-1))/samples
	}
	for i := 0; i <= samples; i++ {
		var point vector.Vector
		if vertical(a, b) {
			newY := min(a.Y, b.Y) + i*sampleSize
			point = vector.New(a.X, newY)
		} else if horizontal(a, b) {
			newX := min(a.X, b.X) + i*sampleSize
			point = vector.New(newX, a.Y)
		}
		if !isValidPoint(corners, point) {
			return false
		}
	}
	return true
}

func isValidPoint(corners []vector.Vector, point vector.Vector) bool {
	return isOnBoundry(corners, point) || isInside(corners, point)
}

func isOnBoundry(corners []vector.Vector, point vector.Vector) bool {
	for i, corner := range corners {
		nextCorner := corners[(i+1)%len(corners)]
		if isPointOnLine(corner, nextCorner, point) {
			return true
		}
	}
	return false
}

func isPointOnLine(a, b, p vector.Vector) bool {
	if (((a.Y > p.Y) != (b.Y > p.Y)) && p.X == a.X) || (((a.X > p.X) != (b.X > p.X)) && p.Y == a.Y) {
		return true
	}
	return false
}

func isInside(corners []vector.Vector, point vector.Vector) bool {
	count := 0
	for i, corner := range corners {
		nextCorner := corners[(i+1)%len(corners)]
		if rayIntersectsSegment(point, corner, nextCorner) {
			count++
		}
	}
	return count%2 == 1
}

func rayIntersectsSegment(p, a, b vector.Vector) bool {
	if ((a.Y > p.Y) != (b.Y > p.Y)) && (p.X > a.X) {
		return true
	}
	return false
}

func vertical(a, b vector.Vector) bool {
	return a.Y != b.Y && a.X == b.X
}
func horizontal(a, b vector.Vector) bool {
	return a.Y == b.Y && a.X != b.X
}

