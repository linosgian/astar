package main

import "fmt"

const (
	Tile     = 0
	Obstacle = 1
	Start    = 2
	Goal     = 3
)

type Map [][]*Node

type Node struct {
	X, Y    int
	Type    uint8
	f, g, h float64
	parent  *Node
}

func NewMap(xLen, yLen int) Map {
	m := make([][]*Node, xLen)
	for i := range m {
		m[i] = make([]*Node, yLen)
	}
	return m
}

func (m Map) GetNeighbours(x, y int) []Node {
	nbs := make([]Node, 0)
	if x-1 >= 0 && m[x-1][y].Type != Obstacle {
		nbs = append(nbs, *m[x-1][y])
	}
	if x+1 < len(m) && m[x+1][y].Type != Obstacle {
		nbs = append(nbs, *m[x+1][y])
	}
	if y-1 >= 0 && m[x][y-1].Type != Obstacle {
		nbs = append(nbs, *m[x][y-1])
	}
	if y+1 < len(m[0]) && m[x][y+1].Type != Obstacle {
		nbs = append(nbs, *m[x][y+1])
	}
	return nbs
}
func (m Map) drawMap() {
	for y := range m {
		for x := range m[y] {
			switch m[x][y].Type {
			case Tile:
				fmt.Print(".")
			case Start:
				fmt.Print("0")
			case Goal:
				fmt.Print("X")
			default:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
