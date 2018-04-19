package main

import "fmt"

type Tile uint8

const (
	Normal Tile = iota
	Obstacle
	Start
	Goal
	Path
)

type Node struct {
	X, Y    int
	Type    Tile // Type of node, e.g. Obstacle
	f, g, h int
	parent  *Node
	closed  bool
	open    bool
	index   int
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

// Draw a Map with ASCII characters
func (m Map) drawMap() {
	for y := range m {
		for x := range m[y] {
			switch m[x][y].Type {
			case Normal:
				fmt.Print(".")
			case Start:
				fmt.Print("0")
			case Path:
				fmt.Print("1")
			case Goal:
				fmt.Print("X")
			default:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
