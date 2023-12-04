package main

import (
	"strconv"
	"testing"
)

func digitsInWords() map[string]int {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	digits := map[string]int{}
	for i, word := range words {
		digits[word] = i + 1
	}
	return digits
}

// TestCalibrationValuePartOneEmpty calls CalibrationValuePartOne on an
// empty string, which should return an error.
func TestCalibrationValuePartOneEmpty(t *testing.T) {
	v, err := CalibrationValuePartOne("")
	if err == nil {
		t.Fatalf("expected error value, got nil and %q instead", v)
	}
}

// TestCalibrationValuePartOneNoDigits calls CalibrationValuePartOne on a
// digit-less input.
func TestCalibrationValuePartOneNoDigits(t *testing.T) {
	if c, err := CalibrationValuePartOne("abcdefg"); err == nil {
		t.Fatalf("expected error, succeeded with %d", c)
	}
}

// TestCalibrationValuePartOneNoDigits calls CalibrationValuePartOne on a
// digit-less input.
func TestCalibrationValuePartOneJustDigits(t *testing.T) {
	in := "1234567890"
	out := 10
	if c, err := CalibrationValuePartOne(in); err != nil {
		t.Fatal(err)
	} else if c != out {
		t.Fatalf("expected %d for input %s, got %d\n", out, in, c)
	}
}

// TestCalibrationValuePartOneSingleDigit calls CalibrationValuePartOne
// on a single-digit input.
func TestCalibrationValuePartOneSingleDigit(t *testing.T) {
	for i := 0; i < 10; i++ {
		in := strconv.Itoa(i)
		out := i * 11
		if c, err := CalibrationValuePartOne(in); err != nil {
			t.Fatal(err)
		} else if c != out {
			t.Fatalf("expected %d for input %s, got %d\n", out, in, c)
		}
	}
}

// TestCalibrationValuePartOneSampleInput tests CalibrationValuePartOne
// on the example sample input.
func TestCalibrationValuePartOneSampleInputs(t *testing.T) {
	sampleInput := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet", "8three2h2"}
	expectedOutput := []int{12, 38, 15, 77, 82}
	if len(sampleInput) != len(expectedOutput) {
		t.Fatal("bruh")
	}
	for i, input := range sampleInput {
		expected := expectedOutput[i]
		if output, err := CalibrationValuePartOne(input); err != nil {
			t.Fatalf("input %s errored with %v", input, err)
		} else if output != expected {
			t.Fatalf("expected %d, got %d", expected, output)
		}
	}
}

// TestCalibrationValuePartTwoEmpty calls CalibrationValuePartTwo on an
// empty string, which should return an error.
func TestCalibrationValuePartTwoEmpty(t *testing.T) {
	v, err := CalibrationValuePartTwo("")
	if err == nil {
		t.Fatalf("expected error value, got nil and %q instead", v)
	}
}

// TestCalibrationValuePartTwoDigitsInWords calls CalibrationValuePartTwo
// on the ten digits written out in words.
func TestCalibrationValuePartTwoDigitsInWords(t *testing.T) {
	for digit, val := range digitsInWords() {
		expected := val * 11
		if c, err := CalibrationValuePartTwo(digit); err != nil {
			t.Fatalf("failed for %s with %v\n", digit, err)
		} else if c != expected {
			t.Fatalf("expected %d for %s, got %d instead\n", expected, digit, c)
		}
	}
}

// TestCalibrationValuePartTwoDigitsNumericAndWords tests CalibrationValuePartTwo
// on strings consisting of combinations of numeric digits and their word
// representations.
func TestCalibrationValuePartTwoDigitsNumericAndWords(t *testing.T) {
	for digit, val := range digitsInWords() {
		digitString := strconv.Itoa(val)
		inputs := []string{
			digitString + digit,
			digit + digitString,
			digit + digit,
			digitString + digitString,
		}
		expected := val * 11
		for _, in := range inputs {
			if c, err := CalibrationValuePartTwo(in); err != nil {
				t.Fatalf("failed for %s with %v\n", in, err)
			} else if c != expected {
				t.Fatalf("expected %d for %s, got %d instead\n", expected, in, c)
			}
		}
	}
}

// TestOverlappingDigitWords tests CalibrationValuePartTwo on digits whose
// word representations can overlap.
func TestOverlappingDigitWords(t *testing.T) {
	tricky := map[string]int{
		"oneight":   18,
		"threeight": 38,
		"fiveight":  58,
		"nineight":  98,
	}
	for in, expected := range tricky {
		if c, err := CalibrationValuePartTwo(in); err != nil {
			t.Fatalf("%s failed with %v\n", in, err)
		} else {
			if c != expected {
				t.Fatalf("expected %d for %s, got %d instead\n", expected, in, c)
			}
		}
	}
}
