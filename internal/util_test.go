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
