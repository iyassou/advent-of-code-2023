package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Range struct {
	DestinationStart int
	SourceStart      int
	Length           int
}

var rangeRegex = regexp.MustCompile(`(\d+) (\d+) (\d+)`)

func NewRange(line string) (*Range, error) {
	m := rangeRegex.FindAllStringSubmatch(line, -1)
	if len(m) != 1 || len(m[0]) != 4 {
		return nil, errors.New("failed to parse almanac map range conversion")
	}
	a := &Range{}
	if v, err := strconv.Atoi(m[0][1]); err != nil {
		return nil, err
	} else {
		a.DestinationStart = v
	}
	if v, err := strconv.Atoi(m[0][2]); err != nil {
		return nil, err
	} else {
		a.SourceStart = v
	}
	if v, err := strconv.Atoi(m[0][3]); err != nil {
		return nil, err
	} else {
		a.Length = v
	}
	return a, nil
}

func (r *Range) String() string {
	if r == nil {
		return ""
	}
	return fmt.Sprintf("(%d => %d | %d)", r.SourceStart, r.DestinationStart, r.Length)
}

func (r *Range) Contains(num int) bool {
	return r.SourceStart <= num && num < r.SourceStart+r.Length
}
