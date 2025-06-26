package game_test

import (
	"strings"
	"testing"

	"github.com/thotluna/ttt/internal/game"
)

func TestNewBoard(t *testing.T) {
	board := game.NewBoard()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.GetBoard()[i][j] != '-' {
				t.Errorf("Expected '-' at position [%d][%d], got %c", i, j, board.GetBoard()[i][j])
			}
		}
	}
}

func TestPlaceToken(t *testing.T) {
	board := game.NewBoard()
	token := game.NewToken('X', 1, 1)

	board.PlaceToken(token)
	if board.GetBoard()[1][1] != 'X' {
		t.Error("Token not placed correctly at [1][1]")
	}

	if board.GetBoard()[0][0] != '-' || board.GetBoard()[2][2] != '-' {
		t.Error("Other cells should remain empty")
	}
}

func TestFullBoard(t *testing.T) {
	tests := []struct {
		name   string
		tokens []game.Token
		full   bool
	}{
		{
			name:   "Empty board",
			tokens: []game.Token{},
			full:   false,
		},
		{
			name: "Partially filled board",
			tokens: []game.Token{
				game.NewToken('X', 0, 0),
				game.NewToken('O', 1, 1),
			},
			full: false,
		},
		{
			name: "Full board",
			tokens: []game.Token{
				game.NewToken('X', 0, 0), game.NewToken('O', 0, 1), game.NewToken('X', 0, 2),
				game.NewToken('O', 1, 0), game.NewToken('X', 1, 1), game.NewToken('O', 1, 2),
				game.NewToken('X', 2, 0), game.NewToken('O', 2, 1), game.NewToken('X', 2, 2),
			},
			full: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			board := game.NewBoard()
			for _, token := range tc.tokens {
				board.PlaceToken(token)
			}
			if got := board.FullBoard(); got != tc.full {
				t.Errorf("Expected FullBoard() = %v, got %v", tc.full, got)
			}
		})
	}
}

func TestPlaceToken_Validation(t *testing.T) {
	tests := []struct {
		name        string
		token       game.Token
		expectError bool
		errMsg      string
	}{
		{
			name:        "valid position",
			token:       game.NewToken('X', 1, 1),
			expectError: false,
		},
		{
			name:        "out of bounds row",
			token:       game.NewToken('X', 3, 1),
			expectError: true,
			errMsg:      "position (3,1) is out of bounds",
		},
		{
			name:        "out of bounds column",
			token:       game.NewToken('X', 1, 3),
			expectError: true,
			errMsg:      "position (1,3) is out of bounds",
		},
		{
			name:        "negative row",
			token:       game.NewToken('X', -1, 1),
			expectError: true,
			errMsg:      "position (-1,1) is out of bounds",
		},
		{
			name:        "negative column",
			token:       game.NewToken('X', 1, -1),
			expectError: true,
			errMsg:      "position (1,-1) is out of bounds",
		},
		{
			name:        "position already taken",
			token:       game.NewToken('X', 0, 0),
			expectError: true,
			errMsg:      "position (0,0) is already taken",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			board := game.NewBoard()

			if tc.name == "position already taken" {
				// Place a token at (0,0) for the "position already taken" test
				err := board.PlaceToken(game.NewToken('O', 0, 0))
				if err != nil {
					t.Fatalf("Setup error: %v", err)
				}
			}

			err := board.PlaceToken(tc.token)

			if tc.expectError {
				if err == nil {
					t.Error("Expected an error but got none")
				} else if !strings.Contains(err.Error(), tc.errMsg) {
					t.Errorf("Expected error to contain '%s', got '%s'", tc.errMsg, err.Error())
				}
			} else if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestCheckWin(t *testing.T) {
	tests := []struct {
		name   string
		tokens []game.Token
		symbol rune
		win    bool
	}{
		{
			name: "No win",
			tokens: []game.Token{
				game.NewToken('X', 0, 0),
				game.NewToken('O', 1, 1),
			},
			symbol: 'X',
			win:    false,
		},
		{
			name: "Row win",
			tokens: []game.Token{
				game.NewToken('X', 0, 0),
				game.NewToken('X', 0, 1),
				game.NewToken('X', 0, 2),
			},
			symbol: 'X',
			win:    true,
		},
		{
			name: "Column win",
			tokens: []game.Token{
				game.NewToken('O', 0, 1),
				game.NewToken('O', 1, 1),
				game.NewToken('O', 2, 1),
			},
			symbol: 'O',
			win:    true,
		},
		{
			name: "Diagonal 1 win",
			tokens: []game.Token{
				game.NewToken('X', 0, 0),
				game.NewToken('X', 1, 1),
				game.NewToken('X', 2, 2),
			},
			symbol: 'X',
			win:    true,
		},
		{
			name: "Diagonal 2 win",
			tokens: []game.Token{
				game.NewToken('O', 0, 2),
				game.NewToken('O', 1, 1),
				game.NewToken('O', 2, 0),
			},
			symbol: 'O',
			win:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			board := game.NewBoard()
			for _, token := range tc.tokens {
				board.PlaceToken(token)
			}
			if got := board.CheckWin(tc.symbol); got != tc.win {
				t.Errorf("Expected CheckWin('%c') = %v, got %v", tc.symbol, tc.win, got)
			}
		})
	}
}
