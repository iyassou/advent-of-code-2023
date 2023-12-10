package main

import (
	_ "embed"
	"log"

	"github.com/iyassou/advent-of-code-2023/internal"
)

//go:embed input.txt
var input string

func read_input() ([]History, error) {
	lines := internal.Lines(input)
	histories := make([]History, len(lines))
	for i, line := range internal.Lines(input) {
		if h, err := NewHistory(line); err != nil {
			return nil, err
		} else {
			histories[i] = h
		}
	}
	return histories, nil
}

func main() {
	hs, err := read_input()
	if err != nil {
		log.Fatal(err)
	}
	p1, p2 := 0, 0
	for _, h := range hs {
		if v, err := h.Extrapolate(Forward); err != nil {
			log.Fatal(err)
		} else {
			p1 += v
		}
		if v, err := h.Extrapolate(Backward); err != nil {
			log.Fatal(err)
		} else {
			p2 += v
		}
	}
	log.Println("Part One:", p1)
	log.Println("Part Two:", p2)
}
