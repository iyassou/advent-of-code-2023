package main

import (
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func sampleCardInputs() []string {
	return []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	}
}

func makeCards(inputs []string) []*Card {
	cards := []*Card{}
	for _, s := range inputs {
		if c, err := NewCard(s); err != nil {
			log.Fatal("bruh")
		} else {
			cards = append(cards, c)
		}
	}
	return cards
}

func TestNewCardParsing(t *testing.T) {
	in := sampleCardInputs()
	expected := []*Card{
		{
			ID:             1,
			WinningNumbers: []int{41, 48, 83, 86, 17},
			Numbers:        []int{83, 86, 6, 31, 17, 9, 48, 53},
		},
		{
			ID:             2,
			WinningNumbers: []int{13, 32, 20, 16, 61},
			Numbers:        []int{61, 30, 68, 82, 17, 32, 24, 19},
		},
		{
			ID:             3,
			WinningNumbers: []int{1, 21, 53, 59, 44},
			Numbers:        []int{69, 82, 63, 72, 16, 21, 14, 1},
		},
		{
			ID:             4,
			WinningNumbers: []int{41, 92, 73, 84, 69},
			Numbers:        []int{59, 84, 76, 51, 58, 5, 54, 83},
		},
		{
			ID:             5,
			WinningNumbers: []int{87, 83, 26, 28, 32},
			Numbers:        []int{88, 30, 70, 12, 93, 22, 82, 36},
		},
	}
	if len(in) != len(expected) {
		t.Fatal("bruh")
	}
	for i, line := range in {
		if c, err := NewCard(line); err != nil {
			t.Fatal(err)
		} else if e := expected[i]; !cmp.Equal(c, e) {
			t.Fatalf("expected %v, got %v", e, c)
		}
	}
}

func TestCardIsWinningNumber(t *testing.T) {
	cards := makeCards(sampleCardInputs())
	winningNumbers := [][]int{
		{41, 48, 83, 86, 17},
		{13, 32, 20, 16, 61},
		{1, 21, 53, 59, 44},
		{41, 92, 73, 84, 69},
		{87, 83, 26, 28, 32},
	}
	notWinningNumbers := [][]int{
		{-1, -12331123, 1, 2, 3, 4, 5},
		{777, 12, 31, 19, 15},
		{2, 22, 54, 1000},
		{9, 99, 999, 9999, 99999},
		{-1, -2, -3, -4, -5},
	}
	if len(cards) != len(winningNumbers) && len(cards) != len(notWinningNumbers) {
		t.Fatal("bruh")
	}
	for i, wins := range winningNumbers {
		c, notWins := cards[i], notWinningNumbers[i]
		for _, w := range wins {
			if !c.IsWinningNumber(w) {
				t.Fatalf("expected %d to be winning number for %v", w, c)
			}
		}
		for _, nw := range notWins {
			if c.IsWinningNumber(nw) {
				t.Fatalf("%d should not be a winning number for %v", nw, c)
			}
		}
	}
}

func TestCardPoints(t *testing.T) {
	cards := makeCards(sampleCardInputs())
	points := []int{8, 2, 2, 1, 0}
	if len(points) != len(cards) {
		t.Fatal("bruh")
	}
	for i, c := range cards {
		if p, actual := points[i], c.Points(); p != actual {
			t.Fatalf("expected %d points for %v, got %d", p, c, actual)
		}
	}
}

func TestScratchcardTrackerProcess(t *testing.T) {
	inputs := sampleCardInputs()
	inputs = append(inputs, "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11")
	cards := makeCards(inputs)
	s, err := NewScratchcardTracker()
	if err != nil {
		t.Fatal(err)
	}
	for _, card := range cards {
		s.AddCard(card)
	}
	expected := 30
	actual := s.Total
	if actual != expected {
		t.Fatalf("expected %d, got %d", expected, actual)
	}
}
