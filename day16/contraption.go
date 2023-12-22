package main

import (
	"fmt"

	"github.com/iyassou/advent-of-code-2023/internal"
)

type Direction int

const (
	Right Direction = iota
	Left
	Down
	Up
)

func (d Direction) Move(x, y int) (int, int, error) {
	switch d {
	case Right:
		return x + 1, y, nil
	case Left:
		return x - 1, y, nil
	case Down:
		return x, y + 1, nil
	case Up:
		return x, y - 1, nil
	default:
		return 0, 0, fmt.Errorf("invalid direction %v", d)
	}
}

type TileType byte

const (
	EmptySpace TileType = iota
	ForwardSlash
	Backslash
	VerticalSplitter
	HorizontalSplitter
)

func NewTileType(b byte) (TileType, error) {
	switch b {
	case '.':
		return EmptySpace, nil
	case '/':
		return ForwardSlash, nil
	case '\\':
		return Backslash, nil
	case '|':
		return VerticalSplitter, nil
	case '-':
		return HorizontalSplitter, nil
	default:
		return 0, fmt.Errorf("invalid tile %v", b)
	}
}

var reflections = map[TileType]map[Direction]Direction{
	ForwardSlash: {Up: Right, Down: Left, Left: Down, Right: Up},
	Backslash:    {Up: Left, Down: Right, Left: Up, Right: Down},
}

type Tile struct {
	Type   TileType
	Visits []Direction
}

type Contraption struct {
	Tiles [][]*Tile
}

func NewContraption(input []byte) (*Contraption, error) {
	lines := internal.Lines(input)
	c := &Contraption{Tiles: make([][]*Tile, len(lines))}
	for i, line := range lines {
		c.Tiles[i] = make([]*Tile, len(line))
		for j, b := range line {
			if tt, err := NewTileType(b); err != nil {
				return nil, err
			} else {
				c.Tiles[i][j] = &Tile{
					Type:   tt,
					Visits: nil,
				}
			}
		}
	}
	return c, nil
}

func (c *Contraption) resetTileVisits() {
	if c == nil {
		return
	}
	for j := 0; j < c.Height(); j++ {
		for i := 0; i < c.Width(); i++ {
			c.Tiles[j][i].Visits = []Direction{}
		}
	}
}

func (c *Contraption) Width() int {
	if c == nil {
		return 0
	}
	return len(c.Tiles[0])
}

func (c *Contraption) Height() int {
	if c == nil {
		return 0
	}
	return len(c.Tiles)
}

func (c *Contraption) validCoordinates(x, y int) bool {
	if c == nil {
		return false
	}
	return 0 <= x && x < c.Width() && 0 <= y && y < c.Height()
}

func (c *Contraption) ShineBeam(x, y int, d Direction) error {
	if c == nil {
		return nil
	}
	if !c.validCoordinates(x, y) {
		return fmt.Errorf("reuf")
	}
	tile := c.Tiles[y][x]
	for _, visited := range tile.Visits {
		if visited == d {
			return nil
		}
	}
	tile.Visits = append(tile.Visits, d)
	switch tt := tile.Type; tt {
	case EmptySpace:
		if x, y, err := d.Move(x, y); err != nil {
			return err
		} else if !c.validCoordinates(x, y) {
			return nil
		} else {
			return c.ShineBeam(x, y, d)
		}
	case ForwardSlash, Backslash:
		if reflection, ok := reflections[tt][d]; !ok {
			return fmt.Errorf("bruh")
		} else if x, y, err := reflection.Move(x, y); err != nil {
			return err
		} else if !c.validCoordinates(x, y) {
			return nil
		} else {
			return c.ShineBeam(x, y, reflection)
		}
	case VerticalSplitter:
		if d == Down || d == Up {
			if x, y, err := d.Move(x, y); err != nil {
				return err
			} else if !c.validCoordinates(x, y) {
				return nil
			} else {
				return c.ShineBeam(x, y, d)
			}
		} else {
			dirs := []Direction{Down, Up}
			for _, dir := range dirs {
				if newX, newY, err := dir.Move(x, y); err != nil {
					return err
				} else if !c.validCoordinates(newX, newY) {
					continue
				} else if err := c.ShineBeam(newX, newY, dir); err != nil {
					return err
				}
			}
			return nil
		}
	case HorizontalSplitter:
		if d == Left || d == Right {
			if x, y, err := d.Move(x, y); err != nil {
				return err
			} else if !c.validCoordinates(x, y) {
				return nil
			} else {
				return c.ShineBeam(x, y, d)
			}
		} else {
			dirs := []Direction{Left, Right}
			for _, dir := range dirs {
				if newX, newY, err := dir.Move(x, y); err != nil {
					return err
				} else if !c.validCoordinates(newX, newY) {
					return nil
				} else if err := c.ShineBeam(newX, newY, dir); err != nil {
					return err
				}
			}
			return nil
		}
	default:
		return fmt.Errorf("invalid tile type %v", (*tile).Type)
	}
}

func (c *Contraption) NumberOfEnergisedTiles() int {
	if c == nil {
		return 0
	}
	sum := 0
	for _, row := range c.Tiles {
		for _, t := range row {
			if len(t.Visits) > 0 {
				sum++
			}
		}
	}
	return sum
}
