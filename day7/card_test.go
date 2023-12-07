package main

import "testing"

func TestByteToCamelCard(t *testing.T) {
	s := "23456789TJQKA"
	expected := []CamelCard{
		Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace,
	}
	if len(s) != len(expected) {
		t.Fatal("bruh")
	}
	for i := range s {
		b := s[i]
		e := expected[i]
		if a, err := ByteToCamelCard(b); err != nil {
			t.Fatalf("expected success, failed with %v", err)
		} else if a != e {
			t.Fatalf("expected %v for input %v, got %v instead", e, b, a)
		}
	}

	s = "woijfoiejfir01010101010101010101jfsdnfgoiwebgpohpur1[fn;elfkndf]"
	for i := range s {
		b := s[i]
		if c, err := ByteToCamelCard(b); err == nil {
			t.Fatalf("expected failure for input %v, succeeded with value %v instead", b, c)
		}
	}
}
