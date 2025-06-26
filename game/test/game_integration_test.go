package game_test

import (
	"testing"

	"github.com/thotluna/ttt/game"
	"github.com/thotluna/ttt/testutils"
)

func TestGame_WinCondition(t *testing.T) {
	mock := testutils.NewMockIO("0.0", "1.0", "0.1", "1.1", "0.2")

	g := game.NewGame(mock)
	g.Play()

	if !mock.ContainsOutput("Player X wins!") {
		t.Error("Expected X to win the game")
	}
}

func TestGame_DrawCondition(t *testing.T) {
	mock := testutils.NewMockIO("0.0", "0.1", "0.2", "1.1", "1.0", "2.0", "1.2", "2.2", "2.1")

	g := game.NewGame(mock)
	g.Play()

	if !mock.ContainsOutput("It's a draw!") {
		t.Error("Expected the game to end in a draw")
	}
}

func TestGame_InvalidInputs(t *testing.T) {
	tests := []struct {
		name        string
		inputs      []string
		expectedErr string
	}{
		{
			name:        "invalid format",
			inputs:      []string{"abc", "0.0", "1.1", "0.1", "1.0", "0.2"}, // X gana
			expectedErr: "Error: invalid input: invalid format. Please use 'row.col' (e.g., '1.2')",
		},
		{
			name:        "out of bounds",
			inputs:      []string{"5.5", "0.0", "1.1", "0.1", "1.0", "0.2"}, // X gana
			expectedErr: "position is out of bounds (0-2,0-2)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := testutils.NewMockIO(tt.inputs...)

			g := game.NewGame(mock)
			g.Play()

			if !mock.ContainsOutput(tt.expectedErr) {
				t.Errorf("Expected error containing: %s", tt.expectedErr)
			}

			if !mock.ContainsOutput("Player X wins!") {
				t.Error("Expected game to end with X winning")
			}
		})
	}
}
