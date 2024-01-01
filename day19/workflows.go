package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/iyassou/advent-of-code-2023/internal"
)

type workflow struct {
	name  string
	rules []*rule
}

func newWorkflow(line string) (*workflow, error) {
	lbrace := strings.Index(line, "{")
	if lbrace == -1 {
		return nil, errors.New("bruh")
	}
	rs := strings.Split(line[lbrace+1:len(line)-1], ",")
	w := &workflow{name: line[:lbrace], rules: make([]*rule, len(rs))}
	for i, r := range rs {
		if rule, err := newRule(r); err != nil {
			return nil, err
		} else {
			w.rules[i] = rule
		}
	}
	return w, nil
}

func (w *workflow) evaluate(p part) string {
	for _, r := range w.rules {
		if !r.predicate(p) {
			continue
		}
		return r.iftrue
	}
	return ""
}

func (w *workflow) processRatingsInterval(ri ratingsInterval) map[string][]ratingsInterval {
	jobs := map[string][]ratingsInterval{}
	current := ri
	for _, r := range w.rules {
		filter, complement := r.splitRatingsInterval(current)
		jobs[r.iftrue] = append(jobs[r.iftrue], filter)
		current = complement
	}
	return jobs
}

type workflows map[string]*workflow

func NewWorkflows(input string) (workflows, error) {
	wfs := internal.Lines(input)
	w := make(workflows, len(wfs))
	for _, line := range wfs {
		if flow, err := newWorkflow(line); err != nil {
			return nil, err
		} else {
			w[flow.name] = flow
		}
	}
	return w, nil
}

func (w workflows) Accepts(p part, starting string) (bool, error) {
	name := starting
	for {
		wf := w[name]
		next := wf.evaluate(p)
		if next == "R" {
			return false, nil
		}
		if next == "A" {
			return true, nil
		}
		if next == "" {
			return false, fmt.Errorf("failed to decide part %v starting at %v", p, starting)
		}
		name = next
	}
}

func (w workflows) DistinctCombinations(starting string) (int, error) {
	accepts := []ratingsInterval{}
	firstBatch := w[starting].processRatingsInterval(newRatingsInterval())
	stack := []map[string][]ratingsInterval{firstBatch}
	for len(stack) != 0 {
		removeFirst := len(stack)
		for _, jobs := range stack {
			for name, ris := range jobs {
				wf := w[name]
				for _, ri := range ris {
					newJobs := wf.processRatingsInterval(ri)
					// check new jobs for accepts/rejects i.e. dead ends
					if a, ok := newJobs["A"]; ok {
						accepts = append(accepts, a...)
						delete(newJobs, "A")
					}
					delete(newJobs, "R")
					// add new jobs to stack
					if len(newJobs) > 0 {
						stack = append(stack, newJobs)
					}
					delete(jobs, name)
				}
			}
		}
		stack = stack[removeFirst:]
	}
	sum := 0
	for _, r := range accepts {
		sum += r.distinctCombinations()
	}
	return sum, nil
}
