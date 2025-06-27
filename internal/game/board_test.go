package game_test

import (
	"strings"
	"testing"

	"github.com/thotluna/ttt/internal/game"
	"github.com/thotluna/ttt/testutils"
)

func TestNewBoard(t *testing.T) {
	mockIO := &testutils.MockIO{}
	board := game.NewBoard(mockIO)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.GetBoard()[i][j] != '-' {
				t.Errorf("Expected '-' at position [%d][%d], got %c", i, j, board.GetBoard()[i][j])
			}
		}
	}
}

func TestPlaceToken(t *testing.T) {
	mockIO := &testutils.MockIO{}
	board := game.NewBoard(mockIO)
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
			mockIO := &testutils.MockIO{}
			board := game.NewBoard(mockIO)
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
		setup       func(*game.Board) error
		expectError bool
		errMsg      string
	}{
		{
			name:        "valid position",
			token:       game.NewToken('X', 1, 1),
			setup:       func(b *game.Board) error { return nil },
			expectError: false,
		},
		{
			name:        "invalid token",
			token:       game.Token{},
			setup:       func(b *game.Board) error { return nil },
			expectError: true,
			errMsg:      "invalid token",
		},
		{
			name:        "position already taken",
			token:       game.NewToken('X', 0, 0),
			setup:       func(b *game.Board) error { return b.PlaceToken(game.NewToken('O', 0, 0)) },
			expectError: true,
			errMsg:      "position (0,0) is already taken",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockIO := &testutils.MockIO{}
			board := game.NewBoard(mockIO)

			if err := tc.setup(board); err != nil {
				t.Fatalf("Setup error: %v", err)
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
			mockIO := &testutils.MockIO{}
			board := game.NewBoard(mockIO)
			for _, token := range tc.tokens {
				board.PlaceToken(token)
			}
			if got := board.CheckWin(tc.symbol); got != tc.win {
				t.Errorf("CheckWin() = %v, want %v", got, tc.win)
			}
		})
	}
}

func TestGetBoardReturnsCopy(t *testing.T) {
	mockIO := &testutils.MockIO{}
	board := game.NewBoard(mockIO)
	
	// Obtener una copia del tablero
	boardCopy := board.GetBoard()
	
	// Modificar la copia
	boardCopy[0][0] = 'X'
	
	// Verificar que el tablero original no fue modificado
	if board.GetBoard()[0][0] != '-' {
		t.Error("GetBoard() should return a copy of the board, not a reference")
	}
}

func TestPrint(t *testing.T) {
	mockIO := &testutils.MockIO{}
	board := game.NewBoard(mockIO)
	
	// Colocar algunas fichas para probar la salida
	board.PlaceToken(game.NewToken('X', 0, 0))
	board.PlaceToken(game.NewToken('O', 1, 1))
	
	// Llamar al método Print
	board.Print()
	
	// Verificar que se llamó a PrintBoard en el mock
	output := mockIO.GetOutput()
	if len(output) == 0 {
		t.Fatal("Expected Print to produce output")
	}
	
	// Verificar que la salida contiene las fichas colocadas
	outputStr := strings.Join(output, "\n")
	if !strings.Contains(outputStr, "X") || !strings.Contains(outputStr, "O") {
		t.Error("Expected output to contain placed tokens")
	}
}
