package main

import (
	"bytes"
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Subset struct {
	Red   int
	Blue  int
	Green int
}

func (s *Subset) Power() int {
	return s.Red * s.Blue * s.Green
}

type Game struct {
	ID      int
	Reveals []Subset
}

func NewGame(input string) (*Game, error) {
	idRegex, err := regexp.Compile(`Game (\d*):`)
	if err != nil {
		return nil, err
	}
	idString := idRegex.FindStringSubmatch(input)
	if idString == nil {
		return nil, errors.New("couldn't identify game id")
	}
	id, err := strconv.Atoi(idString[1])
	if err != nil {
		return nil, err
	}
	g := &Game{ID: id}
	colourRegex, err := regexp.Compile(`(\d*) (red|green|blue)`)
	if err != nil {
		return nil, err
	}
	for _, r := range strings.Split(input[len(idString[0]):], ";") {
		reveal := Subset{}
		m := colourRegex.FindAllStringSubmatch(r, -1)
		for _, cube := range m {
			if i, err := strconv.Atoi(cube[1]); err != nil {
				return nil, err
			} else if cube[2] == "red" {
				reveal.Red = i
			} else if cube[2] == "blue" {
				reveal.Blue = i
			} else {
				reveal.Green = i
			}
		}
		g.Reveals = append(g.Reveals, reveal)
	}
	return g, nil
}

func (g *Game) Possible(bag *Subset) bool {
	if g.Reveals == nil {
		return false
	}
	for _, r := range g.Reveals {
		if r.Red > bag.Red || r.Green > bag.Green || r.Blue > bag.Blue {
			return false
		}
	}
	return true
}

func (g *Game) SmallestSubset() *Subset {
	s := &Subset{}
	if g.Reveals == nil {
		return s
	}
	for _, r := range g.Reveals {
		if r.Red > s.Red {
			s.Red = r.Red
		}
		if r.Green > s.Green {
			s.Green = r.Green
		}
		if r.Blue > s.Blue {
			s.Blue = r.Blue
		}
	}
	return s
}

func main() {
	if b, err := os.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		idSum := 0
		powerSum := 0
		part1 := &Subset{Red: 12, Green: 13, Blue: 14}
		for _, lb := range bytes.Split(b, []byte{'\r', '\n'}) {
			g, err := NewGame(string(lb))
			if err != nil {
				log.Fatal(err)
			}
			if g.Possible(part1) {
				idSum += g.ID
			}
			powerSum += g.SmallestSubset().Power()
		}
		log.Println("Part One:", idSum)
		log.Println("Part Two:", powerSum)
	}
}
