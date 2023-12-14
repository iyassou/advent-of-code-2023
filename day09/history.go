package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type History []int
type Direction bool

const (
	Forward  Direction = true
	Backward Direction = false
)

func NewHistory(line string) (History, error) {
	h := History{}
	for _, field := range strings.Fields(line) {
		if v, err := strconv.Atoi(field); err != nil {
			return h, err
		} else {
			h = append(h, v)
		}
	}
	return h, nil
}

func (h History) Extrapolate(dir Direction) (int, error) {
	if len(h) == 0 {
		return 0, errors.New("cannot extrapolate empty history")
	}
	if !(dir == Forward || dir == Backward) {
		return 0, fmt.Errorf("invalid direction %v", dir)
	}
	diffs := make([]int, len(h))
	copy(diffs, h)
	trunc := len(h)
	prediction := 0
	updateDiffs := func() {
		for i, v := range diffs[1:trunc] {
			diffs[i] = v - diffs[i]
		}
	}
	updatePrediction := func() { prediction += diffs[trunc-1] }
	if dir == Backward {
		for i, j := 0, len(diffs)-1; i < j; i, j = i+1, j-1 {
			diffs[i], diffs[j] = diffs[j], diffs[i]
		}
	}
	keepGoing := func() bool {
		for _, v := range diffs[:trunc] {
			if v != 0 {
				return true
			}
		}
		return false
	}
	for ; keepGoing(); trunc-- {
		updatePrediction()
		updateDiffs()
	}
	return prediction, nil
}
