package game

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.board[i][j] != '-' {
				t.Errorf("Expected '-' at position [%d][%d], got %c", i, j, board.board[i][j])
			}
		}
	}
}

func TestPlaceToken(t *testing.T) {
	board := NewBoard()
	token := NewToken('X', 1, 1)
	
	board.PlaceToken(token)
	if board.board[1][1] != 'X' {
		t.Error("Token not placed correctly at [1][1]")
	}

	// Verificar que otras celdas sigan vacías
	if board.board[0][0] != '-' || board.board[2][2] != '-' {
		t.Error("Other cells should remain empty")
	}
}

func TestFullBoard(t *testing.T) {
	tests := []struct {
		name   string
		tokens []Token
		full   bool
	}{
		{
			name:   "Empty board",
			tokens: []Token{},
			full:   false,
		},
		{
			name: "Partially filled board",
			tokens: []Token{
				{'X', 0, 0},
				{'O', 1, 1},
			},
			full: false,
		},
		{
			name: "Full board",
			tokens: []Token{
				{'X', 0, 0}, {'O', 0, 1}, {'X', 0, 2},
				{'O', 1, 0}, {'X', 1, 1}, {'O', 1, 2},
				{'X', 2, 0}, {'O', 2, 1}, {'X', 2, 2},
			},
			full: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			board := NewBoard()
			for _, token := range tc.tokens {
				board.PlaceToken(token)
			}
			if got := board.FullBoard(); got != tc.full {
				t.Errorf("Expected FullBoard() = %v, got %v", tc.full, got)
			}
		})
	}
}

func TestCheckWin(t *testing.T) {
	tests := []struct {
		name   string
		tokens []Token
		symbol rune
		win    bool
	}{
		{
			name:   "No win",
			tokens: []Token{{'X', 0, 0}, {'O', 1, 1}},
			symbol: 'X',
			win:    false,
		},
		{
			name:   "Row win",
			tokens: []Token{{'X', 0, 0}, {'X', 0, 1}, {'X', 0, 2}},
			symbol: 'X',
			win:    true,
		},
		{
			name:   "Column win",
			tokens: []Token{{'O', 0, 1}, {'O', 1, 1}, {'O', 2, 1}},
			symbol: 'O',
			win:    true,
		},
		{
			name:   "Diagonal 1 win",
			tokens: []Token{{'X', 0, 0}, {'X', 1, 1}, {'X', 2, 2}},
			symbol: 'X',
			win:    true,
		},
		{
			name:   "Diagonal 2 win",
			tokens: []Token{{'O', 0, 2}, {'O', 1, 1}, {'O', 2, 0}},
			symbol: 'O',
			win:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			board := NewBoard()
			for _, token := range tc.tokens {
				board.PlaceToken(token)
			}
			if got := board.CheckWin(tc.symbol); got != tc.win {
				t.Errorf("Expected CheckWin('%c') = %v, got %v", tc.symbol, tc.win, got)
			}
		})
	}
}
