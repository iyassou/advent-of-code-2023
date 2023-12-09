package internal

import (
	"bytes"
	"strings"
)

type line interface {
	string | []byte
}

var newlineString = "\r\n"
var newlineBytes = []byte{'\r', '\n'}

func Lines[T line](input T) []T {
	result := []T{}
	switch input := any(input).(type) {
	case string:
		for _, s := range strings.Split(input, newlineString) {
			result = append(result, T(s))
		}
	case []byte:
		for _, b := range bytes.Split(input, newlineBytes) {
			result = append(result, T(b))
		}
	}
	return result
}

func ShortestRepeatingSubstring(input string) string {
	if len(input) < 2 {
		return input
	}
	runes := []rune(input)
	j, shortest := 0, []rune{runes[0]}
	for i, r := range runes {
		if r == shortest[j] {
			j = (j + 1) % len(shortest)
		} else {
			j = 0
			shortest = runes[:i+1]
		}
	}
	if j == 0 {
		return string(shortest)
	}
	return input
}
