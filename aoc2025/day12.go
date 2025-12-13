package aoc2025

import (
	"fmt"

	"github.com/greeneca/advent-of-code-go/vector"
)

type box struct {
	xSize int
	ySize int
	contents map[vector.Vector]bool
}

func newBox(xSize, ySize int) *box {
	return &box{
		xSize: xSize,
		ySize: ySize,
		contents: map[vector.Vector]bool{},
	}
}

func (b *box) canPlace(presentShape map[vector.Vector]bool, at vector.Vector) bool {
	for vec := range presentShape {
		placeVec := vec.Add(at)
		if placeVec.X < 0 || placeVec.X >= b.xSize || placeVec.Y < 0 || placeVec.Y >= b.ySize {
			return false
		}
		if b.contents[placeVec] {
			return false
		}
	}
	return true
}

func (b *box) place(presentShape map[vector.Vector]bool, at vector.Vector) {
	for vec := range presentShape {
		placeVec := vec.Add(at)
		b.contents[placeVec] = true
	}
}

func copyBox(b *box) *box {
	newBox := newBox(b.xSize, b.ySize)
	for vec := range b.contents {
		newBox.contents[vec] = true
	}
	return newBox
}

type present struct {
	index int
	shape map[vector.Vector]bool
	orentations map[string]map[vector.Vector]bool
}

func newPresent(index int, shape map[vector.Vector]bool) *present {
	p := &present{
		index: index,
		shape: shape,
	}
	p.generateOrentations()
	return p
}

func (p *present) generateOrentations() {
	p.orentations = map[string]map[vector.Vector]bool{}
	shape := p.shape
	p.orentations[p.shapeKey(shape)] = shape
	for range 4 {
		shape = rotateShape(shape)
		p.orentations[p.shapeKey(shape)] = shape
	}
	shape = reflectShape(p.shape)
	p.orentations[p.shapeKey(shape)] = shape
	for range 4 {
		shape = rotateShape(shape)
		p.orentations[p.shapeKey(shape)] = shape
	}
}

func (p *present) shapeKey(shape map[vector.Vector]bool) string {
	key := ""
	for y := range 3 {
		for x := range 3 {
			if shape[vector.New(x,y)] {
				key += "#"
			} else {
				key += "."
			}
		}
	}
	return key
}

func (p *present) toString() string {
	str := ""
	for y := range 3 {
		for x := range 3 {
			if p.shape[vector.New(x,y)] {
				str += "#"
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str
}

func rotateShape(shape map[vector.Vector]bool) map[vector.Vector]bool {
	pivot := vector.New(1,1)
	newShape := map[vector.Vector]bool{}
	for vec := range shape {
		newVec := vec.Subtract(pivot)
	    newVec.Rotate90()
		newVec = newVec.Add(pivot)
		newShape[newVec] = true
	}
	return newShape
}

func reflectShape(shape map[vector.Vector]bool) map[vector.Vector]bool {
	pivot := vector.New(1,1)
	newShape := map[vector.Vector]bool{}
	for vec := range shape {
		newVec := vec.Subtract(pivot)
	    newVec.X = -newVec.X
		newVec = newVec.Add(pivot)
		newShape[newVec] = true
	}
	return newShape
}


func day12Part1(data []string) string {
	presents := []*present{}
	for i := range 6 {
		shape := map[vector.Vector]bool{}
		for y := range 3 {
			for x := range 3 {
				if data[i*5+y+1][x] == '#' {
					shape[vector.New(x,y)] = true
				}
			}
		}
		presents = append(presents, newPresent(i, shape))
	}
	data = data[6*5:]
	sum := 0
	for _, line := range data {
		if line == "" { continue }
		var xSize, ySize, num0, num1, num2, num3, num4, num5 int
		fmt.Sscanf(line, "%dx%d: %d %d %d %d %d %d", &xSize, &ySize, &num0, &num1, &num2, &num3, &num4, &num5)
		region := newBox(xSize, ySize)
		presentCounts := []int{num0, num1, num2, num3, num4, num5}
		if presentsFit(region, presents, presentCounts) {
			sum++
		}
	}
	return fmt.Sprintf("%d", sum)
}

func presentsFit(region *box, presents []*present, presentCounts []int) bool {
	lowerBound := 0
	for i, count := range presentCounts {
		lowerBound += count * len(presents[i].shape)
	}
	if lowerBound > region.xSize * region.ySize {
		return false
	}
	if sum(presentCounts)*9 <= region.xSize * region.ySize {
		return true
	}
	println("Need to do full search - unimplemented")
	return false
}

func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}

func day12Part2(data []string) string {
	// Implementation for Day X Part 2
	return "Unimplemented"
}
