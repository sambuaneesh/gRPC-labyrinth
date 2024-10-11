package game

import (
	"bufio"
	"fmt"
	"os"
)

type TileType int

const (
	Empty TileType = iota
	Wall
	Coin
)

const (
	SUCCESS MoveStatus = iota
	FAILURE
	VICTORY
	DEATH
)

type MoveStatus int

type Tile struct {
	Type TileType
}

type Labyrinth struct {
	Width  int
	Height int
	Tiles  [][]Tile
}

func (l *Labyrinth) IsValidPosition(position Position) bool {
	return position.X >= 0 && position.X < l.Width && position.Y >= 0 && position.Y < l.Height
}

func (l *Labyrinth) GetTile(position Position) Tile {
	if !l.IsValidPosition(position) {
		return Tile{Type: Wall}
	}
	return l.Tiles[position.Y][position.X]
}

func (l *Labyrinth) SetTile(position Position, tile Tile) {
	if !l.IsValidPosition(position) {
		return
	}
	l.Tiles[position.Y][position.X] = tile
}

func NewLabyrinth(filename string) (*Labyrinth, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	labyrinth := &Labyrinth{}

	// parsing labyrinth
	scanner := bufio.NewScanner(file)
	var tiles [][]Tile

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue // just in case
		}

		row := make([]Tile, len(line))
		for i, char := range line {
			switch char {
			case '.':
				row[i] = Tile{Type: Empty}
			case '#':
				row[i] = Tile{Type: Wall}
			case 'C':
				row[i] = Tile{Type: Coin}
			default:
				return nil, fmt.Errorf("invalid character in labyrinth file: %c", char)
			}
		}
		tiles = append(tiles, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(tiles) == 0 {
		return nil, fmt.Errorf("labyrinth file is empty")
	}

	labyrinth.Height = len(tiles)
	labyrinth.Width = len(tiles[0])
	labyrinth.Tiles = tiles

	return labyrinth, nil
}
