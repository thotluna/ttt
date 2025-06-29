package game

import (
	"github.com/thotluna/ttt/internal/view"
)

const (
	NumberRows = 3
	NumberCols = 3
)

type Board struct {
	board [NumberRows][NumberCols]rune
	io    view.IO
}

func NewBoard(io view.IO) *Board {
	return &Board{
		io: io,
		board: [NumberRows][NumberCols]rune{
			{'-', '-', '-'},
			{'-', '-', '-'},
			{'-', '-', '-'},
		},
	}
}

func (b *Board) GetBoard() [NumberRows][NumberCols]rune {
	return b.board
}

func (b *Board) SetBoard(board [NumberRows][NumberCols]rune) {
	b.board = board
}

func (b *Board) PlaceToken(symbol rune, coor Coordinate) error {
	if coor.row < 0 || coor.row >= NumberRows || coor.col < 0 || coor.col >= NumberCols {
		return NewGameError(ErrOutOfBounds, FormatPositionOutOfBounds(coor.row, coor.col))
	}

	if b.board[coor.row][coor.col] != '-' {
		return NewGameError(ErrPositionOccupied,
			FormatPositionTaken(coor.row, coor.col))
	}

	b.board[coor.row][coor.col] = symbol
	return nil
}

func (b *Board) GetTokenBy(symbol rune) []Coordinate {
	var coordinates []Coordinate
	for i := 0; i < NumberRows; i++ {
		for j := 0; j < NumberCols; j++ {
			if b.board[i][j] == symbol {
				coordinates = append(coordinates, Coordinate{i, j})
			}
		}
	}
	return coordinates
}

func (b *Board) FullBoard() bool {
	for i := 0; i < NumberRows; i++ {
		for j := 0; j < NumberCols; j++ {
			if b.board[i][j] == '-' {
				return false
			}
		}
	}
	return true
}

func (b *Board) Print() {
	b.io.PrintBoard(b.board)
}
