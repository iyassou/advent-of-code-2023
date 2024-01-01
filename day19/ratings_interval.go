package main

import "github.com/iyassou/advent-of-code-2023/internal/interval"

var startDefault = 1
var endDefault = 4000

type ratingsInterval map[string]*interval.Interval

func newRatingsInterval() ratingsInterval {
	r := ratingsInterval{}
	for _, cat := range categories {
		r[cat], _ = interval.NewInterval(startDefault, endDefault)
	}
	return r
}

func (r ratingsInterval) copy() ratingsInterval {
	other := ratingsInterval{}
	for _, cat := range categories {
		ival := r[cat]
		other[cat], _ = interval.NewInterval(ival.Start, ival.End)
	}
	return other
}

func (r ratingsInterval) distinctCombinations() int {
	prod := 1
	for _, i := range r {
		prod *= (i.Size() + 1)
	}
	return prod
}
