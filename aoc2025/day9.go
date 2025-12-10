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
	print(corners, largest)
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
	for i, pointA := range rect {
		pointB := rect[(i+1)%len(rect)]
		if crossesShape(corners, pointA, pointB) {
			return 0
		}
	}
	size := (x[1] - x[0] + 1) * (y[1] - y[0] + 1)
	return size
}

func crossesShape(corners []vector.Vector, pointA vector.Vector, pointB vector.Vector) bool {
	for i, pointC := range corners {
		pointD := corners[(i+1)%len(corners)]
		if linesCross(pointA, pointB, pointC, pointD) {
			return true
		}
	}
	return false
}

func linesCross(a1, a2, b1, b2 vector.Vector) bool {
	ax := []int{a1.X, a2.X}
	sort.Ints(ax)
	ay := []int{a1.Y, a2.Y}
	sort.Ints(ay)
	bx := []int{b1.X, b2.X}
	sort.Ints(bx)
	by := []int{b1.Y, b2.Y}
	sort.Ints(by)
	if ax[0] == ax[1] {
		if bx[1] >= ax[1] && bx[0] <= ax[1] {
			if ay[1] >= by[0] && ay[0] <= by[1] {
				return true
			}
		}
	} else if ay[0] == ay[1] {
		if by[1] >= ay[1] && by[0] <= ay[1] {
			if ax[1] >= bx[0] && ax[0] <= bx[1] {
				return true
			}
		}
	}
	return false
}


func print(corners []vector.Vector, largestPair []int) {
	minX, minY := corners[0].X, corners[0].Y
	maxX, maxY := corners[0].X, corners[0].Y
	cornersMap := map[vector.Vector]bool{}
	for _, corner := range corners {
		if corner.X < minX { minX = corner.X }
		if corner.X > maxX { maxX = corner.X }
		if corner.Y < minY { minY = corner.Y }
		if corner.Y > maxY { maxY = corner.Y }
		cornersMap[corner] = true
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			p := vector.New(x, y)
			if x == corners[largestPair[0]].X && y == corners[largestPair[0]].Y {
				fmt.Print("A")
			} else if x == corners[largestPair[1]].X && y == corners[largestPair[1]].Y {
				fmt.Print("A")
			} else if cornersMap[p] {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
