package main

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type Number struct {
	Start int
	End   int
	Value int
}

type Engine struct {
	Schematic      [][]byte
	Numbers        map[int][]Number
	CandidateGears [][]int
}

func NewEngine(schematic []byte) (*Engine, error) {
	e := &Engine{
		Schematic: bytes.Split(schematic, []byte{'\r', '\n'}),
		Numbers:   map[int][]Number{},
	}
	numRegex, err := regexp.Compile(`(\d+)`)
	if err != nil {
		return nil, err
	}
	candidateGearRegex, err := regexp.Compile(`\*`)
	if err != nil {
		return nil, err
	}
	for i, line := range e.Schematic {
		nums := numRegex.FindAllIndex(line, -1)
		for _, num := range nums {
			n := Number{Start: num[0], End: num[1]}
			if val, err := strconv.Atoi(string(line[n.Start:n.End])); err != nil {
				return nil, err
			} else {
				n.Value = val
				e.Numbers[i] = append(e.Numbers[i], n)
			}
		}
		gears := candidateGearRegex.FindAllIndex(line, -1)
		for _, gear := range gears {
			e.CandidateGears = append(e.CandidateGears, []int{i, gear[0]})
		}
	}
	return e, nil
}

func (e *Engine) MaxHeight() int {
	return len(e.Schematic) - 1
}

func (e *Engine) MaxWidth(line int) int {
	return len(e.Schematic[line]) - 1
}

func IsSymbol(b byte) bool {
	return !(unicode.IsDigit(rune(b)) || b == '.')
}

func (e *Engine) PartNumbers() []int {
	p := []int{}
	maxHeight := e.MaxHeight()
	for line, nums := range e.Numbers {
		maxWidth := e.MaxWidth(line)
		for _, num := range nums {
			// determine adjacent cells' bbox
			minI, minJ := line, num.Start
			if line > 0 {
				minI--
			}
			if num.Start > 0 {
				minJ--
			}
			maxI, maxJ := line, num.End-1
			if line < maxHeight {
				maxI++
			}
			if num.End < maxWidth {
				maxJ++
			}
			// add adjacent cells' coordinates
			adj := [][]int{}
			for j := minJ; j <= maxJ; j++ {
				adj = append(adj, []int{minI, j})
				adj = append(adj, []int{maxI, j})
			}
			for i := minI + 1; i < maxI; i++ {
				adj = append(adj, []int{i, minJ})
				adj = append(adj, []int{i, maxJ})
			}
			// check for symbols
			for _, loc := range adj {
				if IsSymbol(e.Schematic[loc[0]][loc[1]]) {
					p = append(p, num.Value)
					break
				}
			}
		}
	}
	return p
}

func (e *Engine) GearRatios() []int {
	gearRatios := []int{}
	maxHeight := e.MaxHeight()
	for _, loc := range e.CandidateGears {
		gx, gy := loc[0], loc[1]
		// determine the bbox
		maxWidth := e.MaxWidth(loc[0])
		minI, minJ := gx, gy
		if minI > 0 {
			minI--
		}
		if minJ > 0 {
			minJ--
		}
		maxI, maxJ := gx, gy
		if maxI < maxHeight {
			maxI++
		}
		if maxJ < maxWidth {
			maxJ++
		}
		// create adjacent cells' coordinates
		adj := [][]int{}
		for i := minI; i <= maxI; i++ {
			for j := minJ; j <= maxJ; j++ {
				if i == gx && j == gy {
					continue
				}
				adj = append(adj, []int{i, j})
			}
		}
		// check for adjacent numbers
		nums := map[Number]bool{}
		for _, aloc := range adj {
			nx, ny := aloc[0], aloc[1]
			for _, n := range e.Numbers[nx] {
				if n.Start <= ny && ny < n.End {
					nums[n] = true
					continue
				}
			}
		}
		// calculate gear ratio
		if len(nums) == 2 {
			ratio := 1
			for n := range nums {
				ratio *= n.Value
			}
			gearRatios = append(gearRatios, ratio)
		}
	}
	return gearRatios
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: <program> <input.txt>")
	}
	if b, err := os.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else if e, err := NewEngine(b); err != nil {
		log.Fatal("bruh")
	} else {
		parts := e.PartNumbers()
		ratios := e.GearRatios()
		one, two := 0, 0
		for _, p := range parts {
			one += p
		}
		for _, r := range ratios {
			two += r
		}
		log.Println("Part One:", one)
		log.Println("Part Two:", two)
	}
}
