package main

import (
	"fmt"
	"io"
)

type Tile uint8

// Types of Tiles
const (
	Normal Tile = iota
	Obstacle
	Start
	Goal
	Path
)

type Node struct {
	X, Y    int
	Type    Tile  // Type of node, e.g. Obstacle
	f, g, h int   // Distance costs
	parent  *Node // Used to reconstruct the shortest path
	closed  bool  // Indicates if this node is in the "closed set"
	open    bool  // Same as above, for the open set
	index   int   // Used by the Priority Queue
}

// Represents the map on which we'll run A*
type Map [][]*Node

// Initializes and returns a Map
func NewMap(xLen, yLen int) Map {
	m := make([][]*Node, xLen)
	for i := range m {
		m[i] = make([]*Node, yLen)
	}
	return m
}

// Given a *Node, returns its neighbours
// Diagonal movement is not allowed
func (m Map) GetNeighbours(n *Node) []*Node {
	nbs := make([]*Node, 0)
	x, y := n.X, n.Y
	if x-1 >= 0 && m[x-1][y].Type != Obstacle {
		nbs = append(nbs, m[x-1][y])
	}
	if x+1 < len(m) && m[x+1][y].Type != Obstacle {
		nbs = append(nbs, m[x+1][y])
	}
	if y-1 >= 0 && m[x][y-1].Type != Obstacle {
		nbs = append(nbs, m[x][y-1])
	}
	if y+1 < len(m[0]) && m[x][y+1].Type != Obstacle {
		nbs = append(nbs, m[x][y+1])
	}
	return nbs
}

// Draws a Map with ASCII characters
// Writes everything to the given io.Writer
func (m Map) drawMap(d io.Writer) {
	for y := range m {
		for x := range m[y] {
			switch m[x][y].Type {
			case Normal:
				fmt.Fprint(d, ".")
			case Start:
				fmt.Fprint(d, "0")
			case Path:
				fmt.Fprint(d, "1")
			case Goal:
				fmt.Fprint(d, "X")
			default:
				fmt.Fprint(d, "#")
			}
		}
		fmt.Fprintln(d)
	}
}

func (n Node) String() string {
	return fmt.Sprintf("Node[X: %d, Y: %d, Type: %d]", n.X, n.Y, n.Type)
}
