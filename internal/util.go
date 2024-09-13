package internal

import (
	"bytes"
	"errors"
	"strings"
)

type line interface {
	string | []byte
}

var newlineString = "\r\n"
var newlineBytes = []byte{'\r', '\n'}
var ErrNegativeNumbers = errors.New("expected positive numbers")
var ErrInsufficientArgs = errors.New("expected at least 2 numbers")

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

func gcd(a, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return 0, ErrNegativeNumbers
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a, nil
}

func GCD(args ...int) (int, error) {
	if len(args) < 2 {
		return 0, ErrInsufficientArgs
	}
	ans, err := gcd(args[0], args[1])
	if err != nil {
		return 0, err
	}
	for i := 2; i < len(args); i++ {
		ans, err = gcd(ans, args[i])
		if err != nil {
			return 0, err
		}
	}
	return ans, nil
}

func lcm(a, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return 0, ErrNegativeNumbers
	}
	if gcd, err := GCD(a, b); err != nil {
		return 0, err
	} else {
		return a * b / gcd, nil
	}
}

func LCM(args ...int) (int, error) {
	if len(args) < 2 {
		return 0, ErrInsufficientArgs
	}
	ans, err := lcm(args[0], args[1])
	if err != nil {
		return 0, err
	}
	for i := 2; i < len(args); i++ {
		ans, err = lcm(ans, args[i])
		if err != nil {
			return 0, err
		}
	}
	return ans, nil
}
