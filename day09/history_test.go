package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func sampleHistoryInputs() []string {
	return []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
}

func TestHistoryParsing(t *testing.T) {
	inputs := sampleHistoryInputs()
	expected := []History{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}
	if len(inputs) != len(expected) {
		t.Fatal("bruh")
	}
	for i, inp := range inputs {
		if actual, err := NewHistory(inp); err != nil {
			t.Fatal(err)
		} else if exp := expected[i]; !cmp.Equal(actual, exp) {
			t.Fatalf("input %q: expected %v, got %v", inp, exp, actual)
		}
	}
}

func TestHistoryExtrapolate(t *testing.T) {
	inputs := sampleHistoryInputs()
	expected := [][2]int{
		{18, -3}, {28, 0}, {68, 5},
	}
	if len(inputs) != len(expected) {
		t.Fatal("reuf")
	}
	for i, inp := range inputs {
		h, err := NewHistory(inp)
		if err != nil {
			t.Fatal("wesh")
		}
		exp := expected[i][0]
		if act, err := h.Extrapolate(Forward); err != nil {
			t.Fatal(err)
		} else if exp != act {
			t.Fatalf("[Forward] input %q: expected %d, got %d", inp, exp, act)
		}
		exp = expected[i][1]
		if act, err := h.Extrapolate(Backward); err != nil {
			t.Fatal(err)
		} else if exp != act {
			t.Fatalf("[Backward] input %q: expected %d, got %d", inp, exp, act)
		}
	}
}
