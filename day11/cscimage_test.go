package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func sampleImage() []byte {
	return []byte("...#......\r\n.......#..\r\n#.........\r\n..........\r\n......#...\r\n.#........\r\n.........#\r\n..........\r\n.......#..\r\n#...#.....")
}

func sampleImageExpanded() []byte {
	return []byte("....#........\r\n.........#...\r\n#............\r\n.............\r\n.............\r\n........#....\r\n.#...........\r\n............#\r\n.............\r\n.............\r\n.........#...\r\n#....#.......")
}

func TestNewCSCImage(t *testing.T) {
	if _, err := NewCSCImage(sampleImage()); err != nil {
		t.Fatal(err)
	}
}

func TestCSCImageNumColumns(t *testing.T) {
	inputs := [][]byte{
		sampleImage(), sampleImageExpanded(),
	}
	expected := []int{10, 13}
	if len(inputs) != len(expected) {
		t.Fatal("ruh roh")
	}
	for i, inp := range inputs {
		c, err := NewCSCImage(inp)
		if err != nil {
			t.Error(err)
		}
		exp := expected[i]
		actual := c.NumColumns()
		if exp != actual {
			t.Errorf("[%d] %v\n", i, c.ColumnIndex)
			t.Errorf("[%d] expected %d, got %d", i, exp, actual)
		}
	}
}

func TestCSCImageGalaxyCount(t *testing.T) {
	inputs := [][]byte{
		sampleImage(), sampleImageExpanded(),
	}
	expected := []int{9, 9}
	if len(inputs) != len(expected) {
		t.Fatal("wesh")
	}
	for i, inp := range inputs {
		c, err := NewCSCImage(inp)
		if err != nil {
			t.Error(err)
		}
		exp := expected[i]
		actual := c.galaxyCount()
		if exp != actual {
			t.Errorf("[%d] %v\n", i, c.ColumnIndex)
			t.Errorf("[%d] expected %d, got %d", i, exp, actual)
		}
	}
}

func TestCSCImageDimensions(t *testing.T) {
	inputs := [][]byte{
		sampleImage(),
		sampleImageExpanded(),
	}
	expected := [][2]int{
		{10, 10},
		{12, 13},
	}
	if len(inputs) != len(expected) {
		t.Fatal("reuf")
	}
	for i, inp := range inputs {
		exp := expected[i]
		c, err := NewCSCImage(inp)
		if err != nil {
			t.Fatal(err)
		}
		expectedM, expectedN := exp[0], exp[1]
		M, N := c.dimensions()
		if expectedM != M {
			t.Errorf("[%d] expected %d rows, got %d", i, expectedM, M)
		}
		if expectedN != N {
			t.Errorf("[%d] expected %d columns, got %d", i, expectedN, N)
		}
	}
}

func TestCSCImageGalaxyAt(t *testing.T) {
	// PROTIP: search and replace for the expected bools
	inputs := [][]byte{
		sampleImage(), sampleImageExpanded(),
	}
	expectedValues := [][][]bool{
		{
			{false, false, false, true, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, true, false, false},
			{true, false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, true, false, false, false},
			{false, true, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false, true},
			{false, false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, true, false, false},
			{true, false, false, false, true, false, false, false, false, false},
		},
		{
			{false, false, false, false, true, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false, true, false, false, false},
			{true, false, false, false, false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, true, false, false, false, false},
			{false, true, false, false, false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false, false, false, false, true},
			{false, false, false, false, false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false, true, false, false, false},
			{true, false, false, false, false, true, false, false, false, false, false, false, false},
		},
	}
	if len(inputs) != len(expectedValues) {
		t.Fatal("b r u h")
	}
	for i, inp := range inputs {
		expected := expectedValues[i]
		c, err := NewCSCImage(inp)
		if err != nil {
			t.Fatal(err)
		}
		M, N := c.dimensions()
		prefix := func(a, b int) string { return fmt.Sprintf("[%d | (%d,%d)]", i, a, b) }
		for x := 0; x < M; x++ {
			for y := 0; y < N; y++ {
				exp := expected[x][y]
				if actual, err := c.galaxyAt(x, y); err != nil {
					t.Errorf("%s %v", prefix(x, y), err)
				} else if actual != exp {
					t.Errorf("%s expected %t, got %t", prefix(x, y), exp, actual)
				}
			}
		}
		if _, err := c.galaxyAt(0, N); err == nil {
			t.Errorf("%s expected error, passed", prefix(0, N))
		}
		if _, err := c.galaxyAt(M, 0); err == nil {
			t.Errorf("%s expected error, passed", prefix(M, 0))
		}
		if _, err := c.galaxyAt(M, N); err == nil {
			t.Errorf("%s expected error, passed", prefix(M, N))
		}
	}
}

