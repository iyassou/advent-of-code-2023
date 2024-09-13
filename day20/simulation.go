package main

import (
	"strings"

	"github.com/iyassou/advent-of-code-2023/internal"
	"github.com/iyassou/advent-of-code-2023/internal/queue"
)

func ParseModules(input string) map[string]Module {
	// Store all modules on a first pass and make a note of the conjunction modules.
	modules := map[string]Module{}
	conjunctions := map[string]struct{}{}
	for _, line := range internal.Lines(input) {
		fields := strings.Split(line, " -> ")
		fullName := fields[0]
		outputs := strings.Split(fields[1], ", ")
		if strings.HasPrefix("broadcaster", fullName) {
			modules[fullName] = NewBroadcaster(outputs)
		} else {
			name := fullName[1:]
			if strings.HasPrefix(fullName, "%") {
				flipflop := NewFlipFlop(name, outputs)
				modules[name] = &flipflop
			} else {
				conj := NewConjunction(name, outputs)
				modules[name] = &conj
				conjunctions[name] = struct{}{}
			}
		}
	}
	// On a second pass update the inputs to conjunctions.
	for _, line := range internal.Lines(input) {
		fields := strings.Split(line, " -> ")
		name := fields[0]
		if !strings.HasPrefix(name, "broadcaster") {
			name = name[1:]
		}
		outputs := strings.Split(fields[1], ", ")
		for _, output := range outputs {
			if _, ok := conjunctions[output]; ok {
				modules[output].(*Conjunction).state[name] = false
			}
		}
	}
	return modules
}

func Simulate(modules map[string]Module, buttonPresses int) map[bool]int {
	modules["output"] = DummyModule{"output"}
	queue := queue.NewQueue[Pulse]()
	pulsesSent := map[bool]int{}
	for buttonPresses > 0 {
		buttonPresses--
		queue.Enqueue(Pulse{"button", "broadcaster", false})
		for queue.Len() > 0 {
			pulse, _ := queue.Dequeue()
			pulsesSent[pulse.high]++
			destinationModule, ok := modules[pulse.destination]
			if !ok {
				modules[pulse.destination] = &DummyModule{pulse.destination}
				destinationModule = modules[pulse.destination]
			}
			newPulses := destinationModule.ProcessPulse(pulse)
			for _, p := range newPulses {
				queue.Enqueue(p)
			}
		}
	}
	return pulsesSent
}

// FindPeriods returns the observed high or low periodicity of a list of modules
// for a given maximum number of button presses.
// moi = modules of interest
// soi = state of interest
func FindObservedPeriods(modules map[string]Module, moi []string, soi bool, buttonPresses int) map[string][]int {
	modules["output"] = DummyModule{"output"}
	modules["rx"] = DummyModule{"rx"}
	queue := queue.NewQueue[Pulse]()
	pressedAt := map[string][]int{}
	for i := 0; i < buttonPresses; i++ {
		queue.Enqueue(Pulse{"button", "broadcaster", false})
		for queue.Len() > 0 {
			pulse, _ := queue.Dequeue()
			if pulse.high == soi {
				for _, mod := range moi {
					if pulse.destination == mod {
						L := len(pressedAt[mod])
						if L == 0 || pressedAt[mod][L-1] != i {
							pressedAt[mod] = append(pressedAt[mod], i)
						}
					}
				}
			}
			newPulses := modules[pulse.destination].ProcessPulse(pulse)
			for _, p := range newPulses {
				queue.Enqueue(p)
			}
		}
	}
	periods := map[string][]int{}
	for k, presses := range pressedAt {
		periods[k] = []int{}
		for i, p := range presses {
			if i == 0 {
				periods[k] = append(periods[k], p+1)
			} else {
				periods[k] = append(periods[k], p-presses[i-1])
			}
		}
	}
	return periods
}
