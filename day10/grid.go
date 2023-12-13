package main

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/iyassou/advent-of-code-2023/internal"
)

type Grid struct {
	Tiles [][]*Tile
	Start Coord
}

func NewGrid(input []byte) (*Grid, error) {
	if input == nil {
		return nil, errors.New("cannot make grid from empty input")
	}
	lines := internal.Lines(input)
	g := &Grid{Tiles: make([][]*Tile, len(lines))}
	for i, line := range lines {
		g.Tiles[i] = make([]*Tile, len(line))
		for j, b := range line {
			g.Tiles[i][j] = NewTile(b)
			if g.Tiles[i][j].Type == Start {
				g.Start = Coord([2]int{i, j})
			}
		}
	}
	return g, nil
}

func (g *Grid) String() string {
	if g.Tiles == nil {
		return ""
	}
	var sb strings.Builder
	for _, line := range g.Tiles {
		for _, tile := range line {
			sb.WriteByte(tile.Char)
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (g *Grid) GetMainLoop() []Coord {
	path := []Coord{}
	visited := func(c Coord) bool {
		for _, v := range path {
			if c == v {
				return true
			}
		}
		return false
	}
	current := g.Start
	for !visited(current) {
		path = append(path, current)
		for _, c := range g.canTravelTo(current) {
			if !visited(c) {
				current = c
				break
			}
		}
	}
	return path
}

func (g *Grid) LoopInterior(loop []Coord) []Coord {
	onLoop := func(i, j int) bool {
		for _, p := range loop {
			if p[0] == i && p[1] == j {
				return true
			}
		}
		return false
	}
	isInteriorPoint := func(i, j int) bool {
		// ASSUMPTION: (i,j) is not on the loop
		notLoop := true
		togglesLeft := 0
		for y := 0; y < j; y++ {
			if notLoop == onLoop(i, y) {
				togglesLeft++
				notLoop = !notLoop
			}
		}
		if togglesLeft == 0 {
			return false
		}
		if !notLoop {
			togglesLeft++
		}
		togglesRight := 0
		for y := j + 1; y < len(g.Tiles[i]); y++ {
			if notLoop == onLoop(i, y) {
				togglesRight++
				notLoop = !notLoop
			}
		}
		if togglesRight == 0 {
			return false
		}
		if !notLoop {
			togglesRight++
		}
		toggles := togglesLeft + togglesRight
		return toggles%2 == 0
	}
	topleft, bottomright := g.bbox(loop)
	interior := []Coord{}
	for i := topleft[0]; i <= bottomright[0]; i++ {
		for j := topleft[1]; j <= bottomright[1]; j++ {
			if !onLoop(i, j) && isInteriorPoint(i, j) {
				interior = append(interior, Coord{i, j})
			}
		}
	}
	return interior
}

func (g *Grid) getTile(c Coord) (*Tile, error) {
	x, y := c[0], c[1]
	if !(0 <= x && x < len(g.Tiles)) {
		return nil, fmt.Errorf("invalid x coordinate %d", x)
	}
	if !(0 <= y && y < len(g.Tiles[x])) {
		return nil, fmt.Errorf("invalid y coordinate %d", y)
	}
	return g.Tiles[x][y], nil
}

func (g *Grid) canTravelTo(c Coord) []Coord {
	coords := []Coord{}
	t, err := g.getTile(c)
	if err != nil {
		return coords
	}
	for dir := range directionCoords {
		if !t.CanFitIn(dir) {
			continue
		}
		offset := directionCoords[dir.Opposite()]
		neighbour := c.Add(offset)
		if ot, err := g.getTile(neighbour); err != nil {
			continue
		} else if ot.CanFitIn(dir.Opposite()) {
			coords = append(coords, neighbour)
		}
	}
	return coords
}

func (g *Grid) highlightPath(path []Coord) string {
	if path == nil {
		return ""
	}
	inPath := func(i, j int) bool {
		c := Coord{i, j}
		for _, node := range path {
			if c == node {
				return true
			}
		}
		return false
	}
	var sb strings.Builder
	for i, line := range g.Tiles {
		for j, tile := range line {
			if inPath(i, j) {
				sb.WriteByte(tile.Char)
			} else {
				sb.WriteByte(' ')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (g *Grid) bbox(points []Coord) (Coord, Coord) {
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	for _, p := range points {
		x, y := p[0], p[1]
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}
	return Coord{minX, minY}, Coord{maxX, maxY}
}
