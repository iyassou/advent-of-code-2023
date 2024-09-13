package main

type FlipFlop struct {
	name    string
	state   bool
	outputs []string
}

func NewFlipFlop(name string, outputs []string) FlipFlop {
	return FlipFlop{name: name, outputs: outputs}
}

func (f FlipFlop) GetName() string {
	return f.name
}

func (f FlipFlop) GetOutputs() []string {
	return f.outputs
}

func (f *FlipFlop) ProcessPulse(p Pulse) []Pulse {
	if p.high {
		return []Pulse{}
	}
	f.state = !f.state
	pulses := make([]Pulse, len(f.outputs))
	for i, output := range f.outputs {
		pulses[i].source = f.name
		pulses[i].destination = output
		pulses[i].high = f.state
	}
	return pulses
}
