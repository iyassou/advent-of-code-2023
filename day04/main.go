package main

import (
	"bytes"
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var cardRegex = regexp.MustCompile(`Card\s+(\d+):([\d|\s]+)\|([\d|\s]+)`)

type CardID int

type Card struct {
	ID             CardID
	WinningNumbers []int
	Numbers        []int
}

func NewCard(line string) (*Card, error) {
	m := cardRegex.FindAllStringSubmatch(line, -1)
	if len(m) != 1 {
		return nil, errors.New("could not match Card regex")
	}
	if len(m[0]) != 4 {
		return nil, errors.New("could not match all Card regex groups")
	}
	c := &Card{
		WinningNumbers: []int{},
		Numbers:        []int{},
	}
	if v, err := strconv.Atoi(m[0][1]); err != nil {
		return nil, err
	} else {
		c.ID = CardID(v)
	}
	for _, s := range strings.Fields(m[0][2]) {
		if v, err := strconv.Atoi(s); err != nil {
			return nil, err
		} else {
			c.WinningNumbers = append(c.WinningNumbers, v)
		}
	}
	for _, s := range strings.Fields(m[0][3]) {
		if v, err := strconv.Atoi(s); err != nil {
			return nil, err
		} else {
			c.Numbers = append(c.Numbers, v)
		}
	}
	return c, nil
}

func (c *Card) IsWinningNumber(num int) bool {
	for _, w := range c.WinningNumbers {
		if w == num {
			return true
		}
	}
	return false
}

func (c *Card) Wins() int {
	wins := 0
	for _, n := range c.Numbers {
		if c.IsWinningNumber(n) {
			wins++
		}
	}
	return wins
}

func (c *Card) Points() int {
	wins := c.Wins()
	if wins > 0 {
		return 1 << (wins - 1)
	}
	return 0
}

type ScratchcardTracker struct {
	Counts map[CardID]int
	Total  int
}

func NewScratchcardTracker() (*ScratchcardTracker, error) {
	return &ScratchcardTracker{Counts: map[CardID]int{}}, nil
}

func (s *ScratchcardTracker) AddCard(card *Card) {
	id := card.ID
	copies := s.Counts[id] + 1
	s.Total += copies
	wins := CardID(card.Wins())
	for j := id + 1; j < id+1+wins; j++ {
		s.Counts[j] += copies
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: <program> <input.txt>")
	}
	if b, err := os.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else if s, err := NewScratchcardTracker(); err != nil {
		log.Fatal(err)
	} else {
		one := 0
		lines := bytes.Split(b, []byte{'\r', '\n'})
		for _, line := range lines {
			if c, err := NewCard(string(line)); err != nil {
				log.Fatal(err)
			} else {
				one += c.Points()
				s.AddCard(c)
			}
		}
		log.Println("Part One:", one)
		log.Println("Part Two:", s.Total)
	}
}
