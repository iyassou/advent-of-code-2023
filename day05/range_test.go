package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func sampleRangeInputs() []string {
	return []string{
		"50 98 2",
		"52 50 48",
	}
}

func sampleRanges() ([]*Range, error) {
	ranges := []*Range{}
	for _, input := range sampleRangeInputs() {
		if r, err := NewRange(input); err != nil {
			return nil, err
		} else {
			ranges = append(ranges, r)
		}
	}
	return ranges, nil
}

func TestRangeParsing(t *testing.T) {
	ranges, err := sampleRanges()
	if err != nil {
		t.Fatal(err)
	}
	expected := []*Range{
		{DestinationStart: 50, SourceStart: 98, Length: 2},
		{DestinationStart: 52, SourceStart: 50, Length: 48},
	}
	if len(ranges) != len(expected) {
		t.Fatal("bruh")
	}
	for i, r := range ranges {
		exp := expected[i]
		if !cmp.Equal(r, exp) {
			t.Fatalf("actual %v differs from expected %v", r, exp)
		}
	}
}

func TestRangeContains(t *testing.T) {
	inputs := [][]int{
		{98, 99, 97, 100, 1, 12},
		{50, 97, 67, 54, 88, 79, 72, 49, 98, -1},
	}
	expected := [][]bool{
		{true, true, false, false, false, false},
		{true, true, true, true, true, true, true, false, false, false},
	}
	if len(inputs) != len(expected) {
		t.Fatal("bruh")
	} else {
		for j, inp := range inputs {
			if len(inp) != len(expected[j]) {
				t.Fatal("brother")
			}
		}
	}
	if ranges, err := sampleRanges(); err != nil {
		t.Fatal(err)
	} else {
		for i, r := range ranges {
			for j, inp := range inputs[i] {
				exp := expected[i][j]
				actual := r.Contains(inp)
				if actual != exp {
					t.Fatalf("expected %t for input %v and range %v", exp, inp, r)
				}
			}
		}
	}
}
