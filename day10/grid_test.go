package main

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func sampleGridInput() []byte {
	return []byte(".....\r\n.F-7.\r\n.|.|.\r\n.L-J.\r\n.....")
}

func sampleGridInputWithStart() []byte {
	return []byte(".....\r\n.S-7.\r\n.|.|.\r\n.L-J.\r\n.....")
}

func getGrid(input []byte) (*Grid, error) {
	return NewGrid(input)
}

func TestNewGrid(t *testing.T) {
	inputs := [][]byte{sampleGridInput(), sampleGridInputWithStart()}
	for _, input := range inputs {
		if _, err := getGrid(input); err != nil {
			t.Fatal(err)
		}
	}
}

func TestGridGetTile(t *testing.T) {
	expected := []*Tile{
		{Char: '.', Type: Ground}, {Char: '.', Type: Ground}, {Char: '.', Type: Ground}, {Char: '.', Type: Ground}, {Char: '.', Type: Ground},
		{Char: '.', Type: Ground}, {Char: 'F', Type: SouthEast}, {Char: '-', Type: WestEast}, {Char: '7', Type: SouthWest}, {Char: '.', Type: Ground},
		{Char: '.', Type: Ground}, {Char: '|', Type: NorthSouth}, {Char: '.', Type: Ground}, {Char: '|', Type: NorthSouth}, {Char: '.', Type: Ground},
		{Char: '.', Type: Ground}, {Char: 'L', Type: NorthEast}, {Char: '-', Type: WestEast}, {Char: 'J', Type: NorthWest}, {Char: '.', Type: Ground},
		{Char: '.', Type: Ground}, {Char: '.', Type: Ground}, {Char: '.', Type: Ground}, {Char: '.', Type: Ground}, {Char: '.', Type: Ground},
	}
	inputs := [][]byte{sampleGridInput(), sampleGridInputWithStart()}
	if len(expected) != len(inputs) {
		t.Fatal("bruh")
	}
	for x, input := range inputs {
		g, err := getGrid(input)
		if err != nil {
			t.Fatal(err)
		}
		if x == 1 {
			expected[6].Char = 'S'
			expected[6].Type = Start
		}
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				coord := i*5 + j
				tile, err := g.getTile(Coord{i, j})
				if err != nil {
					t.Fatal(err)
				} else if exp := expected[coord]; !cmp.Equal(tile, exp) {
					t.Errorf("expected %v for (%d,%d), got %v", exp, i, j, tile)
				}
			}
		}
	}
}

func TestGridCanTravelTo(t *testing.T) {
	expected := [][]Coord{
		{}, {}, {}, {}, {},
		//  F/S               -                  7
		{}, {{1, 2}, {2, 1}}, {{1, 1}, {1, 3}}, {{1, 2}, {2, 3}}, {},
		//  |                     |
		{}, {{1, 1}, {3, 1}}, {}, {{1, 3}, {3, 3}}, {},
		//  L                 -                 J
		{}, {{2, 1}, {3, 2}}, {{3, 1}, {3, 3}}, {{2, 3}, {3, 2}}, {},
		{}, {}, {}, {}, {},
	}
	for _, input := range [][]byte{sampleGridInput(), sampleGridInputWithStart()} {
		g, err := getGrid(input)
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				exp := expected[i*5+j]
				actual := g.canTravelTo(Coord{i, j})
				if len(actual) != len(exp) {
					t.Fatalf("[%d,%d] expected %v, got %v", i, j, exp, actual)
				}
				sort.Slice(actual, func(i, j int) bool { return actual[i].LessThan(actual[j]) })
				if !cmp.Equal(actual, exp) {
					t.Fatalf("[%d,%d] expected %v, got %v", i, j, exp, actual)
				}
			}
		}
	}
}

func TestGridGetMainLoop(t *testing.T) {
	g, err := getGrid(sampleGridInputWithStart())
	if err != nil {
		t.Fatal(err)
	}
	expected := []Coord{
		{1, 1}, {1, 2}, {1, 3},
		{2, 1}, {2, 3},
		{3, 1}, {3, 2}, {3, 3},
	}
	sort.Slice(expected, func(i, j int) bool { return expected[i].LessThan(expected[j]) })
	actual := g.GetMainLoop()
	sort.Slice(actual, func(i, j int) bool { return actual[i].LessThan(actual[j]) })
	for i, exp := range expected {
		act := actual[i]
		if act != exp {
			t.Fatalf("expected %v, got %v", exp, act)
		}
	}
}

func TestGridBbox(t *testing.T) {
	g, err := getGrid(sampleGridInputWithStart())
	if err != nil {
		t.Fatal("bruh")
	}
	loop := g.GetMainLoop()
	min, max := g.bbox(loop)
	expectedMin := Coord{1, 1}
	expectedMax := Coord{3, 3}
	if min != expectedMin {
		t.Errorf("[min] expected %v, got %v", expectedMin, min)
	}
	if max != expectedMax {
		t.Errorf("[max] expected %v, got %v", expectedMax, max)
	}
}

func TestGridLoopArea(t *testing.T) {
	inputs := [][]byte{
		sampleGridInputWithStart(),
		[]byte("...........\r\n.S-------7.\r\n.|F-----7|.\r\n.||.....||.\r\n.||.....||.\r\n.|L-7.F-J|.\r\n.|..|.|..|.\r\n.L--J.L--J.\r\n..........."),
		[]byte("..........\r\n.S------7.\r\n.|F----7|.\r\n.||....||.\r\n.||....||.\r\n.|L-7F-J|.\r\n.|..||..|.\r\n.L--JL--J.\r\n.........."),
		[]byte(".F----7F7F7F7F-7....\r\n.|F--7||||||||FJ....\r\n.||.FJ||||||||L7....\r\nFJL7L7LJLJ||LJ.L-7..\r\nL--J.L7...LJS7F-7L7.\r\n....F-J..F7FJ|L7L7L7\r\n....L7.F7||L7|.L7L7|\r\n.....|FJLJ|FJ|F7|.LJ\r\n....FJL-7.||.||||...\r\n....L---J.LJ.LJLJ..."),
		[]byte("FF7FSF7F7F7F7F7F---7\r\nL|LJ||||||||||||F--J\r\nFL-7LJLJ||||||LJL-77\r\nF--JF--7||LJLJ7F7FJ-\r\nL---JF-JLJ.||-FJLJJ7\r\n|F|F-JF---7F7-L7L|7|\r\n|FFJF7L7F-JF7|JL---7\r\n7-L-JL7||F7|L7F-7F7|\r\nL.L7LFJ|||||FJL7||LJ\r\nL7JLJL-JLJLJL--JLJ.L"),
	}
	expected := []int{
		1,
		4,
		4,
		8,
		10,
	}
	if len(inputs) != len(expected) {
		t.Fatalf("bruh")
	}
	for i, inp := range inputs {
		g, err := getGrid(inp)
		if err != nil {
			t.Fatal("bruh")
		}
		exp := expected[i]
		loop := g.GetMainLoop()
		interior := g.LoopInterior(loop)
		if area := len(interior); area != exp {
			t.Errorf("expected area %d, got %d instead\n%s\n", exp, area, g.highlightPath(interior))
		}
	}
}
