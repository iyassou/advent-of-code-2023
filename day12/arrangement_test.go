package main

import "testing"

func TestArrangementBruteForce(t *testing.T) {
	inputs := [][]byte{
		[]byte("???.### 1,1,3"),
		[]byte(".??..??...?##. 1,1,3"),
		[]byte("?#?#?#?#?#?#?#? 1,3,1,6"),
		[]byte("????.#...#... 4,1,1"),
		[]byte("????.######..#####. 1,6,5"),
		[]byte("?###???????? 3,2,1"),
	}
	outputs := []int{
		1, 4, 1, 1, 4, 10,
	}
	if len(inputs) != len(outputs) {
		t.Fatal("bruh")
	}
	for i, inp := range inputs {
		a, err := NewArrangement(inp)
		if err != nil {
			t.Fatal(err)
		}
		expected := outputs[i]
		actual := a.BruteForce()
		if expected != actual {
			t.Fatalf("input %s, expected %d, got %d", string(inp), expected, actual)
		}
	}
}
