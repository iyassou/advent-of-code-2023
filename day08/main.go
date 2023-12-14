package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/iyassou/advent-of-code-2023/internal"
)

//go:embed input.txt
var input string

func read_input() (string, string, error) {
	lines := strings.Split(input, "\r\n\r\n")
	if l := len(lines); l != 2 {
		return "", "", fmt.Errorf("expected 2 part, got %d", l)
	}
	return internal.ShortestRepeatingSubstring(lines[0]), lines[1], nil
}

func main() {
	path, graph, err := read_input()
	if err != nil {
		log.Fatal("bruh")
	}
	if g, err := NewGraph(graph); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Part One:", g.TravelFromTo(path, "AAA", "ZZZ"))
		log.Println("Part Two:", g.GhostTravel(path, "A", "Z"))
	}
}
