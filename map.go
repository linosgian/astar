package main

import "fmt"

const (
	Tile     = 0
	Obstacle = 1
	Start    = 2
	Goal     = 3
)

type Map [][]uint8

type Node struct {
	X, Y    int
	Type    uint8
	f, g, h float64
	parent  *Node
}

func (m Map) GetNeighbours(x, y int) []uint8 {
	fmt.Println(m[x][y-1])
	nbs := []uint8{
		m[x+1][y],
		m[x-1][y],
		m[x][y+1],
		m[x][y-1],
	}
	return nbs
}
