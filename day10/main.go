package main

import (
	_ "embed"
	"log"
)

//go:embed input.txt
var input []byte

func main() {
	input = []byte("FF7FSF7F7F7F7F7F---7\r\nL|LJ||||||||||||F--J\r\nFL-7LJLJ||||||LJL-77\r\nF--JF--7||LJLJ7F7FJ-\r\nL---JF-JLJ.||-FJLJJ7\r\n|F|F-JF---7F7-L7L|7|\r\n|FFJF7L7F-JF7|JL---7\r\n7-L-JL7||F7|L7F-7F7|\r\nL.L7LFJ|||||FJL7||LJ\r\nL7JLJL-JLJLJL--JLJ.L")
	g, err := NewGrid(input)
	if err != nil {
		log.Fatal(err)
	}
	loop := g.GetMainLoop()
	// log.Println("Part One:", len(loop)/2)
	log.Printf("\n%s\n", g)
	log.Printf("\n%s\n", g.highlightPath(loop))
}
