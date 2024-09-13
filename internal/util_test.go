package internal

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func FuzzShortestRepeatingSubstring(f *testing.F) {
	testcases := []string{"catcatcat", "catdog", "dogcatdogcat", "ccccccccc", ""}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		if !utf8.ValidString(input) {
			t.Skip("skipping invalid utf8")
		}
		sub := ShortestRepeatingSubstring(input)
		if len(input) < 2 {
			if input != sub {
				t.Fatalf("expected %[1]q for %[1]q, got %[2]q", input, sub)
			}
			return
		}
		times := len(input) / len(sub)
		repeated := strings.Repeat(sub, times)
		if repeated != input {
			t.Fatalf("substring %q times %d != input %q", sub, times, input)
		}
	})
}

func TestGCD(t *testing.T) {
	testcases := [][3]int{
		{20, 15, 5},
		{15, 20, 5},
		{31, 27, 1},
	}
	for _, tc := range testcases {
		a, b, expected := tc[0], tc[1], tc[2]
		if actual, err := GCD(a, b); err != nil {
			t.Fatal(err)
		} else if actual != expected {
			t.Fatalf("expected GCD(%d, %d) = %d, got %d", a, b, expected, actual)
		}
	}
}

func TestLCM(t *testing.T) {
	testcases := [][3]int{
		{20, 15, 60},
		{15, 20, 60},
		{31, 27, 31 * 27},
	}
	for _, tc := range testcases {
		a, b, expected := tc[0], tc[1], tc[2]
		if actual, err := LCM(a, b); err != nil {
			t.Fatal(err)
		} else if actual != expected {
			t.Fatalf("expected LCM(%d, %d) = %d, got %d", a, b, expected, actual)
		}
	}
}
