package main

import "fmt"

type CamelCard int

const (
	Two   CamelCard = 2
	Three CamelCard = 3
	Four  CamelCard = 4
	Five  CamelCard = 5
	Six   CamelCard = 6
	Seven CamelCard = 7
	Eight CamelCard = 8
	Nine  CamelCard = 9
	Ten   CamelCard = 10
	Jack  CamelCard = 11
	Queen CamelCard = 12
	King  CamelCard = 13
	Ace   CamelCard = 14
)

func ByteToCamelCard(b byte) (c CamelCard, err error) {
	if '2' <= b && b <= '9' {
		return CamelCard(b - 48), nil
	}
	if b == 'T' {
		return Ten, nil
	}
	if b == 'J' {
		return Jack, nil
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
