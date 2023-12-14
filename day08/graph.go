package main

import (
	"fmt"
	"strings"

	"github.com/iyassou/advent-of-code-2023/internal"
)

type Graph struct {
	Nodes map[*Node]bool
}

func NewGraph(input string) (*Graph, error) {
	lines := internal.Lines(input)
	if len(lines) == 0 {
		return nil, fmt.Errorf("invalid input %q", input)
	}
	g := &Graph{Nodes: map[*Node]bool{}}
	for _, line := range internal.Lines(input) {
		m := nodeRegex.FindAllStringSubmatch(line, -1)
		if len(m) != 1 || len(m[0]) != 4 {
			return nil, fmt.Errorf("invalid node regex %q", line)
		}

		nval := m[0][1]
		node, contains := g.Get(nval)
		if !contains {
			node = &Node{Value: nval}
			g.Nodes[node] = true
		}

		lval := m[0][2]
		if lnode, contains := g.Get(lval); contains {
			node.Left = lnode
		} else {
			node.Left = &Node{Value: lval}
			g.Nodes[node.Left] = true
		}

		rval := m[0][3]
		if rnode, contains := g.Get(rval); contains {
			node.Right = rnode
		} else {
			node.Right = &Node{Value: rval}
			g.Nodes[node.Right] = true
		}
	}
	return g, nil
}

func (g *Graph) Get(val string) (*Node, bool) {
	for n := range g.Nodes {
		if n.Value == val {
			return n, true
		}
	}
	return nil, false
}

func (g *Graph) TravelFromTo(path string, from, to string) int {
	if _, contains := g.Get(to); !contains {
		return -1
	}
	current, contains := g.Get(from)
	if !contains {
		return -1
	}
	steps := 0
	for current.Value != to {
		direction := path[steps%len(path)]
		steps++
		if direction == 'L' {
			current = current.Left
		} else if direction == 'R' {
			current = current.Right
		} else {
			return -1
		}
	}
	return steps
}

func (g *Graph) GhostTravel(path string, fromEnd, toEnd string) int {
	currentNodes := []*Node{}
	destinationPossible := false
	for node := range g.Nodes {
		if strings.HasSuffix(node.Value, fromEnd) {
			currentNodes = append(currentNodes, node)
		}
		if strings.HasSuffix(node.Value, toEnd) {
			destinationPossible = true
		}
	}
	if len(currentNodes) == 0 {
		return -1
	}
	if !destinationPossible {
		return -1
	}
	for _, p := range path {
		if !(p == 'L' || p == 'R') {
			return -1
		}
	}
	cycles := make([][2]int, len(currentNodes))
	for i, current := range currentNodes {
		// Floyd's cycle detection algorithm
		moveAnimal := func() func(animal *Node, by int) *Node {
			a := 0
			return func(animal *Node, by int) *Node {
				for ; by > 0; by-- {
					a = (a + 1) % len(path)
					if path[a] == 'L' {
						animal = animal.Left
					} else {
						animal = animal.Right
					}
				}
				return animal
			}
		}
		moveTortoise, moveHare := moveAnimal(), moveAnimal()
		var tortoise, hare *Node
		for tortoise, hare = moveTortoise(current, 1), moveHare(current, 2); tortoise != hare; {
			tortoise = moveTortoise(tortoise, 1)
			hare = moveHare(hare, 2)
		}
		mu := 0
		tortoise = current
		for ; tortoise != hare; mu++ {
			tortoise = moveTortoise(tortoise, 1)
			hare = moveHare(hare, 1)
		}
		lambda := 1
		hare = moveHare(tortoise, 1)
		for tortoise != hare {
			hare = moveHare(hare, 1)
			lambda++
		}
		cycles[i] = [2]int{lambda, mu}
		fmt.Printf("Cycle\t%d: λ=%d, μ=%d\n", i, lambda, mu)
		// iter := current
		// for j := 0; j < mu; j++ {
		// 	fmt.Printf(buildSuffix(iter.Value))
		// 	iter = moveAnimal(iter, 1)
		// }
	}

	return -1
}
