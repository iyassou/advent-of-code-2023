package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func sampleInput() []byte {
	return []byte("O....#....\r\nO.OO#....#\r\n.....##...\r\nOO.#O....O\r\n.O.....O#.\r\nO.#..O.#.#\r\n..O..#O..O\r\n.......O..\r\n#....###..\r\n#OO..#....")
}

func TestNewPlatform(t *testing.T) {
	expected := &Platform{
		Height: 10,
		Width:  10,
		Data:   []byte("O....#....O.OO#....#.....##...OO.#O....O.O.....O#.O.#..O.#.#..O..#O..O.......O..#....###..#OO..#...."),
	}
	if actual, err := NewPlatform(sampleInput()); err != nil {
		t.Fatalf("bruh: %v", err)
	} else if !cmp.Equal(expected, actual) {
		t.Fatalf("expected:\n%v\nnot the same as actual:\n%v", expected, actual)
	}
}

func TestPlatformLinearCoordinate(t *testing.T) {
	p, err := NewPlatform(sampleInput())
	if err != nil {
		t.Fatal(err)
	}
	inputs := [][2]int{
		{0, 0},
		{0, 1},
		{2, 3},
		{9, 9},
	}
	outputs := []int{
		0,
		10,
		32,
		99,
	}
	if len(inputs) != len(outputs) {
		t.Fatal("bruh")
	}
	for i, in := range inputs {
		x, y := in[0], in[1]
		out := outputs[i]
		if actual, err := p.linearCoordinate(x, y); err != nil {
			t.Errorf("[%d] errored: %v", i, err)
		} else if actual != out {
			t.Errorf("[%d] expected %v, got %v", i, out, actual)
		}
	}
	badInputs := [][2]int{
		{-1, -1},
		{p.Width, 0},
		{0, p.Height},
	}
	for i, bad := range badInputs {
		x, y := bad[0], bad[1]
		if v, err := p.linearCoordinate(x, y); err == nil {
			t.Errorf("[%d] expected error, succeeded with %v", i, v)
		}
	}
}

func TestPlatformTiltNorth(t *testing.T) {
	p, err := NewPlatform(sampleInput())
	if err != nil {
		t.Fatal(err)
	}
	expected := &Platform{
		Height: 10,
		Width:  10,
		Data:   []byte("OOOO.#.O..OO..#....#OO..O##..OO..#.OO...........#...#....#.#..O..#.O.O..O.......#....###..#....#...."),
	}
	if err := p.TiltNorth(); err != nil {
		t.Fatal(err)
	} else if !cmp.Equal(p, expected) {
		t.Fatalf("got:\n%v\nexpected:\n%v", p, expected)
	}
}

func TestPlatformTotalNorthLoad(t *testing.T) {
	p, err := NewPlatform(sampleInput())
	if err != nil {
		t.Fatal(err)
	}
	if err := p.TiltNorth(); err != nil {
		t.Fatal(err)
	}
	expected := 136
	actual := p.TotalNorthLoad()
	if expected != actual {
		t.Fatalf("expected %d, got %d", expected, actual)
	}
}

func TestPlatformFacing(t *testing.T) {
	p, err := NewPlatform(sampleInput())
	if err != nil {
		t.Fatal(err)
	}
	dirs := []Orientation{East, South, West, North}
	for _, dir := range dirs {
		p.rotateRight()
		if p.Facing != dir {
			t.Fatalf("platform facing %s, expected %s", p.Facing, dir)
		}
	}
}

func TestPlatformLinearCoordinateWithRotateRight(t *testing.T) {
	p, err := NewPlatform(sampleInput())
	if err != nil {
		t.Fatal(err)
	}
	checks := []map[[2]int]int{
		{
			{0, 0}:                      (p.Height - 1) * p.Width,
			{0, p.Height - 1}:           p.Width - 1 + (p.Height-1)*p.Width,
			{p.Width - 1, 0}:            0,
			{p.Width - 1, p.Height - 1}: p.Width - 1,
		},
		{
			{0, 0}:                      p.Width - 1,
			{0, p.Height - 1}:           0,
			{p.Width - 1, 0}:            p.Width - 1 + (p.Height-1)*p.Width,
			{p.Width - 1, p.Height - 1}: (p.Height - 1) * p.Width,
		},
		{
			{0, 0}:                      p.Width - 1 + (p.Height-1)*p.Width,
			{0, p.Height - 1}:           p.Width - 1,
			{p.Width - 1, 0}:            (p.Height - 1) * p.Width,
			{p.Width - 1, p.Height - 1}: 0,
		},
	}
	for i, values := range checks {
		p.rotateRight()
		if p.Facing != Orientation(i+1) {
			t.Fatalf("bruh: ")
		}
		for coord, expected := range values {
			x, y := coord[0], coord[1]
			if actual, err := p.linearCoordinate(x, y); err != nil {
				t.Errorf("(x,y) = (%d,%d), facing %v: %s", x, y, p.Facing, err)
			} else if actual != expected {
				t.Errorf("(x,y) = (%d,%d), facing %v: expected %d, got %d", x, y, p.Facing, expected, actual)
			}
		}
	}
}

func TestPlatformSpinCycle(t *testing.T) {
	expectedOutputs := [][]byte{
		[]byte(".....#....\r\n....#...O#\r\n...OO##...\r\n.OO#......\r\n.....OOO#.\r\n.O#...O#.#\r\n....O#....\r\n......OOOO\r\n#...O###..\r\n#..OO#...."),
		[]byte(".....#....\r\n....#...O#\r\n.....##...\r\n..O#......\r\n.....OOO#.\r\n.O#...O#.#\r\n....O#...O\r\n.......OOO\r\n#..OO###..\r\n#.OOO#...O"),
		[]byte(".....#....\r\n....#...O#\r\n.....##...\r\n..O#......\r\n.....OOO#.\r\n.O#...O#.#\r\n....O#...O\r\n.......OOO\r\n#...O###.O\r\n#.OOO#...O"),
	}
	for cycle, data := range expectedOutputs {
		if actual, err := NewPlatform(sampleInput()); err != nil {
			t.Fatal(err)
		} else if err := actual.SpinCycle(cycle + 1); err != nil {
			t.Fatal(err)
		} else if expected, err := NewPlatform(data); err != nil {
			t.Fatal(err)
		} else if !cmp.Equal(actual, expected) {
			t.Fatalf("[%d] expected:\n%s\ngot:\n%s", cycle, expected, actual)
		}
	}
	/*if p, err := NewPlatform(sampleInput()); err != nil {
		t.Fatal(err)
	} else if err := p.SpinCycle(1_000_000_000); err != nil {
		t.Fatal(err)
	} else if actual := p.TotalNorthLoad(); actual != 64 {
		t.Fatalf("expected north load of 64, got %d", actual)
	}*/
}
