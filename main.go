package main

import (
	"container/heap"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"

	"golang.org/x/image/bmp"
)

var (
	Black = color.RGBA{0, 0, 0, 255}       // Obstacle
	White = color.RGBA{255, 255, 255, 255} // Traversable tile
	Blue  = color.RGBA{0, 0, 255, 255}     // Starting point
	Red   = color.RGBA{255, 0, 0, 255}     // Goal point
)

func main() {
	const MapFilePath = "./assets/Map50_3.bmp"

	m, start, goal, err := initializeMap(MapFilePath)
	if err != nil {
		log.Fatalf("map was not initalized: %q", err)
	}
	fmt.Printf("Starting at (%d,%d) towards (%d,%d)\n", start.X, start.Y, goal.X, goal.Y)
	Astar(m, start, goal)
	m.drawMap(os.Stdout)
}

// Receives a grid map, a start and a goal Node.
// Finds the shortest path from start to goal.
func Astar(m Map, start, goal *Node) {
	openSet := &PriorityQueue{}
	heap.Init(openSet)

	heap.Push(openSet, start)

	for openSet.Len() > 0 {
		// Pop current node with highest priority
		c := heap.Pop(openSet).(*Node)

		if c == goal {
			break
		}
		c.open = false
		c.closed = true

		nbs := m.GetNeighbours(c)
		for _, nb := range nbs {
			// This is a potential new path to the neighbouring node
			nbTempCost := c.g + 1 // NOTE: this is the distance between current - neighbour

			// if this path has a lower cost, remove the old one
			// and follow the new one
			if nbTempCost < nb.g {
				if nb.open {
					heap.Remove(openSet, nb.index)
				}
				nb.open = false
				nb.closed = false
			}
			if !nb.open && !nb.closed {
				nb.g = nbTempCost
				nb.f = nb.g + ManhattanDistance(nb, goal)
				nb.open = true
				nb.parent = c
				heap.Push(openSet, nb)
			}
		}
	}
	// Visualize the optimal path
	for _, n := range backTracePath(m, start, goal) {
		m[n.X][n.Y].Type = Path
	}
}

// Given the Map, a start and goal Nodes,
// it reconstructs the shortest path by following the parent pointers
func backTracePath(m Map, start, goal *Node) []*Node {
	path := make([]*Node, 0)

	// Skip the goal point
	c := goal.parent
	for c != start {
		path = append(path, c)
		c = c.parent
	}
	// Skip starting point
	return path
}

// Extracts the color coded tiles that build the Map representation
// as a 2D slice of *Node
func initializeMap(MapFilePath string) (m Map, start *Node, goal *Node, err error) {

	r, err := os.Open(MapFilePath)
	if err != nil {
		log.Fatalf("could not open map file: %q", err)
	}
	img, err := bmp.Decode(r)
	if err != nil {
		log.Fatalf("could not decode map file: %q", err)
	}
	bs := img.Bounds()

	// Initialize 2D slice
	m = NewMap(bs.Max.X, bs.Max.Y)

	// Fill the Map with
	for y := range m {
		for x := range m[y] {
			n := &Node{X: x, Y: y}

			switch img.At(x, y) {
			case White:
				n.Type = Normal
			case Blue:
				n.Type = Start
				start = n
			case Red:
				n.Type = Goal
				goal = n
			default:
				n.Type = Obstacle
			}

			m[x][y] = n
		}
	}

	if start == nil || goal == nil {
		return nil, nil, nil, fmt.Errorf("could not find start or goal")
	}
	return m, start, goal, nil
}

// Takes the current and goal nodes and estimates their distance
// by using the Manhattan distance
func ManhattanDistance(c, g *Node) int {
	return int(math.Abs(float64(c.X-g.X)) + math.Abs(float64(c.Y-g.Y)))
}
