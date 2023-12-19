package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func sampleInputs() [][]byte {
	return [][]byte{
		[]byte("#.##..##.\r\n..#.##.#.\r\n##......#\r\n##......#\r\n..#.##.#.\r\n..##..##.\r\n#.#.##.#."),
		[]byte("#...##..#\r\n#....#..#\r\n..##..###\r\n#####.##.\r\n#####.##.\r\n..##..###\r\n#....#..#"),
	}
}

func makePatterns(inputs [][]byte) []*Pattern {
	ps := make([]*Pattern, len(inputs))
	for i, inp := range inputs {
		ps[i] = NewPattern(inp)
	}
	return ps
}

func TestNewPattern(t *testing.T) {
	inputs := makePatterns(sampleInputs())
	outputs := []*Pattern{
		{
			Rows:    []int{205, 180, 259, 259, 180, 204, 181},
			Columns: []int{77, 12, 115, 33, 82, 82, 33, 115, 12},
		},
		{
			Rows:    []int{305, 289, 460, 223, 223, 460, 289},
			Columns: []int{91, 24, 60, 60, 25, 67, 60, 60, 103},
		},
	}
	if len(inputs) != len(outputs) {
		t.Fatal("bruh")
	}
	for i, actual := range inputs {
		expected := outputs[i]
		if !cmp.Equal(expected, actual) {
			t.Errorf("[%d] expected %v differs from actual %v", i, expected, actual)
		}
	}
}

func TestPatternFindReflection(t *testing.T) {
	inputs := makePatterns(sampleInputs())
	outputs := []int{
		-5, 4,
	}
	if len(inputs) != len(outputs) {
		t.Fatal("bruh")
	}
	for i, inp := range inputs {
		expected := outputs[i]
		actual := inp.FindReflection()
		if expected != actual {
			t.Errorf("[%d] expected %d, got %d", i, expected, actual)
		}
	}
}

func TestPatternWipeSmudgeTwice(t *testing.T) {
	inputs := sampleInputs()
	smudges := [][2]int{
		{0, 0}, {0, 4},
	}
	if len(inputs) != len(smudges) {
		t.Fatal("bruh")
	}
	for i, inp := range inputs {
		x, y := smudges[i][0], smudges[i][1]
		original := NewPattern(inp)
		twiceWiped := NewPattern(inp)
		twiceWiped.WipeSmudge(x, y)
		twiceWiped.WipeSmudge(x, y)
		if !cmp.Equal(original, twiceWiped) {
			t.Errorf("[%d] expected two wipes to cancel out: original=%v, twiceWiped=%v\n", i, original, twiceWiped)
		}
	}
}

func TestPatternFindSmudgedReflection(t *testing.T) {
	inputs := makePatterns(sampleInputs())
	outputs := []int{
		3, 1,
	}
	if len(inputs) != len(outputs) {
		t.Fatal("bruh")
	}
	for i, inp := range inputs {
		expected := outputs[i]
		actual := inp.FindSmudgedReflection(0)
		if expected != actual {
			t.Errorf("[%d] expected %d, got %d", i, expected, actual)
		}
	}
}
