package main

type Broadcaster struct {
	outputs []string
}

func NewBroadcaster(outputs []string) Broadcaster {
	return Broadcaster{outputs: outputs}
}

func (b Broadcaster) GetName() string {
	return "broadcaster"
}

func (b Broadcaster) GetOutputs() []string {
	return b.outputs
}

func (b Broadcaster) ProcessPulse(p Pulse) []Pulse {
	pulses := make([]Pulse, len(b.outputs))
	for i, output := range b.outputs {
		pulses[i].source = b.GetName()
		pulses[i].destination = output
		pulses[i].high = p.high
	}
	return pulses
}
