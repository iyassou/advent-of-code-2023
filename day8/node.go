package main

import (
	"fmt"
	"regexp"
)

var node = `([A-Z]{3})`
var nodeRegex = regexp.MustCompile(fmt.Sprintf(`%[1]s = \(%[1]s, %[1]s\)`, node))

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func (n *Node) String() string {
	l, r := " ", " "
	if n.Left != nil {
		l = n.Left.Value
	}
	if n.Right != nil {
		r = n.Right.Value
	}
	return fmt.Sprintf("Node[%s](%s|%s)", n.Value, l, r)
}
