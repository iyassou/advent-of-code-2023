package main

import (
	_ "embed"
	"log"
)

//go:embed input.txt
var input []byte

func partone() {
	c, err := NewContraption(input)
	if err != nil {
		log.Fatal(err)
	}
	if err := c.ShineBeam(0, 0, Right); err != nil {
		log.Fatal(err)
	}
	log.Println("Part One:", c.NumberOfEnergisedTiles())
}

func parttwo() {
	c, err := NewContraption(input)
	if err != nil {
		log.Fatal(err)
	}
	max := -1
	for _, dir := range []Direction{Left, Right, Up, Down} {
		var coords [][2]int
		switch dir {
		case Left, Right:
			coords = make([][2]int, c.Height())
			var x int
			if dir == Left {
				x = c.Width() - 1
			}
			for y := 0; y < c.Height(); y++ {
				coords[y] = [2]int{x, y}
			}
		case Up, Down:
			coords = make([][2]int, c.Width())
			var y int
			if dir == Up {
				y = c.Height() - 1
			}
			for x := 0; x < c.Width(); x++ {
				coords[x] = [2]int{x, y}
			}
		}
		for _, coord := range coords {
			x, y := coord[0], coord[1]
			if err := c.ShineBeam(x, y, dir); err != nil {
				log.Fatal(err)
			} else if energy := c.NumberOfEnergisedTiles(); energy > max {
				max = energy
			}
			c.resetTileVisits()
		}
	}
	log.Println("Part Two:", max)
}

func main() {
	partone()
	parttwo()
}
