package game_test

import (
	"testing"

	"github.com/thotluna/ttt/internal/game"
	"github.com/thotluna/ttt/testutils"
)

func TestPlayer_MoveCount(t *testing.T) {
	mockIO := testutils.NewMockIO(
		"0.0",
		"0.1",
		"0.2",
	)

	board := game.NewBoard(mockIO)
	player := game.NewPlayer(game.PlayerX, mockIO, board)

	if player.MoveCount() != 0 {
		t.Error("El contador debe empezar en 0")
	}

	player.Play()
	if player.MoveCount() != 1 {
		t.Error("El contador debe ser 1 después del primer movimiento")
	}

	player.Play()
	if player.MoveCount() != 2 {
		t.Error("El contador debe ser 2 después del segundo movimiento")
	}

	player.Play()
	if player.MoveCount() != 3 {
		t.Error("El contador debe ser 3 después del tercer movimiento")
	}
}

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
					WithRow(0, game.EmptyCell, game.EmptyCell, game.EmptyCell).
					WithRow(1, game.EmptyCell, game.EmptyCell, game.EmptyCell).
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
					WithRow(0, game.PlayerX, game.EmptyCell, game.EmptyCell).
					WithRow(1, game.EmptyCell, game.EmptyCell, game.EmptyCell).
					WithRow(2, game.PlayerX, game.EmptyCell, game.EmptyCell).
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
					WithRow(0, game.PlayerX, game.EmptyCell, game.EmptyCell).
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
					WithRow(0, game.EmptyCell, game.EmptyCell, game.PlayerX).
					WithRow(1, game.EmptyCell, game.PlayerX, game.EmptyCell).
					WithRow(2, game.EmptyCell, game.EmptyCell, game.EmptyCell).
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

			player := game.NewPlayer(tt.symbol, mockIO, board)
			isWin := player.Play()

			if isWin != tt.expected {
				t.Errorf("%s: expected win=%v, got %v", tt.name, tt.expected, isWin)
			}
		})
	}
}
