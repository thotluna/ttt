package game

import (
	"fmt"
)

type Direction string

const (
	Vertical   Direction = "vertical"
	Horizontal Direction = "horizontal"
	Diagonal   Direction = "diagonal"
	Inverter   Direction = "inverter"
	None       Direction = "none"
)

type Coordinate struct {
	row, col int
}

func NewCoordinate(row, col int) (Coordinate, error) {
	interval := NewGameInterval()
	if !interval.Contains(row) || !interval.Contains(col) {
		return Coordinate{}, NewGameError(ErrOutOfBounds, fmt.Sprintf(MsgPositionOutOfBounds, row, col))
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
	interval := NewGameInterval()
	return interval.Contains(c.row) && interval.Contains(c.col)
}

func (c Coordinate) Equals(other Coordinate) bool {
	return c.row == other.row && c.col == other.col
}

func (c Coordinate) IsVerticalTo(other Coordinate) bool {
	return c.col == other.col && c.row != other.row
}

func (c Coordinate) IsHorizontalTo(other Coordinate) bool {
	return c.row == other.row && c.col != other.col
}

func (c Coordinate) IsDiagonalTo(other Coordinate) bool {
	return c.row == c.col && other.row == other.col && c.row != other.row
}

func (c Coordinate) IsInverterTo(other Coordinate) bool {
	return (c.row+c.col == other.row+other.col) && (c.row != other.row) && (c.col != other.col)
}

func (c Coordinate) Direction(other Coordinate) Direction {
	if c.IsVerticalTo(other) {
		return Vertical
	}
	if c.IsHorizontalTo(other) {
		return Horizontal
	}
	if c.IsDiagonalTo(other) {
		return Diagonal
	}
	if c.IsInverterTo(other) {
		return Inverter
	}
	return None
}
