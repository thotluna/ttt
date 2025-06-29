package game

import (
	"github.com/thotluna/ttt/internal/view"
)

type Board struct {
	board [3][3]rune
	io    view.IO
}

func NewBoard(io view.IO) *Board {
	return &Board{
		io: io,
		board: [3][3]rune{
			{'-', '-', '-'},
			{'-', '-', '-'},
			{'-', '-', '-'},
		},
	}
}

func (b *Board) GetBoard() [3][3]rune {
	return b.board
}

func (b *Board) SetBoard(board [3][3]rune) {
	b.board = board
}

func (b *Board) PlaceToken(symbol rune, coor Coordinate) error {
	if coor.row < 0 || coor.row >= 3 || coor.col < 0 || coor.col >= 3 {
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
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.board[i][j] == symbol {
				coordinates = append(coordinates, Coordinate{i, j})
			}
		}
	}
	return coordinates
}

func (b *Board) FullBoard() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
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
