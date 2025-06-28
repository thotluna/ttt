package game

import "errors"

var (
	ErrInvalidPosition = errors.New("position must be between 0 and 2")
)

type Coordinate struct {
	row, col int
}

func NewCoordinate(row, col int) (Coordinate, error) {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return Coordinate{}, ErrInvalidPosition
	}
	return Coordinate{row: row, col: col}, nil
}

func (c Coordinate) Row() int {
	return c.row
}

func (c Coordinate) Col() int {
	return c.col
}

func (c Coordinate) IsValid() bool {
	return c.row >= 0 && c.row <= 2 && c.col >= 0 && c.col <= 2
}

func (c Coordinate) Equals(other Coordinate) bool {
	return c.row == other.row && c.col == other.col
}
