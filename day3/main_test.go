package main

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func sampleEngineInput() []byte {
	return []byte("467..114..\r\n...*......\r\n..35..633.\r\n......#...\r\n617*......\r\n.....+.58.\r\n..592.....\r\n......755.\r\n...$.*....\r\n.664.598..")
}

func TestNewEngineSchematicsParsing(t *testing.T) {
	in := sampleEngineInput()
	expected := &Engine{
		Schematic: [][]byte{
			[]byte(`467..114..`),
			[]byte(`...*......`),
			[]byte(`..35..633.`),
			[]byte(`......#...`),
			[]byte(`617*......`),
			[]byte(`.....+.58.`),
			[]byte(`..592.....`),
			[]byte(`......755.`),
			[]byte(`...$.*....`),
			[]byte(`.664.598..`),
		},
		Numbers: map[int][]Number{
			0: {
				{Value: 467, Start: 0, End: 3},
				{Value: 114, Start: 5, End: 8},
			},
			2: {
				{Value: 35, Start: 2, End: 4},
				{Value: 633, Start: 6, End: 9},
			},
			4: {{Value: 617, Start: 0, End: 3}},
			5: {{Value: 58, Start: 7, End: 9}},
			6: {{Value: 592, Start: 2, End: 5}},
			7: {{Value: 755, Start: 6, End: 9}},
			9: {
				{Value: 664, Start: 1, End: 4},
				{Value: 598, Start: 5, End: 8},
			},
		},
		CandidateGears: [][]int{{1, 3}, {4, 3}, {8, 5}},
	}
	if e, err := NewEngine(in); err != nil {
		t.Fatal(err)
	} else if !cmp.Equal(e, expected) {
		t.Fatalf("expected %v\ngot %v", expected, e)
	}
}

func TestIsSymbol(t *testing.T) {
	symbols := []byte{'*', '/', '$', '-'}
	notSymbols := []byte{'.', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, s := range symbols {
		if !IsSymbol(s) {
			t.Fatalf("%q not recognised as a symbol", s)
		}
	}
	for _, n := range notSymbols {
		if IsSymbol(n) {
			t.Fatalf("%q recognised as a symbol", n)
		}
	}
}

func TestEnginePartNumbers(t *testing.T) {
	expected := []int{35, 467, 592, 598, 617, 633, 664, 755}
	e, err := NewEngine(sampleEngineInput())
	if err != nil {
		t.Fatal(err)
	}
	p := e.PartNumbers()
	sort.Ints(p)
	if !cmp.Equal(p, expected) {
		t.Fatalf("expected %v, got %v", expected, p)
	}
}

func TestGearRatios(t *testing.T) {
	expected := []int{16345, 451490}
	e, err := NewEngine(sampleEngineInput())
	if err != nil {
		t.Fatal(err)
	}
	g := e.GearRatios()
	sort.Ints(g)
	if !cmp.Equal(g, expected) {
		t.Fatalf("expected %v, got %v", expected, g)
	}
}

func TestCandidateGearParsing(t *testing.T) {
	in := []byte("*...*....\r\n...***...\r\n....*....\r\n....*....\r\n....*....\r\n")
	expected := [][]int{{0, 0}, {0, 4}, {1, 3}, {1, 4}, {1, 5}, {2, 4}, {3, 4}, {4, 4}}
	e, err := NewEngine(in)
	if err != nil {
		t.Fatal(err)
	}
	if actual := e.CandidateGears; !cmp.Equal(actual, expected) {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
