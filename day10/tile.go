package main

import "fmt"

type TileType int
type Tile struct {
	Char byte
	Type TileType
}

const (
	UnknownTile TileType = iota
	NorthSouth
	WestEast
	NorthEast
	NorthWest
	SouthWest
	SouthEast
	Ground
	Start
)

var char2tiletype = map[byte]TileType{
	'|': NorthSouth,
	'-': WestEast,
	'L': NorthEast,
	'J': NorthWest,
	'7': SouthWest,
	'F': SouthEast,
	'.': Ground,
	'S': Start,
}

type Direction int

const (
	UnknownDirection Direction = iota
	North
	South
	West
	East
)

func (d Direction) Opposite() Direction {
	if d == North {
		return South
	}
	if d == South {
		return North
	}
	if d == West {
		return East
	}
	if d == East {
		return West
	}
	return UnknownDirection
}

func (d Direction) String() string {
	if d == North {
		return "North"
	}
	if d == South {
		return "South"
	}
	if d == West {
		return "West"
	}
	if d == East {
		return "East"
	}
	return "?"
}

func NewTile(b byte) *Tile {
	t := &Tile{Char: b}
	if tt, ok := char2tiletype[b]; ok {
		t.Type = tt
	} else {
		t.Type = UnknownTile
	}
	return t
}

func (t *Tile) CanFitIn(dir Direction) bool {
	if t == nil {
		return false
	}
	tt := t.Type
	if tt == Ground {
		return false
	}
	if tt == Start {
		return true
	}
	switch dir {
	case North:
		return tt == SouthWest || tt == NorthSouth || tt == SouthEast
	case South:
		return tt == NorthWest || tt == NorthSouth || tt == NorthEast
	case West:
		return tt == NorthEast || tt == WestEast || tt == SouthEast
	case East:
		return tt == NorthWest || tt == WestEast || tt == SouthWest
	}
	return false
}

func (t *Tile) String() string {
	if t == nil {
		return "Tile[]"
	}
	return fmt.Sprintf("Tile[%s]", string(t.Char))
}
