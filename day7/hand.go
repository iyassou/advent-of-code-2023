package main

import (
	"fmt"
	"strconv"
	"strings"
)

type HandType int

const (
	HighCard     HandType = 1
	OnePair      HandType = 2
	TwoPair      HandType = 3
	ThreeOfAKind HandType = 4
	FullHouse    HandType = 5
	FourOfAKind  HandType = 6
	FiveOfAKind  HandType = 7
)

type HandStrength int

type Hand struct {
	Cards [5]CamelCard
	Bid   int
	Type  HandType
}

func getHandType(h *Hand) HandType {
	freqs := map[CamelCard]int{}
	for _, c := range h.Cards {
		freqs[c]++
	}
	if len(freqs) == 1 {
		return FiveOfAKind
	}
	if len(freqs) == 2 {
		for _, f := range freqs {
			if f == 4 || f == 1 {
				return FourOfAKind
			}
			return FullHouse
		}
	}
	if len(freqs) == 3 {
		for _, f := range freqs {
			if f == 3 {
				return ThreeOfAKind
			}
			if f == 2 {
				return TwoPair
			}
		}
	}
	if len(freqs) == 4 {
		return OnePair
	}
	return HighCard
}

func NewHand(input string) (*Hand, error) {
	fields := strings.Fields(input)
	if l := len(fields); l != 2 {
		return nil, fmt.Errorf("expected 2 fields, read %d", l)
	}
	if l := len(fields[0]); l != 5 {
		return nil, fmt.Errorf("expected 5 CamelCards in hand, read %d", l)
	}
	h := &Hand{Cards: [5]CamelCard{}}
	for i := range fields[0] {
		b := fields[0][i]
		if c, err := ByteToCamelCard(b); err != nil {
			return nil, err
		} else {
			h.Cards[i] = c
		}
	}
	if b, err := strconv.Atoi(fields[1]); err != nil {
		return nil, err
	} else {
		h.Bid = b
	}
	h.Type = getHandType(h)
	return h, nil
}

func (h *Hand) BetterThan(other *Hand) bool {
	ht, ot := h.Type, other.Type
	if ht > ot {
		return true
	} else if ht < ot {
		return false
	}
	for i, hc := range h.Cards {
		oc := other.Cards[i]
		if hc > oc {
			return true
		} else if hc < oc {
			return false
		}
	}
	return false
}
