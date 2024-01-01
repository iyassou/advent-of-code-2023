package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var categories = []string{"x", "m", "a", "s"}
var partRegex *regexp.Regexp

func init() {
	fields := []string{}
	for _, cat := range categories {
		fields = append(fields, fmt.Sprintf("%[1]s=(?P<%[1]s>\\d+)", cat))
	}
	together := strings.Join(fields, ",")
	partRegex = regexp.MustCompile(fmt.Sprintf("{%s}", together))
}

type part map[string]int

func NewPart(line string) (part, error) {
	match := partRegex.FindAllStringSubmatch(line, -1)
	if len(match) != 1 || len(match[0]) != 5 {
		return nil, errors.New("invalid part line")
	}
	p := part{}
	for _, cat := range categories {
		i := partRegex.SubexpIndex(cat)
		if rating, err := strconv.Atoi(match[0][i]); err != nil {
			return nil, fmt.Errorf("failed to parse %v", match[0][i])
		} else {
			p[cat] = rating
		}
	}
	return p, nil
}

func (p part) RatingsSum() int {
	sum := 0
	for _, cat := range categories {
		sum += p[cat]
	}
	return sum
}
