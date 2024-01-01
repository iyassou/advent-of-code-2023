package main

import (
	_ "embed"
	"log"
	"strings"

	"github.com/iyassou/advent-of-code-2023/internal"
)

//go:embed input.txt
var input string

func parse(input string) (workflows, []part) {
	inputs := strings.Split(input, "\r\n\r\n")
	workflows, err := NewWorkflows(inputs[0])
	if err != nil {
		log.Fatal(err)
	}
	lines := internal.Lines(inputs[1])
	parts := make([]part, len(lines))
	for i, line := range lines {
		if p, err := NewPart(line); err != nil {
			log.Fatal(err)
		} else {
			parts[i] = p
		}
	}

	return workflows, parts
}

func partone(wfs workflows, parts []part) {
	sum := 0
	for _, p := range parts {
		if ok, err := wfs.Accepts(p, "in"); err != nil {
			log.Fatal(err)
		} else if ok {
			sum += p.RatingsSum()
		}
	}
	log.Println("Part One:", sum)
}

func parttwo(wfs workflows) {
	combos, err := wfs.DistinctCombinations("in")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Part Two:", combos)
}

func main() {
	wfs, parts := parse(input)
	partone(wfs, parts)
	parttwo(wfs)
}
