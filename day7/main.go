package main

import (
	_ "embed"
	"log"
	"sort"

	"github.com/iyassou/advent-of-code-2023/internal"
)

//go:embed input.txt
var input string

func part1(input string) (int, error) {
	lines := internal.Lines(input)
	hands := make([]*Hand, len(lines))
	for i, line := range internal.Lines(input) {
		if h, err := NewHand(line); err != nil {
			return 0, err
		} else {
			hands[i] = h
		}
	}
	sort.Slice(hands, func(i, j int) bool {
		return hands[j].BetterThan(hands[i])
	})
	winnings := 0
	for i, h := range hands {
		winnings += (i + 1) * h.Bid
	}
	return winnings, nil
}

func main() {
	if winnings, err := part1(input); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Part One:", winnings)
	}
}
