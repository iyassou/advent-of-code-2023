package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func sampleAlmanacInputs() [][]string {
	return [][]string{
		{
			"seeds: 79 14 55 13",
			"seed-to-soil map:\r\n50 98 2\r\n52 50 48",
			"soil-to-fertilizer map:\r\n0 15 37\r\n37 52 2\r\n39 0 15",
			"fertilizer-to-water map:\r\n49 53 8\r\n0 11 42\r\n42 0 7\r\n57 7 4",
		},
	}
}

func sampleAlmanacs() ([]*Almanac, error) {
	almanacs := []*Almanac{}
	for _, input := range sampleAlmanacInputs() {
		if a, err := NewAlmanac(input); err != nil {
			return nil, err
		} else {
			almanacs = append(almanacs, a)
		}
	}
	return almanacs, nil
}

func TestAlmanacParsing(t *testing.T) {
	as, err := sampleAlmanacs()
	if err != nil {
		t.Fatal(err)
	}
	expected := []*Almanac{
		{
			Seeds: []int{79, 14, 55, 13},
			Maps: []*Map{
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
			},
		},
	}
	if len(as) != len(expected) {
		t.Fatal("bruh")
	}
	for i, a := range as {
		exp := expected[i]
		if !cmp.Equal(a, exp) {
			t.Fatalf("actual %v differs from expected %v", a, exp)
		}
	}
}

func TestAlmanacLocationNumber(t *testing.T) {
	as, err := sampleAlmanacs()
	if err != nil {
		t.Fatal(err)
	}
	inputs := [][]int{
		{79, 14, 55, 13},
	}
	expected := [][]int{
		{81, 49, 53, 41},
	}
	if len(inputs) != len(expected) {
		t.Fatal("bruh")
	}
	if len(inputs) != len(as) {
		t.Fatal("reuf")
	}
	for i, a := range as {
		for j, inp := range inputs[i] {
			exp := expected[i][j]
			actual := a.LocationNumber(inp)
			if actual != exp {
				t.Fatalf("actual location number %d for seed %d differs from expected %d", actual, inp, exp)
			}
		}
	}
}
