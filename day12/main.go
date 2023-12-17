package main

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/iyassou/advent-of-code-2023/internal"
)

//go:embed input.txt
var input []byte

func partone() {
	sum := 0
	for _, line := range internal.Lines(input) {
		if a, err := NewArrangement(line); err != nil {
			log.Fatal(err)
		} else {
			sum += a.BruteForce()
		}
	}
	log.Println("Part One:", sum)
}

func parttwo() {
	sum := 0
	for _, line := range internal.Lines(input) {
		f := bytes.Fields(line)
		condition, groups := f[0], f[1]
		unfurledCondition := make([]byte, len(f[0])*5+4)
		unfurledGroups := make([]byte, len(f[1])*5+4)
		for i := 0; i < 5; i++ {
			copy(unfurledCondition[i+i*len(f[0]):i+(i+1)*len(f[0])], condition)
			copy(unfurledGroups[i+i*len(f[1]):i+(i+1)*len(f[1])], groups)
			if i != 4 {
				unfurledCondition[i+(i+1)*len(f[0])] = '?'
				unfurledGroups[i+(i+1)*len(f[1])] = ','
			}
		}
		unfurled := bytes.Join([][]byte{unfurledCondition, unfurledGroups}, []byte{' '})
		if a, err := NewArrangement(unfurled); err != nil {
			log.Fatal(err)
		} else {
			sum += a.BruteForce()
		}
	}
	log.Println("Part Two:", sum)
}

func main() {
	// partone()
	parttwo()
}
