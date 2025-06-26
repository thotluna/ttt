package game

import (
	"fmt"
)

type Board struct {
	board [3][3]rune
}

func NewBoard() *Board {
	return &Board{
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

func (b *Board) PlaceToken(token Token) error {
	if token.row < 0 || token.row >= 3 || token.col < 0 || token.col >= 3 {
		return NewGameError(ErrOutOfBounds,
			fmt.Sprintf("position (%d,%d) is out of bounds", token.row, token.col))
	}

	if b.board[token.row][token.col] != '-' {
		return NewGameError(ErrPositionOccupied,
			FormatPositionTaken(token.row, token.col))
			FormatPositionTaken(token.row, token.col))
	}

	b.board[token.row][token.col] = token.GetSymbol()
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
	// Check rows
	for i := 0; i < 3; i++ {
		if b.board[i][0] == b.board[i][1] && b.board[i][1] == b.board[i][2] && b.board[i][0] == turn {
			return true
		}
	}

	// Check columns
	for i := 0; i < 3; i++ {
		if b.board[0][i] == b.board[1][i] && b.board[1][i] == b.board[2][i] && b.board[0][i] == turn {
			return true
		}
	}

	// Check diagonals
	if b.board[0][0] == b.board[1][1] && b.board[1][1] == b.board[2][2] && b.board[0][0] == turn {
		return true
	}
	if b.board[0][2] == b.board[1][1] && b.board[1][1] == b.board[2][0] && b.board[0][2] == turn {
		return true
	}

	return false
}
func (b *Board) PrintBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%c ", b.board[i][j])
		}
		fmt.Println()
	}
}
