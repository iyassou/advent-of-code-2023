package main

import (
	"testing"
)

func sampleInputs() []string {
	return []string{
		"32T3K 765", "T55J5 684", "KK677 28", "KTJJT 220", "QQQJA 483",
	}
}

func sampleHands() ([]*Hand, error) {
	inputs := sampleInputs()
	hands := make([]*Hand, len(inputs))
	for i, s := range inputs {
		if h, err := NewHand(s); err != nil {
			return nil, err
		} else {
			hands[i] = h
		}
	}
	return hands, nil
}

func TestHandParsing(t *testing.T) {
	if _, err := sampleHands(); err != nil {
		t.Fatal(err)
	}
}

func TestGetHandType(t *testing.T) {
	hands, err := sampleHands()
	if err != nil {
		t.Fatal(err)
	}
	expected := []HandType{
		OnePair, ThreeOfAKind, TwoPair, TwoPair, ThreeOfAKind,
	}
	if len(hands) != len(expected) {
		t.Fatal("bruh")
	}
	for i, h := range hands {
		e := expected[i]
		if a := getHandType(h); a != e {
			t.Fatalf("expected %v for input %v, got %v instead", e, h, a)
		}
	}
	inputs := []string{"23456 0", "TTTQQ 0", "8888J 0", "AAAAA 0"}
	expected = []HandType{
		HighCard, FullHouse, FourOfAKind, FiveOfAKind,
	}
	if len(inputs) != len(expected) {
		t.Fatal("bruh")
	}
	for j, inp := range inputs {
		h, err := NewHand(inp)
		if err != nil {
			t.Fatal(err)
		}
		e := expected[j]
		if a := getHandType(h); a != e {
			t.Fatalf("expected %v for input %v, got %v instead", e, h, a)
		}
	}
}
