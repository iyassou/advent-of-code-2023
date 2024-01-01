package interval

import "fmt"

type Interval struct {
	Start, End int
}

func NewInterval(start, end int) (*Interval, error) {
	if end < start {
		return nil, fmt.Errorf("invalid interval: [%d, %d]", start, end)
	}
	return &Interval{start, end}, nil
}

func (i *Interval) Size() int {
	return i.End - i.Start
}

func (i *Interval) Intersection(other *Interval) *Interval {
	// x = max(i.Start, other.Start)
	// y = min(i.End, other.End)
	x := i.Start
	if other.Start > x {
		x = other.Start
	}
	y := i.End
	if other.End < y {
		y = other.End
	}
	if x > y {
		// no intersection
		return nil
	}
	intersection, _ := NewInterval(x, y)
	return intersection
}

func (i *Interval) String() string {
	return fmt.Sprintf("[%d, %d]", i.Start, i.End)
}
