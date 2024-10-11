package main

import (
	"fmt"
	"labyrinth/internal/game"
	"math/rand"
	"os"
	"strconv"
)

func GenerateLabyrinth(width int, height int) [][]game.Tile {
	labyrinth := make([][]game.Tile, height)
	for i := range labyrinth {
		labyrinth[i] = make([]game.Tile, width)
	}

	// Generate random tiles
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := rand.Float32()
			if r < 0.2 {
				labyrinth[y][x] = game.Tile{Type: game.Wall}
			} else if r < 0.4 {
				labyrinth[y][x] = game.Tile{Type: game.Coin}
			} else {
				labyrinth[y][x] = game.Tile{Type: game.Empty}
			}
		}
	}

	// Ensure start and end are empty
	labyrinth[0][0] = game.Tile{Type: game.Empty}
	labyrinth[height-1][width-1] = game.Tile{Type: game.Empty}

	// Export to file
	ExportLabyrinth(labyrinth, "../../server/labyrinth.txt")

	return labyrinth
}

func ExportLabyrinth(labyrinth [][]game.Tile, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, row := range labyrinth {
		for _, tile := range row {
			switch tile.Type {
			case game.Empty:
				file.WriteString(".")
			case game.Wall:
				file.WriteString("#")
			case game.Coin:
				file.WriteString("C")
			}
		}
		file.WriteString("\n")
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run generate_labyrinth.go <width> <height>")
		os.Exit(1)
	}

	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid width:", os.Args[1])
		os.Exit(1)
	}

	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid height:", os.Args[2])
		os.Exit(1)
	}

	labyrinth := GenerateLabyrinth(width, height)
	fmt.Println("Labyrinth generated and exported to labyrinth.txt")

	for _, row := range labyrinth {
		for _, tile := range row {
			switch tile.Type {
			case game.Empty:
				fmt.Print(".")
			case game.Wall:
				fmt.Print("#")
			case game.Coin:
				fmt.Print("C")
			}
		}
		fmt.Println()
	}
}
