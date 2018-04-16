package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"golang.org/x/image/bmp"
)

const (
	MapFilePath = "./assets/Map50_3.bmp"
)

var (
	Black = color.RGBA{0, 0, 0, 255}
	White = color.RGBA{255, 255, 255, 255}
	Blue  = color.RGBA{0, 0, 255, 255}
	Red   = color.RGBA{255, 0, 0, 255}
)

func main() {
	m := initializeMap()
	m.drawMap()
	fmt.Println("Starting path finding to goal tile...")
	fmt.Println(m.GetNeighbours(12, 10))
}

func Astar(m Map) {

}
func initializeMap() Map {
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
	m := NewMap(bs.Max.X, bs.Max.Y)

	for y := range m {
		for x := range m[y] {
			n := &Node{
				X:      x,
				Y:      y,
				parent: nil,
			}
			switch img.At(x, y) {
			case White:
				n.Type = Tile
			case Blue:
				n.Type = Start
			case Red:
				n.Type = Goal
			default:
				n.Type = Obstacle
			}
			m[x][y] = n
		}
	}
	return m
}
