package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func simpleExample() string {
	return "broadcaster -> a, b, c\r\n%a -> b\r\n%b -> c\r\n%c -> inv\r\n&inv -> a"
}

func lessSimpleExample() string {
	return "broadcaster -> a\r\n%a -> inv, con\r\n&inv -> b\r\n%b -> con\r\n&con -> output"
}

func TestParseModulesSimpleExample(t *testing.T) {
	expected := map[string]Module{
		"broadcaster": Broadcaster{[]string{"a", "b", "c"}},
		"a":           &FlipFlop{"a", false, []string{"b"}},
		"b":           &FlipFlop{"b", false, []string{"c"}},
		"c":           &FlipFlop{"c", false, []string{"inv"}},
		"inv":         &Conjunction{"inv", map[string]bool{"c": false}, []string{"a"}},
	}
	actual := ParseModules(simpleExample())
	if diff := cmp.Diff(expected, actual, cmp.AllowUnexported(FlipFlop{}, Conjunction{}, Broadcaster{})); diff != "" {
		t.Fatalf("parsing failed:\n%s", diff)
	}
}

func TestParseModulesLessSimpleExample(t *testing.T) {
	expected := map[string]Module{
		"broadcaster": Broadcaster{[]string{"a"}},
		"a":           &FlipFlop{"a", false, []string{"inv", "con"}},
		"inv":         &Conjunction{"inv", map[string]bool{"a": false}, []string{"b"}},
		"b":           &FlipFlop{"b", false, []string{"con"}},
		"con":         &Conjunction{"con", map[string]bool{"a": false, "b": false}, []string{"output"}},
	}
	actual := ParseModules(lessSimpleExample())
	if diff := cmp.Diff(expected, actual, cmp.AllowUnexported(FlipFlop{}, Conjunction{}, Broadcaster{})); diff != "" {
		t.Fatalf("parsing failed:\n%s", diff)
	}
}

func TestSimulateSimpleExampleOnce(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	modules := ParseModules(simpleExample())
	pulsesSent := Simulate(modules, 1)
	if pulsesSent[true] != 4 {
		t.Fatalf("expected 4 high pulses, sent %d", pulsesSent[true])
	}
	if pulsesSent[false] != 8 {
		t.Fatalf("expected 8 low pulses, sent %d", pulsesSent[false])
	}
}

func TestSimulateSimpleExample1000(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	modules := ParseModules(simpleExample())
	pulsesSent := Simulate(modules, 1000)
	if pulsesSent[true] != 4000 {
		t.Fatalf("expected 4000 high pulses, sent %d", pulsesSent[true])
	}
	if pulsesSent[false] != 8000 {
		t.Fatalf("expected 8000 low pulses, sent %d", pulsesSent[false])
	}
}

func TestSimulateLessSimpleExampleOnce(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	modules := ParseModules(lessSimpleExample())
	pulsesSent := Simulate(modules, 1)
	if pulsesSent[true] != 4 {
		t.Fatalf("expected 4 high pulses, sent %d", pulsesSent[true])
	}
	if pulsesSent[false] != 4 {
		t.Fatalf("expected 4 low pulses, sent %d", pulsesSent[false])
	}
}

func TestSimulateLessSimpleExample1000(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	modules := ParseModules(lessSimpleExample())
	pulsesSent := Simulate(modules, 1000)
	if pulsesSent[true] != 2750 {
		t.Fatalf("expected 2750 high pulses, sent %d", pulsesSent[true])
	}
	if pulsesSent[false] != 4250 {
		t.Fatalf("expected 4250 low pulses, sent %d", pulsesSent[false])
	}
}
