package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iyassou/advent-of-code-2023/internal/interval"
)

type rule struct {
	comparand int
	category  string
	lessThan  bool
	tautology bool
	iftrue    string
}

func newRule(line string) (*rule, error) {
	parts := strings.Split(line, ":")
	r := &rule{}
	if len(parts) == 1 {
		r.tautology = true
		r.iftrue = parts[0]
	} else if len(parts) == 2 {
		var operator string
		if strings.Contains(parts[0], "<") {
			r.lessThan = true
			operator = "<"
		} else {
			operator = ">"
		}
		vals := strings.Split(parts[0], operator)
		if len(vals) != 2 {
			return nil, fmt.Errorf("invalid predicate %v", parts[0])
		}
		r.category = vals[0]
		if val, err := strconv.Atoi(vals[1]); err != nil {
			return nil, err
		} else {
			r.comparand = val
		}
		r.iftrue = parts[1]
	}
	return r, nil
}

func (r *rule) predicate(p part) bool {
	if r.tautology {
		return true
	}
	val := p[r.category]
	if r.lessThan {
		return val < r.comparand
	}
	return val > r.comparand
}

func (r *rule) splitRatingsInterval(ri ratingsInterval) (filter, complement ratingsInterval) {
	filter = ri.copy()
	if r.tautology {
		return
	}
	complement = ri.copy()
	var i, iComplement *interval.Interval
	if r.lessThan {
		i, _ = interval.NewInterval(startDefault, r.comparand-1)
		iComplement, _ = interval.NewInterval(r.comparand, endDefault)
	} else {
		i, _ = interval.NewInterval(r.comparand+1, endDefault)
		iComplement, _ = interval.NewInterval(startDefault, r.comparand)
	}
	cat := r.category
	filter[cat] = filter[cat].Intersection(i)
	complement[cat] = complement[cat].Intersection(iComplement)
	return
}
