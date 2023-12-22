package main

import (
	"bytes"
	_ "embed"
	"log"
)

//go:embed input.txt
var input []byte

func partone() {
	var h int
	h, err := hashInitSequence(bytes.TrimSpace(input))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Part One:", h)
}

func parttwo() {
	var fp int
	if h, err := NewHashmap(input); err != nil {
		log.Fatal(err)
	} else if fp, err = h.FocusingPower(); err != nil {
		log.Fatal(err)
	}
	log.Println("Part Two:", fp)
}

func main() {
	partone()
	parttwo()
}
