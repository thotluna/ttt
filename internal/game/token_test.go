package game_test

import (
	"testing"

	"github.com/thotluna/ttt/internal/game"
)

func TestNewToken(t *testing.T) {
	tests := []struct {
		symbol rune
		row    int
		col    int
	}{
		{'X', 0, 0},
		{'O', 1, 2},
		{'X', 2, 1},
	}

	for _, tc := range tests {
		token := game.NewToken(tc.symbol, tc.row, tc.col)

		if token.GetSymbol() != tc.symbol {
			t.Errorf("Expected symbol %c, got %c", tc.symbol, token.GetSymbol())
		}
		if token.GetRow() != tc.row {
			t.Errorf("Expected row %d, got %d", tc.row, token.GetRow())
		}
		if token.GetCol() != tc.col {
			t.Errorf("Expected col %d, got %d", tc.col, token.GetCol())
		}
	}
}

func TestGetSymbol(t *testing.T) {
	tests := []struct {
		symbol   rune
		expected rune
	}{
		{'X', 'X'},
		{'O', 'O'},
		{'A', 'A'}, // Caso con un símbolo no estándar
	}

	for _, tc := range tests {
		token := game.NewToken(tc.symbol, 0, 0)
		result := token.GetSymbol()
		if result != tc.expected {
			t.Errorf("Expected %c, got %c", tc.expected, result)
		}
	}
}

func TestTokenFields(t *testing.T) {
	token := game.NewToken('X', 1, 2)

	// Verificar que los campos no exportados sean accesibles dentro del paquete
	if token.GetSymbol() != 'X' {
		t.Error("Token symbol not set correctly")
	}
	if token.GetRow() != 1 {
		t.Error("Row not set correctly")
	}
	if token.GetCol() != 2 {
		t.Error("Column not set correctly")
	}
}
