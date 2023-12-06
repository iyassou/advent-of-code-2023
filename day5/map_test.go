package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func sampleMapInputs() []string {
	return []string{
		"seed-to-soil map:\r\n50 98 2\r\n52 50 48",
		"soil-to-fertilizer map:\r\n0 15 37\r\n37 52 2\r\n39 0 15",
		"fertilizer-to-water map:\r\n49 53 8\r\n0 11 42\r\n42 0 7\r\n57 7 4",
	}
}

func sampleMaps() ([]*Map, error) {
	maps := []*Map{}
	for _, input := range sampleMapInputs() {
		if m, err := NewMap(input); err != nil {
			return nil, err
		} else {
			maps = append(maps, m)
		}
	}
	return maps, nil
}

func TestMapParsing(t *testing.T) {
	maps, err := sampleMaps()
	if err != nil {
		t.Fatal(err)
	}
	expected := []*Map{
		{From: Entry("seed"), To: Entry("soil"), Ranges: []*Range{
			{DestinationStart: 50, SourceStart: 98, Length: 2},
			{DestinationStart: 52, SourceStart: 50, Length: 48},
		}},
		{From: Entry("soil"), To: Entry("fertilizer"), Ranges: []*Range{
			{DestinationStart: 0, SourceStart: 15, Length: 37},
			{DestinationStart: 37, SourceStart: 52, Length: 2},
			{DestinationStart: 39, SourceStart: 0, Length: 15},
		}},
		{From: Entry("fertilizer"), To: Entry("water"), Ranges: []*Range{
			{DestinationStart: 49, SourceStart: 53, Length: 8},
			{DestinationStart: 0, SourceStart: 11, Length: 42},
			{DestinationStart: 42, SourceStart: 0, Length: 7},
			{DestinationStart: 57, SourceStart: 7, Length: 4},
		}},
	}
	if len(maps) != len(expected) {
		t.Fatal("bruh")
	}
	for i, r := range maps {
		exp := expected[i]
		if !cmp.Equal(r, exp) {
			t.Fatalf("actual %v differs from expected %v", r, exp)
		}
	}
}

func TestMapConvert(t *testing.T) {
	inputs := [][]int{
		{0, 1, 48, 49, 50, 51, 96, 97, 98, 99},
		{81, 14, 57, 13},
		{81, 53, 57, 52},
	}
	expected := [][]int{
		{0, 1, 48, 49, 52, 53, 98, 99, 50, 51},
		{81, 53, 57, 52},
		{81, 49, 53, 41},
	}
	if len(inputs) != len(expected) {
		t.Fatal("bruh")
	} else {
		for j, inp := range inputs {
			if len(inp) != len(expected[j]) {
				t.Fatal("reuf")
			}
		}
	}
	if maps, err := sampleMaps(); err != nil {
		t.Fatal(err)
	} else {
		for i, m := range maps {
			for j, inp := range inputs[i] {
				exp := expected[i][j]
				actual := m.Convert(inp)
				if actual != exp {
					t.Fatalf("expected %v for input %v, got %v", exp, inp, actual)
				}
			}
		}
	}
}
