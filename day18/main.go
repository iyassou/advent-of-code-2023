package main

import (
	_ "embed"
	"log"

	"github.com/iyassou/advent-of-code-2023/internal"
)

//go:embed input.txt
var input string

func routine(lines []string, part string) {
	part1 := false
	if part == "One" {
		part1 = true
	} else if part != "Two" {
		log.Fatal("bruh")
	}
	x, y := 0, 0
	b := 0
	A := 0
	for _, line := range lines {
		ins := NewInstruction(line)
		if !part1 {
			ins.DecodeColour()
		}
		newX, newY := ins.Apply(x, y)
		A += newX*y - newY*x
		b += ins.length
		x, y = newX, newY
	}
	A /= 2
	answer := A + b/2 + 1
	log.Printf("Part %s: %d", part, answer)
}

func main() {
	lines := internal.Lines(input)
	routine(lines, "One")
	routine(lines, "Two")
}
