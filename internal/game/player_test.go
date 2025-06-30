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
		board    [3][3]game.SymbolPlayerCurrent
		symbol   game.SymbolPlayerCurrent
		expected bool
	}{
		{
			name:  "Horizontal win",
			input: "2.1",
			board: [3][3]game.SymbolPlayerCurrent{
				{game.PlayerX, game.PlayerO, game.PlayerO},
				{game.PlayerO, game.PlayerO, game.PlayerX},
				{game.PlayerX, game.EmptyCell, game.PlayerX},
			},
			symbol:   game.PlayerX,
			expected: true,
		},
		{
			name:  "Vertical win",
			input: "1.0",
			board: [3][3]game.SymbolPlayerCurrent{
				{game.PlayerX, game.PlayerO, game.PlayerX},
				{game.EmptyCell, game.PlayerO, game.PlayerO},
				{game.PlayerX, game.PlayerX, game.PlayerO},
			},
			symbol:   game.PlayerX,
			expected: true,
		},
		{
			name:  "Diagonal win",
			input: "2.2",
			board: [3][3]game.SymbolPlayerCurrent{
				{game.PlayerX, game.PlayerO, game.PlayerO},
				{game.EmptyCell, game.PlayerX, game.EmptyCell},
				{game.EmptyCell, game.EmptyCell, game.EmptyCell},
			},
			symbol:   game.PlayerX,
			expected: true,
		},
		{
			name:  "Inverter diagonal win",
			input: "2.0",
			board: [3][3]game.SymbolPlayerCurrent{
				{game.PlayerX, game.PlayerO, game.PlayerX},
				{game.PlayerO, game.PlayerX, game.PlayerO},
				{game.EmptyCell, game.PlayerO, game.PlayerO},
			},
			symbol:   game.PlayerX,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockIO := testutils.NewMockIO(tt.input)
			board := game.NewBoard(mockIO)
			board.SetBoard(tt.board)

			player := game.NewPlayer(tt.symbol, mockIO, board)
			isWin := player.Play()

			if isWin != tt.expected {
				t.Errorf("%s: expected win=%v, got %v", tt.name, tt.expected, isWin)
			}
		})
	}
}
