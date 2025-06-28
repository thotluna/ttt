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

func (c Coordinate) IsVerticalTo(other Coordinate) bool {
	return c.row == other.row && c.col != other.col
}

func (c Coordinate) IsHorizontalTo(other Coordinate) bool {
	return c.col == other.col && c.row != other.row
}

func (c Coordinate) IsDiagonalTo(other Coordinate) bool {
	return c.row == c.col && other.row == other.col
}

func (c Coordinate) IsInverterTo(other Coordinate) bool {
	return c.row+c.col == 3 && other.row+other.col == 3
}

func (c Coordinate) Direction(other Coordinate) string {
	if c.IsVerticalTo(other) {
		return "vertical"
	}
	if c.IsHorizontalTo(other) {
		return "horizontal"
	}
	if c.IsDiagonalTo(other) {
		return "diagonal"
	}
	if c.IsInverterTo(other) {
		return "inverter"
	}
	return "NOT"
}
