package main

import "fmt"

type CamelCard int

const (
	Two CamelCard = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Joker
	Queen
	King
	Ace
)

func (c CamelCard) Value(part int) int {
	if part == 1 {
		if c == Two {
			return 2
		}
		if c == Three {
			return 3
		}
		if c == Four {
			return 4
		}
		if c == Five {
			return 5
		}
		if c == Six {
			return 6
		}
		if c == Seven {
			return 7
		}
		if c == Eight {
			return 8
		}
		if c == Nine {
			return 9
		}
		if c == Ten {
			return 10
		}
		if c == Joker {
			return 11
		}
		if c == Queen {
			return 12
		}
		if c == King {
			return 13
		}
		if c == Ace {
			return 14
		}
		return -1
	} else if part == 2 {
		if c == Joker {
			return 1
		}
		return c.Value(1)
	}
	return -1
}

func ByteToCamelCard(b byte) (c CamelCard, err error) {
	if b == '2' {
		return Two, nil
	}
	if b == '3' {
		return Three, nil
	}
	if b == '4' {
		return Four, nil
	}
	if b == '5' {
		return Five, nil
	}
	if b == '6' {
		return Six, nil
	}
	if b == '7' {
		return Seven, nil
	}
	if b == '8' {
		return Eight, nil
	}
	if b == '9' {
		return Nine, nil
	}
	if b == 'T' {
		return Ten, nil
	}
	if b == 'J' {
		return Joker, nil
	}
	if b == 'Q' {
		return Queen, nil
	}
	if b == 'K' {
		return King, nil
	}
	if b == 'A' {
		return Ace, nil
	}
	return 0, fmt.Errorf("cannot make CamelCard from %v", b)
}
