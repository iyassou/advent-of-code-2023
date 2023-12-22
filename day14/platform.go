package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/iyassou/advent-of-code-2023/internal"
)

type Platform struct {
	Height, Width int
	Data          []byte
	Facing        Orientation
}

type Orientation int

const (
	North Orientation = 0
	East  Orientation = 1
	South Orientation = 2
	West  Orientation = 3
)

func (o Orientation) String() string {
	switch o {
	case 0:
		return "North"
	case 1:
		return "East"
	case 2:
		return "South"
	case 3:
		return "West"
	default:
		return "?"
	}
}

var errNilPlatform = errors.New("nil Platform")

func NewPlatform(input []byte) (*Platform, error) {
	lines := internal.Lines(input)
	if len(lines) == 0 {
		return nil, errors.New("could not read platform rows")
	}
	p := &Platform{
		Height: len(lines),
		Width:  len(lines[0]),
		Data:   bytes.Join(lines, nil),
		Facing: North,
	}
	return p, nil
}

func (p *Platform) String() string {
	if p == nil {
		return ""
	}
	var sb strings.Builder
	for y := 0; y < p.Height; y++ {
		for x := 0; x < p.Width; x++ {
			if c, err := p.linearCoordinate(x, y); err != nil {
				return "X"
			} else {
				sb.WriteByte(p.Data[c])
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (p *Platform) linearCoordinate(x, y int) (int, error) {
	if p == nil {
		return 0, errNilPlatform
	}
	if !(0 <= x && x < p.Width) {
		return 0, fmt.Errorf("invalid x coordinate %d", x)
	}
	if !(0 <= y && y < p.Height) {
		return 0, fmt.Errorf("invalid y coordinate %d", y)
	}
	switch p.Facing {
	case North:
		// x, y
		return x + y*p.Width, nil
	case East:
		// y, p.Width-x
		return y + (p.Width-x-1)*p.Height, nil
	case South:
		// p.Height-x, p.Width-y
		return p.Height - x - 1 + (p.Width-y-1)*p.Width, nil
	case West:
		// y, x
		return (p.Width - 1 - y) + x*p.Height, nil
	default:
		return 0, fmt.Errorf("unrecognised orientation %v", p.Facing)
	}
}

func (p *Platform) Get(x, y int) (byte, error) {
	if p == nil {
		return 0, errNilPlatform
	}
	coord, err := p.linearCoordinate(x, y)
	if err != nil {
		return 0, err
	}
	return p.Data[coord], nil
}

func (p *Platform) Set(x, y int, t byte) error {
	if p == nil {
		return errNilPlatform
	}
	coord, err := p.linearCoordinate(x, y)
	if err != nil {
		return err
	}
	if !(t == '.' || t == 'O' || t == '#') {
		return fmt.Errorf("invalid value %v", t)
	}
	p.Data[coord] = t
	return nil
}

func (p *Platform) TiltNorth() error {
	for x := 0; x < p.Width; x++ {
		emptySlots, roundRocks := []int{}, []int{}
		slideNorth := func() error {
			if len(roundRocks) == 0 || len(emptySlots) == 0 {
				return nil
			}
			for i := len(roundRocks) - 1; i > -1; i, roundRocks = i-1, roundRocks[:i] {
				j := len(emptySlots) - 1
				if roundRocks[i] < emptySlots[j] {
					continue
				}
				if err := p.Set(x, emptySlots[j], 'O'); err != nil {
					return err
				}
				if err := p.Set(x, roundRocks[i], '.'); err != nil {
					return err
				}
				emptySlots[j] = roundRocks[i]
				for k := j - 1; k > -1; k-- {
					if emptySlots[k] < emptySlots[j] {
						emptySlots[k], emptySlots[j] = emptySlots[j], emptySlots[k]
						j = k
					} else {
						break
					}
				}
			}
			return nil
		}
		for y := p.Height - 1; y > -1; y-- {
			if b, err := p.Get(x, y); err != nil {
				return err
			} else if b == '.' {
				emptySlots = append(emptySlots, y)
			} else if b == 'O' {
				roundRocks = append(roundRocks, y)
			} else if b == '#' {
				if err := slideNorth(); err != nil {
					return err
				}
				emptySlots = []int{}
				roundRocks = []int{}
			}
		}
		if err := slideNorth(); err != nil {
			return err
		}
	}
	return nil
}

func (p *Platform) rotateRight() {
	if p == nil {
		return
	}
	p.Height, p.Width = p.Width, p.Height
	p.Facing = (p.Facing + 1) % 4
}

func (p *Platform) SpinCycle(n int) error {
	if p == nil {
		return errNilPlatform
	}
	lastCycleNorthLoads := make([]int, 4)
	currentCycleNorthLoads := make([]int, 4)
	var err error
	for i := 0; i < n; i++ {
		for j := 0; j < 4; j++ {
			if err = p.TiltNorth(); err != nil {
				return err
			}
			if load := p.TotalNorthLoad(); load == -1 {
				return errors.New("could not calculate north load")
			} else {
				currentCycleNorthLoads[j] = load
			}
			p.rotateRight()
		}
		repeatingCycle := true
		for i, lc := range lastCycleNorthLoads {
			if lc != currentCycleNorthLoads[i] {
				repeatingCycle = false
				break
			}
		}
		if repeatingCycle {
			return nil
		}
		copy(lastCycleNorthLoads, currentCycleNorthLoads)
	}
	return nil
}

func (p *Platform) TotalNorthLoad() int {
	if p == nil {
		return -1
	}
	totalLoad := 0
	for y := 0; y < p.Height; y++ {
		load := p.Height - y
		for x := 0; x < p.Width; x++ {
			if b, err := p.Get(x, y); err != nil {
				return -1
			} else if b == 'O' {
				totalLoad += load
			}
		}
	}
	return totalLoad
}
