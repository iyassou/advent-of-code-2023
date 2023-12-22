package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func sampleInput() []byte {
	return []byte(".|...\\....\r\n|.-.\\.....\r\n.....|-...\r\n........|.\r\n..........\r\n.........\\\r\n..../.\\\\..\r\n.-.-/..|..\r\n.|....-|.\\\r\n..//.|....")
}

func TestNewContraption(t *testing.T) {
	expected := &Contraption{
		Tiles: [][]*Tile{
			{{EmptySpace, nil}, {VerticalSplitter, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {Backslash, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}},
			{{VerticalSplitter, nil}, {EmptySpace, nil}, {HorizontalSplitter, nil}, {EmptySpace, nil}, {Backslash, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}},
			{{EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {VerticalSplitter, nil}, {HorizontalSplitter, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}},
			{{EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {VerticalSplitter, nil}, {EmptySpace, nil}},
			{{EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}},
			{{EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {Backslash, nil}},
			{{EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {ForwardSlash, nil}, {EmptySpace, nil}, {Backslash, nil}, {Backslash, nil}, {EmptySpace, nil}, {EmptySpace, nil}},
			{{EmptySpace, nil}, {HorizontalSplitter, nil}, {EmptySpace, nil}, {HorizontalSplitter, nil}, {ForwardSlash, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {VerticalSplitter, nil}, {EmptySpace, nil}, {EmptySpace, nil}},
			{{EmptySpace, nil}, {VerticalSplitter, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {HorizontalSplitter, nil}, {VerticalSplitter, nil}, {EmptySpace, nil}, {Backslash, nil}},
			{{EmptySpace, nil}, {EmptySpace, nil}, {ForwardSlash, nil}, {ForwardSlash, nil}, {EmptySpace, nil}, {VerticalSplitter, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}, {EmptySpace, nil}},
		},
	}
	if actual, err := NewContraption(sampleInput()); err != nil {
		t.Fatal(err)
	} else if !cmp.Equal(actual, expected) {
		t.Fatalf("expected %v\ngot %v", expected, actual)
	}
}

func TestContraptionShineBeam(t *testing.T) {
	expected := 46
	c, err := NewContraption(sampleInput())
	if err != nil {
		t.Fatal(err)
	}
	if err := c.ShineBeam(0, 0, Right); err != nil {
		t.Fatalf("failed with: %v", err)
	} else if actual := c.NumberOfEnergisedTiles(); actual != expected {
		t.Fatalf("expected %d, got %d", expected, actual)
	}
}
