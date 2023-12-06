package main

import (
	_ "embed"
	"log"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	chunks := strings.Split(input, "\r\n\r\n")
	if a, err := NewAlmanac(chunks); err != nil {
		log.Fatal(err)
	} else {
		min := math.MaxInt
		for _, seed := range a.Seeds {
			if loc := a.LocationNumber(seed); loc < min {
				min = loc
			}
		}
		log.Println("Part One:", min)
		min = math.MaxInt
		for _, seed := range a.SeedsFlattened() {
			if loc := a.LocationNumber(seed); loc < min {
				min = loc
			}
		}
		log.Println("Part Two:", min)
	}
}
