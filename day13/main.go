package main

import (
	"bytes"
	_ "embed"
	"log"
)

//go:embed input.txt
var input []byte

func reflectionContribution(r int) int {
	if r < 0 {
		return -r
	}
	if r > 0 {
		return r * 100
	}
	log.Fatalf("bruh %d", r)
	return 0
}

func main() {
	p1, p2 := 0, 0
	for _, inp := range bytes.Split(input, []byte("\r\n\r\n")) {
		p := NewPattern(inp)
		originalReflection := p.FindReflection()
		p1 += reflectionContribution(originalReflection)
		smudgedReflection := p.FindSmudgedReflection(originalReflection)
		p2 += reflectionContribution(smudgedReflection)
	}
	log.Println("Part One:", p1)
	log.Println("Part Two:", p2)
}
