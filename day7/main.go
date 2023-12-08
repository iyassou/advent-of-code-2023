package main

import (
	_ "embed"
	"errors"
	"log"
	"sort"

	"github.com/iyassou/advent-of-code-2023/internal"
)

//go:embed input.txt
var input string

func solution(input string, part int) (int, error) {
	if !(part == 1 || part == 2) {
		return 0, errors.New("bruh")
	}
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
		return hands[j].BetterThan(hands[i], part)
	})
	winnings := 0
	for i, h := range hands {
		winnings += (i + 1) * h.Bid
	}
	return winnings, nil
}

func main() {
	if winnings, err := solution(input, 1); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Part One:", winnings)
	}
	if winnings, err := solution(input, 2); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Part Two:", winnings)
	}
}
