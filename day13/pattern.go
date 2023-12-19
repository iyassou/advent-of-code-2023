package main

import (
	"log"

	"github.com/iyassou/advent-of-code-2023/internal"
)

type Pattern struct {
	Rows    []int
	Columns []int
}

func NewPattern(input []byte) *Pattern {
	// ASSUMPTION: input is rectangular
	rows := internal.Lines(input)
	numRows, numColumns := len(rows), len(rows[0])
	p := &Pattern{Rows: make([]int, numRows), Columns: make([]int, numColumns)}
	for i, row := range rows {
		for j, b := range row {
			if b == '#' {
				p.Rows[i] += 1 << j
				p.Columns[j] += 1 << i
			}
		}
	}
	return p
}

func (p *Pattern) FindReflection() int {
	// Negative return value means the line of reflection is vertical.
	// Positive return value means the line of reflection is horizontal.
	// Zero return value means no reflection was found.
	for i, r := range p.Rows[1:] {
		if r == p.Rows[i] {
			match := true
			for j, k := i-1, i+2; j > -1 && k < len(p.Rows); j, k = j-1, k+1 {
				if p.Rows[j] != p.Rows[k] {
					match = false
					break
				}
			}
			if match {
				log.Println("i:", i)
				return i + 1
			}
		}
	}
	for i, c := range p.Columns[1:] {
		if c == p.Columns[i] {
			match := true
			for j, k := i-1, i+2; j > -1 && k < len(p.Columns); j, k = j-1, k+1 {
				if p.Columns[j] != p.Columns[k] {
					match = false
					break
				}
			}
			if match {
				log.Println("i:", -i)
				return -(i + 1)
			}
		}
	}
	return 0
}

func (p *Pattern) FindPossibleSmudges() [][2]int {
	smudges := [][2]int{}
	for i := 0; i < len(p.Rows); i++ {
		for j := i + 1; j < len(p.Rows); j++ {
			a, b := p.Rows[i], p.Rows[j]
			if a == b {
				continue
			}
			if xor := a ^ b; xor&(xor-1) == 0 {
				y := 0
				for ; xor != 1; xor, y = xor>>1, y+1 {
				}
				smudges = append(smudges, [2]int{i, y}, [2]int{j, y})
			}
		}
	}
	for i := 0; i < len(p.Columns); i++ {
		for j := i + 1; j < len(p.Columns); j++ {
			a, b := p.Columns[i], p.Columns[j]
			if a == b {
				continue
			}
			if xor := a ^ b; xor&(xor-1) == 0 {
				x := 0
				for ; xor != 1; xor, x = xor>>1, x+1 {
				}
				smudges = append(smudges, [2]int{x, i}, [2]int{x, j})
			}
		}
	}
	return smudges
}

func (p *Pattern) FindSmudgedReflection(originalReflection int) int {
	// Negative return value means the line of reflection is vertical.
	// Positive return value means the line of reflection is horizontal.
	// Zero return value means no reflection was found.
	if originalReflection == 0 {
		originalReflection = p.FindReflection()
	}
	reflections := []int{}
	smudges := p.FindPossibleSmudges()
	log.Println("smudges:", smudges)
	for _, smudge := range smudges {
		x, y := smudge[0], smudge[1]
		p.WipeSmudge(x, y)
		if r := p.FindReflection(); r != originalReflection {
			reflections = append(reflections, r)
		} else {
			log.Printf("(original, new) = (%d, %d)", originalReflection, r)
		}
		p.WipeSmudge(x, y)
	}
	if len(reflections) != 1 {
		log.Fatalf("reflections: %v", reflections)
	}
	return reflections[0]
}

func (p *Pattern) WipeSmudge(x, y int) {
	if !(0 <= x && x < len(p.Rows)) {
		return
	}
	if !(0 <= y && y < len(p.Columns)) {
		return
	}
	p.Rows[x] ^= 1 << y
	p.Columns[y] ^= 1 << x
}
