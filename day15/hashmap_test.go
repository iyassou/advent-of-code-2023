package main

import (
	"testing"
)

func sampleInitSequence() string {
	return "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
}

func TestHash(t *testing.T) {
	input := "HASH"
	expected := byte(52)
	if h, err := hash([]byte(input)); err != nil {
		t.Fatal(err)
	} else if h != expected {
		t.Fatalf("expected hash(%s)=%d, got %d", input, expected, h)
	}
}

func TestHashInitSequence(t *testing.T) {
	input := sampleInitSequence()
	expected := 1320
	h, err := hashInitSequence([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if h != expected {
		t.Fatalf("expected hashInitSequence(%s)=%d, got %d", input, expected, h)
	}
}

func TestNewHashmap(t *testing.T) {
	input := sampleInitSequence()
	actual, err := NewHashmap([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	expected, err := NewHashmap(nil)
	if err != nil {
		t.Fatal(err)
	}
	expected.Boxes[0] = []*Lens{
		{Label: []byte("rn"), FocalLength: 1},
		{Label: []byte("cm"), FocalLength: 2},
	}
	expected.Boxes[3] = []*Lens{
		{Label: []byte("ot"), FocalLength: 7},
		{Label: []byte("ab"), FocalLength: 5},
		{Label: []byte("pc"), FocalLength: 6},
	}
	if !actual.Equal(expected) {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

func TestHashmapFocusingPower(t *testing.T) {
	expected := 145
	if h, err := NewHashmap([]byte(sampleInitSequence())); err != nil {
		t.Fatal(err)
	} else if actual, err := h.FocusingPower(); err != nil {
		t.Fatal(err)
	} else if actual != expected {
		t.Fatalf("expected %d, got %d", expected, actual)
	}
}
