package main

import (
	"log"
	"strconv"
	"strings"
)

type Direction int

const (
	R Direction = 0
	D Direction = 1
	L Direction = 2
	U Direction = 3
)

var string2direction = map[string]Direction{
	"R": R,
	"D": D,
	"L": L,
	"U": U,
}

type Instruction struct {
	direction Direction
	length    int
	colour    string
}

func NewInstruction(line string) *Instruction {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		log.Fatal(fields)
	}
	ins := &Instruction{colour: fields[2][2 : len(fields[2])-1]}
	if dir, ok := string2direction[fields[0]]; !ok {
		log.Fatal(fields[0])
	} else {
		ins.direction = dir
	}
	if length, err := strconv.Atoi(fields[1]); err != nil {
		log.Fatal(fields[1])
	} else {
		ins.length = length
	}
	return ins
}

func (i *Instruction) Apply(x, y int) (int, int) {
	switch i.direction {
	case L:
		x -= i.length
	case R:
		x += i.length
	case D:
		y -= i.length
	case U:
		y += i.length
	}
	return x, y
}

func (i *Instruction) DecodeColour() {
	if newLength, err := strconv.ParseInt(i.colour[:5], 16, 0); err != nil {
		log.Fatal(err)
	} else {
		i.length = int(newLength)
	}
	if dir, err := strconv.Atoi(i.colour[5:6]); err != nil {
		log.Fatal(err)
	} else {
		i.direction = Direction(dir)
	}
}
