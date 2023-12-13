package main

import "testing"

func TestTileCanFitIn(t *testing.T) {
	chars := []byte{'|', '-', 'L', 'J', '7', 'F', '.', 'S'}
	tiles := make([]*Tile, len(chars))
	for i, c := range chars {
		tiles[i] = NewTile(c)
	}
	dirs := []Direction{North, South, West, East}
	expected := [][4]bool{
		{true, true, false, false},
		{false, false, true, true},
		{false, true, true, false},
		{false, true, false, true},
		{true, false, false, true},
		{true, false, true, false},
		{false, false, false, false},
		{true, true, true, true},
	}
	if len(expected) != len(chars) {
		t.Fatal("bruh")
	}
	for i, tile := range tiles {
		for j, dir := range dirs {
			e := expected[i][j]
			a := tile.CanFitIn(dir)
			if a != e {
				t.Fatalf("direction %s, expected tile %v fit? %t, but got %t instead", dir, string(tile.Char), e, a)
			}
		}
	}
}
