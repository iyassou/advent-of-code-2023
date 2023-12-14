package main

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/iyassou/advent-of-code-2023/internal"
)

var mapHeaderRegex = regexp.MustCompile(fmt.Sprintf(`%[1]s-to-%[1]s map:`, entryRegex))

type Map struct {
	From   Entry
	To     Entry
	Ranges []*Range
}

func NewMap(input string) (*Map, error) {
	a := &Map{Ranges: []*Range{}}
	for i, line := range internal.Lines(input) {
		if i == 0 {
			m := mapHeaderRegex.FindAllStringSubmatch(line, -1)
			if len(m) != 1 || len(m[0]) != 3 {
				return nil, errors.New("failed to parse almanac map header")
			}
			a.From = Entry(m[0][1])
			a.To = Entry(m[0][2])
			continue
		}
		if r, err := NewRange(line); err != nil {
			return nil, err
		} else {
			a.Ranges = append(a.Ranges, r)
		}
	}
	return a, nil
}

func (m *Map) Convert(num int) int {
	for _, r := range m.Ranges {
		if r.Contains(num) {
			return r.DestinationStart + (num - r.SourceStart)
		}
	}
	return num
}
