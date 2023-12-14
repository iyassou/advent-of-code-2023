package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewGame(t *testing.T) {
	in := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	expected := Game{
		ID: 3,
		Reveals: []Subset{
			{Red: 20, Green: 8, Blue: 6},
			{Red: 4, Green: 13, Blue: 5},
			{Red: 1, Green: 5},
		},
	}
	if g, err := NewGame(in); err != nil {
		t.Fatal(err)
	} else {
		if !cmp.Equal(*g, expected) {
			t.Fatalf("=> expected\n%v\n=> got\n%v\n", expected, g)
		}
	}
}

func TestSubsetPower(t *testing.T) {
	testValues := map[Subset]int{
		{}:                          0,
		{Red: 1, Green: 1, Blue: 1}: 1,
		{Red: 1}:                    0,
		{Red: 2, Green: 3, Blue: 4}: 24,
		{Red: 6, Green: 4, Blue: 5}: 120,
	}
	for s, expected := range testValues {
		if p := s.Power(); p != expected {
			t.Fatalf("expected %d for %q, got %d instead\n", expected, s, p)
		}
	}
}
