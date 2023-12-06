package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	chunks := strings.Split(input, "\r\n\r\n")
	if a, err := NewAlmanac(chunks); err != nil {
		log.Fatal(err)
	} else {
		if m, err := a.MinLocationNumber(1); err != nil {
			log.Fatal(err)
		} else {
			log.Println("Part One:", m)
		}
		if m, err := a.MinLocationNumber(2); err != nil {
			log.Fatal(err)
		} else {
			log.Println("Part Two:", m)
		}
	}
}
