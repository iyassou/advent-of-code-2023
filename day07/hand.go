package main

import (
	"fmt"
	"strconv"
	"strings"
)

type HandType int

const (
	UnknownHandType HandType = -1
	HighCard        HandType = 1
	OnePair         HandType = 2
	TwoPair         HandType = 3
	ThreeOfAKind    HandType = 4
	FullHouse       HandType = 5
	FourOfAKind     HandType = 6
	FiveOfAKind     HandType = 7
)

type Hand struct {
	Cards [5]CamelCard
	Bid   int
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
	return h, nil
}

func (h *Hand) countCards() map[CamelCard]int {
	freqs := map[CamelCard]int{}
	for _, c := range h.Cards {
		freqs[c]++
	}
	return freqs
}

func (h *Hand) getHandType() HandType {
	freqs := h.countCards()
	if len(freqs) == 0 {
		return UnknownHandType
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

func (h *Hand) getHandTypeFunky() HandType {
	freqs := h.countCards()
	if _, ok := freqs[Joker]; !ok {
		return h.getHandType()
	}
	if len(freqs) == 0 {
		return UnknownHandType
	}
	if len(freqs) <= 2 {
		return FiveOfAKind
	}
	if len(freqs) == 3 {
		for c, f := range freqs {
			if f == 3 || (f == 2 && c == Joker) {
				return FourOfAKind
			}
		}
		return FullHouse
	}
	if len(freqs) == 4 {
		return ThreeOfAKind
	}
	return OnePair
}

func (h *Hand) BetterThan(other *Hand, part int) bool {
	var ht, ot HandType
	if part == 1 {
		ht, ot = h.getHandType(), other.getHandType()
	} else if part == 2 {
		ht, ot = h.getHandTypeFunky(), other.getHandTypeFunky()
	}
	if ht > ot {
		return true
	} else if ht < ot {
		return false
	}
	for i, c := range h.Cards {
		v, ov := c.Value(part), other.Cards[i].Value(part)
		if v > ov {
			return true
		} else if v < ov {
			return false
		}
	}
	return false
}
