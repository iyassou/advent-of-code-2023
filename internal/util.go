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
