package game

import (
	"testing"
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
		token := NewToken(tc.symbol, tc.row, tc.col)
		
		if token.token != tc.symbol {
			t.Errorf("Expected symbol %c, got %c", tc.symbol, token.token)
		}
		if token.row != tc.row {
			t.Errorf("Expected row %d, got %d", tc.row, token.row)
		}
		if token.col != tc.col {
			t.Errorf("Expected col %d, got %d", tc.col, token.col)
		}
	}
}

func TestGetSymbol(t *testing.T) {
	tests := []struct {
		symbol rune
		expected rune
	}{
		{'X', 'X'},
		{'O', 'O'},
		{'A', 'A'}, // Caso con un símbolo no estándar
	}

	for _, tc := range tests {
		token := Token{token: tc.symbol}
		result := token.GetSymbol()
		if result != tc.expected {
			t.Errorf("Expected %c, got %c", tc.expected, result)
		}
	}
}

func TestTokenFields(t *testing.T) {
	token := NewToken('X', 1, 2)
	
	// Verificar que los campos no exportados sean accesibles dentro del paquete
	if token.token != 'X' {
		t.Error("Token symbol not set correctly")
	}
	if token.row != 1 {
		t.Error("Row not set correctly")
	}
	if token.col != 2 {
		t.Error("Column not set correctly")
	}
}
