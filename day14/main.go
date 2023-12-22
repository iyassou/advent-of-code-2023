package main

import (
	_ "embed"
	"log"
)

//go:embed input.txt
var input []byte

func partone(input []byte) {
	p, err := NewPlatform(input)
	if err != nil {
		log.Fatal(err)
	}
	if err := p.TiltNorth(); err != nil {
		log.Fatal(err)
	}
	log.Println("Part One:", p.TotalNorthLoad())
}

func parttwo(input []byte) {
	p, err := NewPlatform(input)
	if err != nil {
		log.Fatal(err)
	}
	if err := p.SpinCycle(1_000_000_000); err != nil {
		log.Fatal(err)
	}
	log.Println("Part Two:", p.TotalNorthLoad())
}

func main() {
	partone(input)
	parttwo(input)
}
