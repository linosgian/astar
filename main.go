package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"golang.org/x/image/bmp"
)

const (
	MapFilePath = "/home/lgian/Pictures/Map50_3.bmp"
)

var (
	Black = color.RGBA{0, 0, 0, 255}
	White = color.RGBA{255, 255, 255, 255}
	Blue  = color.RGBA{0, 0, 255, 255}
	Red   = color.RGBA{255, 0, 0, 255}
)

func main() {
	m := initializeMap()
	drawMap(m)
	fmt.Println("Starting path finding to goal tile...")
	fmt.Println(m.GetNeighbours(10, 10))
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
	Map := make([][]uint8, bs.Max.X)
	for i := range Map {
		Map[i] = make([]uint8, bs.Max.Y)
	}

	for y := range Map {
		for x := range Map[y] {
			switch img.At(x, y) {
			case White:
				Map[x][y] = Tile
			case Blue:
				Map[x][y] = Start
			case Red:
				Map[x][y] = Goal
			default:
				Map[x][y] = Obstacle
			}
		}
	}
	return Map
}
func drawMap(m Map) {
	for y := range m {
		for x := range m[y] {
			switch m[x][y] {
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
