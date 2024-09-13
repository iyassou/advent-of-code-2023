package main

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/iyassou/advent-of-code-2023/internal"
)

//go:embed input.txt
var input string

func partone() {
	modules := ParseModules(input)
	pulses := Simulate(modules, 1000)
	log.Println("Part One:", pulses[true]*pulses[false])
}

func parttwo() {
	// rx's only input is hj, a conjunction module
	// Conjunctions send a low pulse if all of their inputs are high.
	modules := ParseModules(input)
	layers := [][][]string{
		{{"&hj"}},
	}
	for j := 0; j < len(layers); j++ {
		layer := layers[j]
		newLayer := [][]string{}
		for _, input := range layer {
			for _, name := range input {
				if module, ok := modules[name[1:]].(*Conjunction); ok {
					inputs := make([]string, len(module.state))
					i := 0
					for key := range module.state {
						var name string
						if _, ok := modules[key].(*Conjunction); ok {
							name = fmt.Sprintf("&%s", key)
						} else if _, ok := modules[key].(*FlipFlop); ok {
							name = fmt.Sprintf("%%%s", key)
						}
						inputs[i] = name
						i++
					}
					newLayer = append(newLayer, inputs)
				}
			}
		}
		if len(newLayer) > 0 {
			layers = append(layers, newLayer)
		}
	}
	moi := []string{}
	for _, mods := range layers[1] {
		for _, mod := range mods {
			moi = append(moi, mod[1:])
		}
	}
	observedPeriods := FindObservedPeriods(modules, moi, false, 10_000) // 10K is arbitrary big number
	ops := make([]int, len(observedPeriods))
	i := 0
	for _, v := range observedPeriods {
		ops[i] = v[0]
		i++
	}
	if lcm, err := internal.LCM(ops...); err != nil {
		log.Fatal("bruh", err)
	} else {
		log.Println("Part Two:", lcm)
	}
}

func main() {
	partone()
	parttwo()
}
