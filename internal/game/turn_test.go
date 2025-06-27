package game_test

import (
	"testing"

	"github.com/thotluna/ttt/internal/game"
	"github.com/thotluna/ttt/testutils"
)

func TestNewTurn(t *testing.T) {
	mock := testutils.NewMockIO()
	players := []*game.Player{
		game.NewPlayer('X', mock, nil),
		game.NewPlayer('O', mock, nil),
	}

	turn := game.NewTurn(players, mock)
	indexPlayer, _ := turn.GetTurn()
	if indexPlayer != 0 {
		t.Errorf("Expected initial turn to be 0, got %d", indexPlayer)
	}
}

func TestGetTurn_InitialPlayer(t *testing.T) {
	mock := testutils.NewMockIO()
	players := []*game.Player{
		game.NewPlayer('X', mock, nil),
		game.NewPlayer('O', mock, nil),
	}

	turn := game.NewTurn(players, mock)
	indexPlayer, player := turn.GetTurn()

	if indexPlayer != 0 || player != 'X' {
		t.Errorf("Expected (0, 'X'), got (%d, %c)", indexPlayer, player)
	}
}

func TestTurnChange_SingleChange(t *testing.T) {
	mock := testutils.NewMockIO()
	players := []*game.Player{
		game.NewPlayer('X', mock, nil),
		game.NewPlayer('O', mock, nil),
	}

	turn := game.NewTurn(players, mock)
	turn.TurnChange()

	indexPlayer, player := turn.GetTurn()
	if indexPlayer != 1 || player != 'O' {
		t.Errorf("After one change, expected (1, 'O'), got (%d, %c)", indexPlayer, player)
	}
}

func TestTurnChange_MultipleChanges(t *testing.T) {
	mock := testutils.NewMockIO()
	players := []*game.Player{
		game.NewPlayer('X', mock, nil),
		game.NewPlayer('O', mock, nil),
	}

	turn := game.NewTurn(players, mock)

	turn.TurnChange()
	indexPlayer, player := turn.GetTurn()
	if indexPlayer != 1 || player != 'O' {
		t.Errorf("After first change, expected (1, 'O'), got (%d, %c)", indexPlayer, player)
	}

	turn.TurnChange()
	indexPlayer, player = turn.GetTurn()
	if indexPlayer != 0 || player != 'X' {
		t.Errorf("After second change, expected (0, 'X'), got (%d, %c)", indexPlayer, player)
	}
}

func TestTurnChange_MultiplePlayers(t *testing.T) {
	mock := testutils.NewMockIO()
	players := []*game.Player{
		game.NewPlayer('X', mock, nil),
		game.NewPlayer('O', mock, nil),
	}

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
		turnNum, symbol := turn.GetTurn()
		if turnNum != tc.expectedNum || symbol != tc.expectedRune {
			t.Errorf("After %d changes, expected (%d, %c), got (%d, %c)",
				i, tc.expectedNum, tc.expectedRune, turnNum, symbol)
		}
	}
}