func TestCSCImageEmptyRows(t *testing.T) {
	inputs := [][]byte{
		sampleImage(),
		sampleImageExpanded(),
	}
	expected := [][]int{
		{3, 7},
		{3, 4, 8, 9},
	}
	if len(inputs) != len(expected) {
		t.Fatal("ወንድም")
	}
	for i, inp := range inputs {
		exp := expected[i]
		c, err := NewCSCImage(inp)
		if err != nil {
			t.Fatal(err)
		}
		actual := c.emptyRows()
		if !cmp.Equal(exp, actual) {
			t.Errorf("[%d] expected %v, got %v", i, exp, actual)
		}
	}
}

func TestCSCImageEmptyColumns(t *testing.T) {
	inputs := [][]byte{
		sampleImage(),
		sampleImageExpanded(),
	}
	expected := [][]int{
		{2, 5, 8},
		{2, 3, 6, 7, 10, 11},
	}
	if len(inputs) != len(expected) {
		t.Fatal("bro")
	}
	for i, inp := range inputs {
		exp := expected[i]
		c, err := NewCSCImage(inp)
		if err != nil {
			t.Fatal(err)
		}
		actual := c.emptyColumns()
		if !cmp.Equal(exp, actual) {
			t.Errorf("[%d] expected %v, got %v", i, exp, actual)
		}
	}
}

func TestCSCImageExpandEmptyRowsBy1(t *testing.T) {
	c, err := NewCSCImage(sampleImage())
	if err != nil {
		t.Fatal(err)
	}
	c.expandEmptyRows(1)
	expected, err := NewCSCImage([]byte("...#......\r\n.......#..\r\n#.........\r\n..........\r\n..........\r\n......#...\r\n.#........\r\n.........#\r\n..........\r\n..........\r\n.......#..\r\n#...#....."))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(c, expected) {
		t.Fatalf("rows:%v\nexpected:\n%v\nactual:\n%v", c.emptyRows(), expected, c)
	}
}

func TestCSCImageExpandEmptyColumnsBy1(t *testing.T) {
	c, err := NewCSCImage(sampleImage())
	if err != nil {
		t.Fatal(err)
	}
	c.expandEmptyColumns(1)
	expected, err := NewCSCImage([]byte("....#........\r\n.........#...\r\n#............\r\n.............\r\n........#....\r\n.#...........\r\n............#\r\n.............\r\n.........#...\r\n#....#......."))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(c, expected) {
		t.Fatalf("columns:%v\nexpected:\n%v\nactual:\n%v", c.emptyColumns(), expected, c)
	}
}

func TestCSCImageExpandBothBy1(t *testing.T) {
	by := 1
	rowsFirst := []bool{true, false}
	for _, rowThenCol := range rowsFirst {
		c, err := NewCSCImage(sampleImage())
		if err != nil {
			t.Fatal(err)
		}
		if rowThenCol {
			c.expandEmptyRows(by)
			c.expandEmptyColumns(by)
		} else {
			c.expandEmptyColumns(by)
			c.expandEmptyRows(by)
		}
		expected, err := NewCSCImage(sampleImageExpanded())
		if err != nil {
			t.Fatal(err)
		}
		if !cmp.Equal(c, expected) {
			t.Errorf("[rowThenCol? %t] expected:\n%v\nactual:\n%v", rowThenCol, expected, c)
		}
	}
}

func TestCSCImageGalaxyLocations(t *testing.T) {
	inputs := [][]byte{
		sampleImage(),
		// sampleImageExpanded(),
	}
	expected := [][][2]int{
		{
			{2, 0}, {9, 0},
			{5, 1},
			{0, 3},
			{9, 4},
			{4, 6},
			{1, 7}, {8, 7},
			{6, 9},
		},
	}
	if len(inputs) != len(expected) {
		t.Fatal("አንበሳ")
	}
	for i, inp := range inputs {
		exp := expected[i]
		c, err := NewCSCImage(inp)
		if err != nil {
			t.Fatal(err)
		}
		actual := c.GalaxyLocations()
		if !cmp.Equal(exp, actual) {
			t.Errorf("[%d] expected %v, got %v", i, exp, actual)
		}
	}
}
