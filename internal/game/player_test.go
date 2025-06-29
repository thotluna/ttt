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
		board    [3][3]rune
		symbol   rune
		expected bool
	}{
		{
			name:  "Horizontal win",
			input: "2.1",
			board: [3][3]rune{
				{'X', 'O', 'O'},
				{'O', 'O', 'X'},
				{'X', '-', 'X'},
			},
			symbol:   'X',
			expected: true,
		},
		{
			name:  "Vertical win",
			input: "1.0",
			board: [3][3]rune{
				{'X', 'O', 'X'},
				{'-', 'O', 'O'},
				{'X', 'X', 'O'},
			},
			symbol:   'X',
			expected: true,
		},
		{
			name:  "Diagonal win",
			input: "2.2",
			board: [3][3]rune{
				{'X', 'O', 'O'},
				{'-', 'X', '-'},
				{'-', '-', '-'},
			},
			symbol:   'X',
			expected: true,
		},
		{
			name:  "Inverter diagonal win",
			input: "2.0",
			board: [3][3]rune{
				{'X', 'O', 'X'},
				{'O', 'X', 'O'},
				{'-', 'O', 'O'},
			},
			symbol:   'X',
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
