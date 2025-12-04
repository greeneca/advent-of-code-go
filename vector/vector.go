package vector

import "strconv"

type Vector struct {
	X, Y int
}

func New(x, y int) Vector {
	return Vector{X: x, Y: y}
}

func (v Vector) Add(other Vector) Vector {
	return Vector{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v Vector) IsAt(x, y int) bool {
	return v.X == x && v.Y == y
}

func (v Vector) ForEachNeighbour(function func(Vector)) {
	for _, dir := range GetAllDirections() {
		neighbour := v.Add(dir)
		function(neighbour)
	}
}

func (v Vector) GetMatchingNeighbours(function func(Vector) bool) []Vector {
	matching := []Vector{}
	for _, dir := range GetAllDirections() {
		neighbour := v.Add(dir)
		if function(neighbour) {
			matching = append(matching, neighbour)
		}
	}
	return matching
}

func (v Vector) Print() {
	println(v.ToString())
}

func (v Vector) ToString() string {
	return "X: " + strconv.Itoa(v.X) + " Y: " + strconv.Itoa(v.Y)
}

func GetDirections() []Vector {
	return []Vector{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
		{X: 0, Y: -1},
	}
}

func GetAllDirections() []Vector {
	return []Vector{
		{X: 1, Y: 0},
		{X: 1, Y: 1},
		{X: 0, Y: 1},
		{X: -1, Y: 1},
		{X: -1, Y: 0},
		{X: -1, Y: -1},
		{X: 0, Y: -1},
		{X: 1, Y: -1},
	}
}
