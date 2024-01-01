package interval

import (
	"math"
	"testing"
)

func FuzzNewInterval(f *testing.F) {
	testcases := []int{0, 1, -1, 3, 12}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, a int) {
		// width=1 interval: valid
		var x, y int
		if a == math.MaxInt {
			x, y = a-1, a
		} else {
			x, y = a, a+1
		}
		if i, err := NewInterval(x, y); err != nil {
			t.Fatalf("[%d, %d] failed with %v", x, y, err)
		} else if s := i.Size(); s != 1 {
			t.Fatalf("[%d, %d] expected size=1, got %d", x, y, s)
		}
		// width=0 interval: valid
		if i, err := NewInterval(x, x); err != nil {
			t.Fatalf("failed for [%d, %d] with %v", x, x, err)
		} else if s := i.Size(); s != 0 {
			t.Fatalf("expected size=0, got %d", s)
		}
		// width=-1 interval: invalid
		x, y = a+1, a
		if a == math.MaxInt {
			x, y = y, x
		}
		if _, err := NewInterval(x, y); err == nil {
			t.Fatalf("expected failure for [%d, %d]", x, y)
		}
	})
}

func FuzzIntervalIntersection(f *testing.F) {
	testcases := []int{-2, -1, 0, 1, math.MaxInt}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, a int) {
		if a > math.MaxInt-3 {
			// I'm testing intervals, not maxints
			return
		}
		// [a, a+1] ∩ [a+1, a+2] = [a+1, a+1]
		x, err := NewInterval(a, a+1)
		if err != nil {
			t.Fatalf("interval(%d, %d) failed with %v", a, a+1, err)
		}
		y, err := NewInterval(a+1, a+2)
		if err != nil {
			t.Fatalf("interval(%d, %d) failed with %v", a+1, a+2, err)
		}
		z, err := NewInterval(a+1, a+1)
		if err != nil {
			t.Fatalf("interval(%d, %d) failed with %v", a+1, a+1, err)
		}
		insect := x.Intersection(y)
		if insect == nil {
			t.Fatalf("[%d, %[2]d] ∩ [%[2]d, %d] returned nil", a, a+1, a+2)
		}
		if z.Start != insect.Start || z.End != insect.End {
			t.Fatalf("expected [%d, %[2]d] ∩ [%[2]d, %d] = [%[2]d, %[2]d], got [%d, %d]", a, a+1, a+2, insect.Start, insect.End)
		}

		// [a, a+1] ∩ [a+2, a+3] = ∅
		x, err = NewInterval(a, a+1)
		if err != nil {
			t.Fatalf("interval(%d, %d) failed with %v", a, a+1, err)
		}
		y, err = NewInterval(a+2, a+3)
		if err != nil {
			t.Fatalf("interval(%d, %d) failed with %v", a+2, a+3, err)
		}
		insect = x.Intersection(y)
		if insect != nil {
			t.Fatalf("expected [%d, %d] ∩ [%d, %d] = ∅, got [%d, %d]", a, a+1, a+2, a+3, insect.Start, insect.End)
		}

		// [a, a+1] ∩ [a, a+2] = [a, a+1]
		x, err = NewInterval(a, a+1)
		if err != nil {
			t.Fatalf("interval(%d, %d) failed with %v", a, a+1, err)
		}
		y, err = NewInterval(a, a+2)
		if err != nil {
			t.Fatalf("interval(%d, %d) failed with %v", a, a+2, err)
		}
		z, err = NewInterval(a, a+1)
		if err != nil {
			t.Fatalf("interval(%d, %d) failed with %v", a, a+1, err)
		}
		insect = x.Intersection(y)
		if insect == nil {
			t.Fatalf("[%[1]d, %d] ∩ [%[1]d, %d] returned nil", a, a+1, a+2)
		}
		if z.Start != insect.Start || z.End != insect.End {
			t.Fatalf("expected [%[1]d, %[2]d] ∩ [%[1]d, %d] = [%[1]d, %[2]d], got [%d, %d]", a, a+1, a+2, insect.Start, insect.End)
		}

		// [a+1, a+2] ∩ [a, a+2] = [a+1, a+2]
		x, err = NewInterval(a+1, a+2)
		if err != nil {
			t.Fatalf("interval(%d, %d) failed with %v", a+1, a+2, err)
		}
		y, err = NewInterval(a, a+2)
		if err != nil {
			t.Fatalf("interval(%d, %d) failed with %v", a, a+2, err)
		}
		z, err = NewInterval(a+1, a+2)
		if err != nil {
			t.Fatalf("interval(%d, %d) failed with %v", a+1, a+2, err)
		}
		insect = x.Intersection(y)
		if insect == nil {
			t.Fatalf("[%d, %[2]d] ∩ [%d, %[2]d] returned nil", a+1, a+2, a)
		}
		if z.Start != insect.Start || z.End != insect.End {
			t.Fatalf("expected [%[1]d, %[2]d] ∩ [%d, %[2]d] = [%[1]d, %[2]d], got [%d, %d]", a+1, a+2, a, insect.Start, insect.End)
		}
	})
}
