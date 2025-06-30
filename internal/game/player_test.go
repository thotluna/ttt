package game_test

import (
	"testing"

	"github.com/thotluna/ttt/internal/game"
	"github.com/thotluna/ttt/testutils"
)

func TestPlayer_CheckWin(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		board    func(board *game.Board) *game.Board
		symbol   game.Symbol
		expected bool
	}{
		{
			name:  "Horizontal win",
			input: "2.1",
			board: func(board *game.Board) *game.Board {
				return newBoardBuilder(board).
					WithRow(0, game.PlayerX, game.PlayerO, game.PlayerO).
					WithRow(1, game.PlayerO, game.PlayerO, game.PlayerX).
					WithRow(2, game.PlayerX, game.EmptyCell, game.PlayerX).
					Build()
			},

			symbol:   game.PlayerX,
			expected: true,
		},
		{
			name:  "Vertical win",
			input: "1.0",
			board: func(board *game.Board) *game.Board {
				return newBoardBuilder(board).
					WithRow(0, game.PlayerX, game.PlayerO, game.PlayerX).
					WithRow(1, game.EmptyCell, game.PlayerO, game.PlayerO).
					WithRow(2, game.PlayerX, game.PlayerX, game.PlayerO).
					Build()
			},
			symbol:   game.PlayerX,
			expected: true,
		},
		{
			name:  "Diagonal win",
			input: "2.2",
			board: func(board *game.Board) *game.Board {
				return newBoardBuilder(board).
					WithRow(0, game.PlayerX, game.PlayerO, game.PlayerO).
					WithRow(1, game.EmptyCell, game.PlayerX, game.EmptyCell).
					WithRow(2, game.EmptyCell, game.EmptyCell, game.EmptyCell).
					Build()
			},
			symbol:   game.PlayerX,
			expected: true,
		},
		{
			name:  "Inverter diagonal win",
			input: "2.0",
			board: func(board *game.Board) *game.Board {
				return newBoardBuilder(board).
					WithRow(0, game.PlayerX, game.PlayerO, game.PlayerX).
					WithRow(1, game.PlayerO, game.PlayerX, game.PlayerO).
					WithRow(2, game.EmptyCell, game.PlayerO, game.PlayerO).
					Build()
			},
			symbol:   game.PlayerX,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockIO := testutils.NewMockIO(tt.input)
			board := game.NewBoard(mockIO)
			board = tt.board(board)
			// board.SetBoard(tt.board(board))

			player := game.NewPlayer(tt.symbol, mockIO, board)
			isWin := player.Play()

			if isWin != tt.expected {
				t.Errorf("%s: expected win=%v, got %v", tt.name, tt.expected, isWin)
			}
		})
	}
}
