package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/iyassou/advent-of-code-2023/internal"
)

//go:embed input.txt
var input string

type Race struct {
	Time     time.Duration
	Distance int64
}

func NewRace(input string) (*Race, error) {
	lines := internal.Lines(input)
	if l := len(lines); l != 2 {
		return nil, fmt.Errorf("expected times and distances, got %d inputs", l)
	}
	timestring := strings.Join(strings.Fields(strings.SplitAfter(lines[0], "Time:")[1]), "") + "ms"
	distancestring := strings.Join(strings.Fields(strings.SplitAfter(lines[1], "Distance:")[1]), "")
	if t, err := time.ParseDuration(timestring); err != nil {
		return nil, err
	} else if d, err := strconv.ParseInt(distancestring, 10, 64); err != nil {
		return nil, err
	} else {
		return &Race{Time: t, Distance: d}, nil
	}
}

func (r *Race) String() string {
	return fmt.Sprintf("Race[t: %v, d: %d]", r.Time, r.Distance)
}

func (r *Race) WaysOfBreakingTheRecord() int {
	recordBreaks := []time.Duration{}
	for i := 0 * time.Millisecond; i < r.Time; i += 1 * time.Millisecond {
		cost := i / time.Millisecond
		canCover := int64(r.Time/time.Millisecond-cost) * int64(cost)
		if canCover > r.Distance {
			recordBreaks = append(recordBreaks, i)
		}
	}
	return len(recordBreaks)
}

type Competition struct {
	Races []*Race
}

func NewCompetition(input string) (*Competition, error) {
	lines := internal.Lines(input)
	if l := len(lines); l != 2 {
		return nil, fmt.Errorf("expected times and distances, got %d inputs", l)
	}
	times := []time.Duration{}
	for _, timestring := range strings.Fields(strings.SplitAfter(lines[0], "Time:")[1]) {
		if t, err := time.ParseDuration(timestring + "ms"); err != nil {
			return nil, err
		} else {
			times = append(times, t)
		}
	}
	distances := []int64{}
	for _, distancestring := range strings.Fields(strings.SplitAfter(lines[1], "Distance:")[1]) {
		if d, err := strconv.ParseInt(distancestring, 10, 64); err != nil {
			return nil, err
		} else {
			distances = append(distances, d)
		}
	}
	if len(times) != len(distances) {
		return nil, fmt.Errorf("parsed %d times and %d distances", len(times), len(distances))
	}
	c := &Competition{Races: []*Race{}}
	for i, t := range times {
		d := distances[i]
		c.Races = append(c.Races, &Race{Time: t, Distance: d})
	}
	return c, nil
}

func main() {
	if c, err := NewCompetition(input); err != nil {
		log.Fatal(err)
	} else {
		prod := 1
		for _, r := range c.Races {
			prod *= r.WaysOfBreakingTheRecord()
		}
		fmt.Println("Part One:", prod)
	}
	if r, err := NewRace(input); err != nil {
		log.Fatal(err)
	} else {
		start := time.Now()
		fmt.Println("Part Two:", r.WaysOfBreakingTheRecord())
		end := time.Now()
		fmt.Printf("Took %v\n", end.Sub(start))
	}
}
