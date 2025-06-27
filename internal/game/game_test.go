package game_test

import (
	"testing"

	"github.com/thotluna/ttt/internal/game"
	"github.com/thotluna/ttt/testutils"
)

func TestGame_Play_XWins(t *testing.T) {
	mockIO := testutils.NewMockIO(
		"0.0", // X
		"1.0", // O
		"0.1", // X
		"1.1", // O
		"0.2", // X wins
	)

	g := game.NewGame(mockIO)
	g.Play()

	if !mockIO.ContainsOutput("Player X wins!") {
		t.Error("Expected X to win the game")
	}
}

func TestGame_Play_Draw(t *testing.T) {
	mockIO := testutils.NewMockIO(
		"0.0", // X
		"0.1", // O
		"0.2", // X
		"1.1", // O
		"1.0", // X
		"2.0", // O
		"1.2", // X
		"2.2", // O
		"2.1", // X - draw
	)

	g := game.NewGame(mockIO)
	g.Play()

	if !mockIO.ContainsOutput("It's a draw!") {
		t.Error("Expected game to end in a draw")
	}
}

func TestGame_Play_InvalidMove(t *testing.T) {
	mockIO := testutils.NewMockIO(
		"3.3", // invalid
		"0.0", // X
		"0.0", // already taken
		"1.0", // O
		"0.1", // X
		"1.1", // O
		"0.2", // X wins
	)

	g := game.NewGame(mockIO)
	g.Play()

	if !mockIO.ContainsOutput("Player X wins!") {
		t.Error("Expected X to win the game after invalid moves")
	}
}
