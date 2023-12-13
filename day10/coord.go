package main

type Coord [2]int

var directionCoords = map[Direction]Coord{
	North: {-1, 0},
	South: {1, 0},
	West:  {0, -1},
	East:  {0, 1},
}

func (c Coord) Add(other Coord) Coord {
	return Coord{c[0] + other[0], c[1] + other[1]}
}

func (c Coord) LessThan(other Coord) bool {
	if c[0] < other[0] {
		return true
	} else if c[0] > other[0] {
		return false
	}
	return c[1] <= other[1]
}
