package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func hash(data []byte) (byte, error) {
	if data == nil {
		return 0, fmt.Errorf("bruh")
	}
	hash := byte(0)
	for _, b := range data {
		if 0x20 <= b && b <= 0x7E {
			hash = (hash + b) * 17
		}
	}
	return hash, nil
}

func hashInitSequence(data []byte) (int, error) {
	sum := 0
	for _, step := range bytes.Split(data, []byte{','}) {
		if h, err := hash(step); err != nil {
			return 0, fmt.Errorf("bruh: %s", step)
		} else {
			sum += int(h)
		}
	}
	return sum, nil
}

type Lens struct {
	Label       []byte
	FocalLength int
}

func NewLens(label []byte, focalLength int) *Lens {
	return &Lens{Label: label, FocalLength: focalLength}
}

type Hashmap struct {
	Boxes [256][]*Lens
}

var errNilHashmap = fmt.Errorf("nil Hashmap")

func NewHashmap(data []byte) (*Hashmap, error) {
	h := &Hashmap{}
	if data == nil {
		return h, nil
	}
	for _, step := range bytes.Split(data, []byte{','}) {
		if parts := bytes.Split(step, []byte{'='}); len(parts) == 2 {
			if focalLength, err := strconv.Atoi(string(parts[1])); err != nil {
				return nil, err
			} else if err := h.Add(NewLens(parts[0], focalLength)); err != nil {
				return nil, err
			}
		} else if label := bytes.TrimSuffix(step, []byte{'-'}); len(label) > 0 {
			if err := h.Discard(label); err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("could not parse lens: %s", step)
		}
	}
	return h, nil
}

func (h *Hashmap) Equal(other *Hashmap) bool {
	if h == nil {
		return other == nil
	}
	if other == nil {
		return false
	}
	for i, box := range h.Boxes {
		otherBox := other.Boxes[i]
		if len(box) != len(otherBox) {
			return false
		}
		for j, lens := range box {
			otherLens := otherBox[j]
			if !bytes.Equal(lens.Label, otherLens.Label) {
				return false
			}
			if lens.FocalLength != otherLens.FocalLength {
				return false
			}
		}
	}
	return true
}

func (h *Hashmap) Add(lens *Lens) error {
	if h == nil {
		return errNilHashmap
	}
	if i, err := hash(lens.Label); err != nil {
		return err
	} else {
		box := &h.Boxes[i]
		for _, otherLens := range *box {
			if bytes.Equal(otherLens.Label, lens.Label) {
				otherLens.FocalLength = lens.FocalLength
				return nil
			}
		}
		*box = append(*box, lens)
	}
	return nil
}

func (h *Hashmap) Discard(label []byte) error {
	if h == nil {
		return errNilHashmap
	}
	if i, err := hash(label); err != nil {
		return err
	} else {
		box := &h.Boxes[i]
		for i, lens := range *box {
			if !bytes.Equal(label, lens.Label) {
				continue
			}
			if i < len(*box)-1 {
				*box = append((*box)[:i], (*box)[i+1:]...)
			} else {
				*box = (*box)[:len(*box)-1]
			}
			return nil
		}
	}
	return nil
}

func (h *Hashmap) FocusingPower() (int, error) {
	if h == nil {
		return 0, errNilHashmap
	}
	sum := 0
	for i, box := range h.Boxes {
		for slot, lens := range box {
			sum += (i + 1) * (slot + 1) * lens.FocalLength
		}
	}
	return sum, nil
}
