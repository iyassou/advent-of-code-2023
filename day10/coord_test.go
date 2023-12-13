package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func FuzzCoordAdd(f *testing.F) {
	testcases := [][]byte{
		{0, 0}, {1, 1},
	}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input []byte) {
		if len(input) != 2 {
			return
		}
		c1 := Coord{int(input[0]), int(input[1])}
		x, y := c1[0], c1[1]
		a, b := x, y
		c2 := c1.Add(c1)
		if c2[0] != x+a {
			t.Fatalf("expected %d for x component, got %d instead", x+a, c2[0])
		}
		if c2[1] != y+b {
			t.Fatalf("expected %d for y component, got %d instead", y+b, c2[1])
		}
		c3 := c2.Add(c1)
		c4 := c1.Add(c2)
		if !cmp.Equal(c3, c4) {
			t.Fatalf("addition not commutative for c1=%v, c2=%v (c3=%v, c4=%v)", c1, c2, c3, c4)
		}
	})
}
