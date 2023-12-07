package main

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func sampleInput() string {
	return "Time:      7  15   30\r\nDistance:  9  40  200"
}

func TestCompetitionParsing(t *testing.T) {
	c, err := NewCompetition(sampleInput())
	if err != nil {
		t.Fatal(err)
	}
	expected := &Competition{
		Races: []*Race{
			{Time: 7 * time.Millisecond, Distance: 9},
			{Time: 15 * time.Millisecond, Distance: 40},
			{Time: 30 * time.Millisecond, Distance: 200},
		},
	}
	if !cmp.Equal(c, expected) {
		t.Fatalf("expected %v, got %v", expected, c)
	}
}

func TestCompetitionProductWaysOfWinning(t *testing.T) {
	c, err := NewCompetition(sampleInput())
	if err != nil {
		t.Fatal(err)
	}
	expected := 288
	if p := c.ProductWaysOfWinning(); p != expected {
		t.Fatalf("expected %d, got %d instead", expected, p)
	}
}

func TestRaceParsing(t *testing.T) {
	r, err := NewRace(sampleInput())
	if err != nil {
		t.Fatal(err)
	}
	expected := &Race{
		Time: 71530 * time.Millisecond, Distance: 940200,
	}
	if !cmp.Equal(r, expected) {
		t.Fatalf("expected %v, got %v", expected, r)
	}
}

func TestRaceWaysOfWinning(t *testing.T) {
	r, err := NewRace(sampleInput())
	if err != nil {
		t.Fatal(err)
	}
	expected := 71503
	if w := r.WaysOfBreakingTheRecord(); w != expected {
		t.Fatalf("expected %d ways of winning, got %d instead", expected, w)
	}
}
