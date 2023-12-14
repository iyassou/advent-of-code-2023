package main

import "strings"

type Entry string

const (
	Seed        Entry = "seed"
	Soil        Entry = "soil"
	Fertilizer  Entry = "fertilizer"
	Water       Entry = "water"
	Light       Entry = "light"
	Temperature Entry = "temperature"
	Humidity    Entry = "humidity"
	Location    Entry = "location"
)

var entryRegex = strings.Join([]string{
	"(",
	string(Seed), string(Soil), string(Fertilizer), string(Water),
	string(Light), string(Temperature), string(Humidity), string(Location),
	")",
}, `|`)
