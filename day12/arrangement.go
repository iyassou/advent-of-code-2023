package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Arrangement struct {
	Conditions   []byte
	GroupLengths []int
}

func NewArrangement(input []byte) (*Arrangement, error) {
	f := bytes.Fields(input)
	if len(f) != 2 {
		return nil, fmt.Errorf("expected 2 fields, got %d", len(f))
	}
	nums := bytes.Split(f[1], []byte{','})
	a := &Arrangement{
		Conditions:   f[0],
		GroupLengths: make([]int, len(nums)),
	}
	for i, num := range nums {
		if v, err := strconv.Atoi(string(num)); err != nil {
			return nil, err
		} else {
			a.GroupLengths[i] = v
		}
	}
	return a, nil
}

func (a *Arrangement) AllPossibleValues() [][]byte {
	if a == nil {
		return nil
	}
	unknowns := []int{}
	for i, c := range a.Conditions {
		if c == '?' {
			unknowns = append(unknowns, i)
		}
	}
	values := [][]byte{}
	numCombos := 1 << len(unknowns)
	for combo := 0; combo < numCombos; combo++ {
		value := bytes.Clone(a.Conditions)
		for i, unknown := range unknowns {
			if combo&(1<<i) != 0 {
				value[unknown] = '.' // operational
			} else {
				value[unknown] = '#' // damaged
			}
		}
		values = append(values, value)
	}
	return values
}

func (a *Arrangement) MakeRegexp() (*regexp.Regexp, error) {
	if a == nil {
		return nil, fmt.Errorf("bruh")
	}
	groups := make([]string, len(a.GroupLengths))
	for i, length := range a.GroupLengths {
		groups[i] = fmt.Sprintf("#{%d}", length)
	}
	groupsTogether := strings.Join(groups, `\.+`)
	return regexp.Compile(groupsTogether)
}

func (a *Arrangement) BruteForce() int {
	if a == nil {
		return -1
	}
	values := a.AllPossibleValues()
	regex, err := a.MakeRegexp()
	if err != nil {
		return -1
	}
	matches := 0
	for _, v := range values {
		loc := regex.FindIndex(v)
		if loc == nil {
			continue
		}
		bruh := false
		for i := 0; i < loc[0]; i++ {
			if v[i] == '#' {
				bruh = true
				break
			}
		}
		if bruh {
			continue
		}
		for i := loc[1]; i < len(v); i++ {
			if v[i] != '.' {
				bruh = true
				break
			}
		}
		if bruh {
			continue
		}
		matches++
	}
	return matches
}
