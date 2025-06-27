package game_test

import (
	"testing"

	"github.com/thotluna/ttt/internal/game"
)

func TestNewToken(t *testing.T) {
	tests := []struct {
		name   string
		symbol rune
		row    int
		col    int
		valid  bool
	}{
		{"Valid token X at 0,0", 'X', 0, 0, true},
		{"Valid token O at 1,2", 'O', 1, 2, true},
		{"Valid token at edge positions", 'X', 2, 2, true},
		{"Symbol is zero rune", 0, 0, 0, false},
		{"Negative row", 'X', -1, 0, false},
		{"Negative column", 'O', 0, -1, false},
		{"Row out of bounds", 'X', 3, 0, false},
		{"Column out of bounds", 'O', 0, 3, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			token := game.NewToken(tc.symbol, tc.row, tc.col)

			if !tc.valid {
				if token.GetSymbol() != 0 {
					t.Error("Expected zero value for symbol in invalid token")
				}
				return
			}
			if token.GetSymbol() != tc.symbol {
				t.Errorf("Expected symbol %c, got %c", tc.symbol, token.GetSymbol())
			}
		})
	}
}


func TestTokenEquality(t *testing.T) {
	tokensEqual := func(t1, t2 game.Token) bool {
		if t1.GetSymbol() == 0 && t2.GetSymbol() == 0 {
			return true
		}
		return t1.GetSymbol() == t2.GetSymbol()
	}

	tests := []struct {
		name     string
		t1       game.Token
		t2       game.Token
		expected bool
	}{
		{
			name:     "Equal tokens",
			t1:       game.NewToken('X', 0, 0),
			t2:       game.NewToken('X', 0, 0),
			expected: true,
		},
		{
			name:     "Different symbols",
			t1:       game.NewToken('X', 0, 0),
			t2:       game.NewToken('O', 0, 0),
			expected: false,
		},
		{
			name:     "Different positions same symbol",
			t1:       game.NewToken('X', 0, 0),
			t2:       game.NewToken('X', 1, 1),
			expected: true,
		},
		{
			name:     "Invalid token",
			t1:       game.NewToken('X', -1, 0),
			t2:       game.NewToken('X', 0, 0),
			expected: false,
		},
		{
			name:     "Both invalid tokens",
			t1:       game.NewToken('X', -1, 0),
			t2:       game.NewToken('Y', -1, 0),
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			isEqual := tokensEqual(tc.t1, tc.t2)
			if isEqual != tc.expected {
				t.Errorf("Expected equality to be %v for %v and %v", tc.expected, tc.t1, tc.t2)
			}
		})
	}
}


func TestTokenString(t *testing.T) {
	tests := []struct {
		name     string
		token    game.Token
		expected string
	}{
		{"X at 0,0", game.NewToken('X', 0, 0), "X"},
		{"O at 1,1", game.NewToken('O', 1, 1), "O"},
		{"A at 2,2", game.NewToken('A', 2, 2), "A"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.token.String(); got != tc.expected {
				t.Errorf("Expected string %q, got %q", tc.expected, got)
			}
		})
	}
}


func TestTokenEdgeCases(t *testing.T) {
	tests := []struct {
		name   string
		symbol rune
		row    int
		col    int
		valid  bool
	}{
		{"Minimum valid values", 'X', 0, 0, true},
		{"Maximum valid values", 'O', 2, 2, true},
		{"Non-standard symbol", 'A', 1, 1, true},
		{"Negative row", 'X', -1, 0, false},
		{"Negative column", 'O', 0, -1, false},
		{"Row out of bounds", 'X', 3, 0, false},
		{"Column out of bounds", 'O', 0, 3, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			token := game.NewToken(tc.symbol, tc.row, tc.col)

			if !tc.valid {
				if token.GetSymbol() != 0 {
					t.Error("Expected zero value for symbol in invalid token")
				}
			} else if token.GetSymbol() != tc.symbol {
				t.Errorf("Expected symbol %c, got %c", tc.symbol, token.GetSymbol())
			}
		})
	}
}
