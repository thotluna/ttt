package game

import (
	"testing"
)

func TestNewTurn(t *testing.T) {
	turn := NewTurn()
	if turn.turn != 0 {
		t.Errorf("Expected initial turn to be 0, got %d", turn.turn)
	}
}

func TestGetTurn_InitialPlayer(t *testing.T) {
	turn := NewTurn()
	turnNum, symbol := turn.GetTurn()
	
	if turnNum != 0 || symbol != 'X' {
		t.Errorf("Expected (0, 'X'), got (%d, %c)", turnNum, symbol)
	}
}

func TestTurnChange_SingleChange(t *testing.T) {
	turn := NewTurn()
	turn.TurnChange()
	
	turnNum, symbol := turn.GetTurn()
	if turnNum != 1 || symbol != 'O' {
		t.Errorf("After one change, expected (1, 'O'), got (%d, %c)", turnNum, symbol)
	}
}

func TestTurnChange_MultipleChanges(t *testing.T) {
	turn := NewTurn()
	
	// First change (X -> O)
	turn.TurnChange()
	turnNum, symbol := turn.GetTurn()
	if turnNum != 1 || symbol != 'O' {
		t.Errorf("After first change, expected (1, 'O'), got (%d, %c)", turnNum, symbol)
	}
	
	// Second change (O -> X)
	turn.TurnChange()
	turnNum, symbol = turn.GetTurn()
	if turnNum != 0 || symbol != 'X' {
		t.Errorf("After second change, expected (0, 'X'), got (%d, %c)", turnNum, symbol)
	}
}

func TestTurnChange_MultiplePlayers(t *testing.T) {
	turn := NewTurn()
	
	// Test multiple cycles of turns
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
