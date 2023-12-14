package main

import (
	_ "embed"
	"log"
)

//go:embed input.txt
var input []byte

func routine(c *CSCImage, part string) {
	by := 1
	if part == "Two" {
		by = 999_999
	} else if part != "One" {
		log.Fatal("bruh")
	}
	c.ExpandUniverse(by)
	loc := c.GalaxyLocations()
	manhattanDistance := func(i, j int) int {
		a, b := loc[i][0], loc[i][1]
		x, y := loc[j][0], loc[j][1]
		dx, dy := 0, 0
		if x > a {
			dx = x - a
		} else {
			dx = a - x
		}
		if y > b {
			dy = y - b
		} else {
			dy = b - y
		}
		return dx + dy
	}
	sum := 0
	for i := 0; i < len(loc); i++ {
		for j := i + 1; j < len(loc); j++ {
			sum += manhattanDistance(i, j)
		}
	}
	log.Printf("Part %s: %d\n", part, sum)
}

func main() {
	for _, part := range []string{"One", "Two"} {
		if img, err := NewCSCImage(input); err != nil {
			log.Fatal(err)
		} else {
			routine(img, part)
		}
	}
}
