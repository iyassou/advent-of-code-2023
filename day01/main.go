package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type FirstLastRunes struct {
	FirstIndex int
	FirstRune  rune
	LastIndex  int
	LastRune   rune
}

func (f *FirstLastRunes) CalibrationValue() (int, error) {
	return strconv.Atoi(string(f.FirstRune) + string(f.LastRune))
}

func FindFirstLastRunes(input string) (*FirstLastRunes, error) {
	flr := FirstLastRunes{}
	runes := []rune(input)
	for i, r := range runes {
		if unicode.IsDigit(r) {
			flr.FirstIndex = i
			flr.FirstRune = r
			break
		}
	}
	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsDigit(runes[i]) {
			flr.LastIndex = i
			flr.LastRune = runes[i]
			break
		}
	}
	if flr.FirstRune == 0 && flr.LastRune == 0 {
		return nil, errors.New("digit-less input string")
	}
	return &flr, nil
}

func CalibrationValuePartOne(input string) (int, error) {
	if len(input) == 0 {
		return 0, errors.New("empty string")
	}
	if flr, err := FindFirstLastRunes(input); err != nil {
		return 0, err
	} else {
		return flr.CalibrationValue()
	}
}

func CalibrationValuePartTwo(input string) (int, error) {
	// Construct regex matches
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	nums := map[*regexp.Regexp]rune{}
	for i, w := range words {
		nums[regexp.MustCompile(w)] = rune(49 + i)
	}
	// Find first and last runes.
	flr, err := FindFirstLastRunes(input)
	if err != nil {
		flr = &FirstLastRunes{
			LastIndex:  -1,
			FirstIndex: math.MaxInt,
		}
	}
	// Check if any regex matches occur before or after the last rune.
	for r, v := range nums {
		if locs := r.FindAllStringIndex(input, -1); locs != nil {
			if locs[0][0] < flr.FirstIndex {
				flr.FirstIndex = locs[0][0]
				flr.FirstRune = v
			}
			if last := len(locs) - 1; locs[last][0] > flr.LastIndex {
				flr.LastIndex = locs[last][0]
				flr.LastRune = v
			}
		}
	}
	if flr.FirstRune == 0 && flr.LastRune != 0 {
		flr.FirstRune = flr.LastRune
	} else if flr.FirstRune != 0 && flr.LastRune == 0 {
		flr.LastRune = flr.FirstRune
	} else if flr.FirstRune == 0 && flr.LastRune == 0 {
		return 0, errors.New("digit-less (word or numeric) input")
	}
	return flr.CalibrationValue()
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: <program> <input.txt>")
	}
	if b, err := os.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		var lines []string
		for _, lb := range bytes.Split(b, []byte{'\r', '\n'}) {
			lines = append(lines, string(lb))
		}
		total1, total2 := 0, 0
		for _, line := range lines {
			if c, err := CalibrationValuePartOne(line); err != nil {
				log.Fatal(err)
			} else {
				total1 += c
			}
			if c, err := CalibrationValuePartTwo(line); err != nil {
				log.Fatal(err)
			} else {
				total2 += c
			}
		}
		fmt.Println("Part One:", total1)
		fmt.Println("Part Two:", total2)
	}
}
