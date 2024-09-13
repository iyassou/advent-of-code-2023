package main

type Conjunction struct {
	name    string
	state   map[string]bool
	outputs []string
}

func NewConjunction(name string, outputs []string) Conjunction {
	return Conjunction{name: name, state: map[string]bool{}, outputs: outputs}
}

func (c Conjunction) GetName() string {
	return c.name
}

func (c Conjunction) GetOutputs() []string {
	return c.outputs
}

func (c *Conjunction) ProcessPulse(p Pulse) []Pulse {
	c.state[p.source] = p.high
	send := true
	if p.high {
		send = false
		for _, state := range c.state {
			if !state {
				send = true
				break
			}
		}
	}
	pulses := make([]Pulse, len(c.outputs))
	for i, output := range c.outputs {
		pulses[i].source = c.name
		pulses[i].destination = output
		pulses[i].high = send
	}
	return pulses
}
