package main

import "fmt"

type Pulse struct {
	source      string
	destination string
	high        bool
}

func (p Pulse) String() string {
	return fmt.Sprintf(
		"%s -%s-> %s",
		p.source,
		map[bool]string{true: "high", false: "low"}[p.high],
		p.destination,
	)
}

type Module interface {
	GetName() string
	GetOutputs() []string
	ProcessPulse(Pulse) []Pulse
}

type DummyModule struct{ name string }

func (d DummyModule) GetName() string            { return d.name }
func (d DummyModule) GetOutputs() []string       { return []string{} }
func (d DummyModule) ProcessPulse(Pulse) []Pulse { return nil }
