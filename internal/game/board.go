package game

import "github.com/thotluna/ttt/internal/view"

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
func (b *Board) CheckWin(turn rune) bool {
	for i := 0; i < 3; i++ {
		if b.board[i][0] == b.board[i][1] && b.board[i][1] == b.board[i][2] && b.board[i][0] == turn {
			return true
		}
	}

	for i := 0; i < 3; i++ {
		if b.board[0][i] == b.board[1][i] && b.board[1][i] == b.board[2][i] && b.board[0][i] == turn {
			return true
		}
	}

	if b.board[0][0] == b.board[1][1] && b.board[1][1] == b.board[2][2] && b.board[0][0] == turn {
		return true
	}
	if b.board[0][2] == b.board[1][1] && b.board[1][1] == b.board[2][0] && b.board[0][2] == turn {
		return true
	}

	return false
}

func (b *Board) Print() {
	b.io.PrintBoard(b.board)
}
