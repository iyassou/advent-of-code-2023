package main

import "testing"

func TestPart1(t *testing.T) {
	expected := 6440
	input := "32T3K 765\r\nT55J5 684\r\nKK677 28\r\nKTJJT 220\r\nQQQJA 483"
	if w, err := solution(input, 1); err != nil {
		t.Fatal(err)
	} else if w != expected {
		t.Fatalf("expected to win %d, won %d instead", expected, w)
	}
}

func TestPart2(t *testing.T) {
	expected := 5905
	input := "32T3K 765\r\nT55J5 684\r\nKK677 28\r\nKTJJT 220\r\nQQQJA 483"
	if w, err := solution(input, 2); err != nil {
		t.Fatal(err)
	} else if w != expected {
		t.Fatalf("expected to win %d, won %d instead", expected, w)
	}
}
