package game_test

import (
	"testing"

	"github.com/thotluna/ttt/internal/game"
	"github.com/thotluna/ttt/testutils"
)

func TestNewTurn(t *testing.T) {
	mock, players := createMapPlayers()

	turn := game.NewTurn(players, mock)
	player := turn.GetTurn()
	if player != game.PlayerX {
		t.Errorf("Expected 'X', got %c", player)
	}
}

func TestTurnChange_SingleChange(t *testing.T) {
	mock, players := createMapPlayers()

	turn := game.NewTurn(players, mock)
	turn.TurnChange()

	player := turn.GetTurn()
	if player != game.PlayerO {
		t.Errorf("After one change, expected 'O', got %c", player)
	}
}

func TestTurnChange_MultipleChanges(t *testing.T) {
	mock, players := createMapPlayers()

	turn := game.NewTurn(players, mock)

	turn.TurnChange()
	player := turn.GetTurn()
	if player != game.PlayerO {
		t.Errorf("After first change, expected 'O', got %c", player)
	}

	turn.TurnChange()
	player = turn.GetTurn()
	if player != game.PlayerX {
		t.Errorf("After second change, expected 'X', got %c", player)
	}
}

func TestTurnChange_MultiplePlayers(t *testing.T) {
	mock, players := createMapPlayers()

	turn := game.NewTurn(players, mock)

	testCases := []struct {
		expectedSymbol game.Symbol
	}{
		{game.PlayerX},
		{game.PlayerO},
		{game.PlayerX},
		{game.PlayerO},
	}

	for i, tc := range testCases {
		if i > 0 {
			turn.TurnChange()
		}
		if turn.GetTurn() != tc.expectedSymbol {
			t.Errorf("After %d changes, expected %c, got %c", i, tc.expectedSymbol, turn.GetTurn())
		}
	}
}

func createMapPlayers() (*testutils.MockIO, map[game.Symbol]game.Player) {
	mock := testutils.NewMockIO()
	players := map[game.Symbol]game.Player{
		game.PlayerX: *game.NewPlayer(game.PlayerX, mock, nil),
		game.PlayerO: *game.NewPlayer(game.PlayerO, mock, nil),
	}
	return mock, players
}
