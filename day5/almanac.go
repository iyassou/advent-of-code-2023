package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var almanacSeedsRegex = regexp.MustCompile(`seeds: ([\d|\s]+)`)

type Almanac struct {
	Seeds []int
	Maps  []*Map
}

func NewAlmanac(chunks []string) (*Almanac, error) {
	a := &Almanac{}
	if len(chunks) == 0 {
		return a, nil
	}
	m := almanacSeedsRegex.FindStringSubmatch(chunks[0])
	if len(m) != 2 {
		return nil, fmt.Errorf("invalid seeds")
	}
	for _, seed := range strings.Fields(m[1]) {
		if s, err := strconv.Atoi(seed); err != nil {
			return nil, err
		} else {
			a.Seeds = append(a.Seeds, s)
		}
	}
	if len(a.Seeds)%2 != 0 {
		return nil, fmt.Errorf("seeds must come in pairs")
	}
	a.Maps = []*Map{}
	for _, chunk := range chunks[1:] {
		if m, err := NewMap(chunk); err != nil {
			return nil, err
		} else {
			if l := len(a.Maps); l != 0 {
				to, from := a.Maps[l-1].To, m.From
				if to != from {
					return nil, fmt.Errorf("expected map to go from %s, not %s", to, from)
				}
			}
			a.Maps = append(a.Maps, m)
		}
	}
	return a, nil
}

func (a *Almanac) LocationNumber(seed int) int {
	for _, m := range a.Maps {
		seed = m.Convert(seed)
	}
	return seed
}

func (a *Almanac) MinLocationNumber(part int) (int, error) {
	if part != 1 && part != 2 {
		return 0, fmt.Errorf("part must be 1 or 2, not %d", part)
	}
	min := math.MaxInt
	if part == 1 {
		for _, seed := range a.Seeds {
			if loc := a.LocationNumber(seed); loc < min {
				min = loc
			}
		}
	} else {
		for i := 0; i < len(a.Seeds); i += 2 {
			start, length := a.Seeds[i], a.Seeds[i+1]
			for seed := start; seed < start+length; seed++ {
				if loc := a.LocationNumber(seed); loc < min {
					min = loc
				}
			}
		}
	}
	return min, nil
}
