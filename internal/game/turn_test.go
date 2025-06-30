package game_test

import (
	"testing"

	"github.com/thotluna/ttt/internal/game"
	"github.com/thotluna/ttt/testutils"
)

func TestNewTurn(t *testing.T) {
	mock, players := createMapPlayers()

	turn := game.NewTurn(players, mock)
	_, player := turn.GetTurn()
	if player != 'X' {
		t.Errorf("Expected 'X', got %c", player)
	}
}

func TestTurnChange_SingleChange(t *testing.T) {
	mock, players := createMapPlayers()

	turn := game.NewTurn(players, mock)
	turn.TurnChange()

	_, player := turn.GetTurn()
	if player != 'O' {
		t.Errorf("After one change, expected 'O', got %c", player)
	}
}

func TestTurnChange_MultipleChanges(t *testing.T) {
	mock, players := createMapPlayers()

	turn := game.NewTurn(players, mock)

	turn.TurnChange()
	_, player := turn.GetTurn()
	if player != 'O' {
		t.Errorf("After first change, expected 'O', got %c", player)
	}

	turn.TurnChange()
	_, player = turn.GetTurn()
	if player != 'X' {
		t.Errorf("After second change, expected 'X', got %c", player)
	}
}

func TestTurnChange_MultiplePlayers(t *testing.T) {
	mock, players := createMapPlayers()

	turn := game.NewTurn(players, mock)

	testCases := []struct {
		expectedNum  int
		expectedRune rune
	}{
		{0, 'X'}, // Initial
		{1, 'O'}, // After 1 change
		{0, 'X'}, // After 2 changes
		{1, 'O'}, // After 3 changes
	}

	for i, tc := range testCases {
		if i > 0 {
			turn.TurnChange()
		}
		_, symbol := turn.GetTurn()
		if symbol != tc.expectedRune {
			t.Errorf("After %d changes, expected %c, got %c", i, tc.expectedRune, symbol)
		}
	}
}

func createMapPlayers() (*testutils.MockIO, map[game.SymbolPlayerCurrent]game.Player) {
	mock := testutils.NewMockIO()
	players := map[game.SymbolPlayerCurrent]game.Player{
		game.PlayerX: *game.NewPlayer(game.PlayerX, mock, nil),
		game.PlayerO: *game.NewPlayer(game.PlayerO, mock, nil),
	}
	return mock, players
}
