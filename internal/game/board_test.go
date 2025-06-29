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
	boardData := board.GetBoard()
	for i := 0; i < game.BoardSize; i++ {
		for j := 0; j < game.BoardSize; j++ {
			if boardData[i][j] != '-' {
				t.Errorf("Expected '-' at position [%d][%d], got %c", i, j, boardData[i][j])
			}
		}
	}
}

func TestPlaceToken(t *testing.T) {
	mockIO := &testutils.MockIO{}
	board := game.NewBoard(mockIO)
	coor, _ := game.NewCoordinate(1, 1)
	board.PlaceToken('X', coor)
	boardData := board.GetBoard()
	if boardData[1][1] != 'X' {
		t.Error("Token not placed correctly at [1][1]")
	}

	if boardData[0][0] != '-' || boardData[2][2] != '-' {
		t.Error("Other cells should remain empty")
	}
}

func TestFullBoard(t *testing.T) {

	tests := []struct {
		name   string
		symbol []rune
		coor   []game.Coordinate
		full   bool
	}{
		{
			name:   "Empty board",
			symbol: []rune{},
			coor:   []game.Coordinate{},
			full:   false,
		},
		{
			name:   "Partially filled board",
			symbol: []rune{'X', 'O'},
			coor: []game.Coordinate{
				testutils.MustNewCoordinate(t, 0, 0),
				testutils.MustNewCoordinate(t, 1, 1),
			},
			full: false,
		},
		{
			name:   "Full board",
			symbol: []rune{'X', 'O', 'X', 'O', 'X', 'O', 'X', 'O', 'X'},
			coor: []game.Coordinate{
				testutils.MustNewCoordinate(t, 0, 0),
				testutils.MustNewCoordinate(t, 0, 1),
				testutils.MustNewCoordinate(t, 0, 2),
				testutils.MustNewCoordinate(t, 1, 0),
				testutils.MustNewCoordinate(t, 1, 1),
				testutils.MustNewCoordinate(t, 1, 2),
				testutils.MustNewCoordinate(t, 2, 0),
				testutils.MustNewCoordinate(t, 2, 1),
				testutils.MustNewCoordinate(t, 2, 2),
			},
			full: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockIO := &testutils.MockIO{}
			board := game.NewBoard(mockIO)
			for index, coor := range tc.coor {
				board.PlaceToken(tc.symbol[index], coor)
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
		symbol      rune
		coor        game.Coordinate
		setup       func(*game.Board) error
		expectError bool
		errMsg      string
	}{
		{
			name:        "valid position",
			symbol:      'X',
			coor:        testutils.MustNewCoordinate(t, 1, 1),
			setup:       func(b *game.Board) error { return nil },
			expectError: false,
		},
		{
			name:        "position already taken",
			symbol:      'X',
			coor:        testutils.MustNewCoordinate(t, 0, 0),
			setup:       func(b *game.Board) error { return b.PlaceToken('O', testutils.MustNewCoordinate(t, 0, 0)) },
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

			err := board.PlaceToken(tc.symbol, tc.coor)

			if tc.expectError {
				if !strings.Contains(err.Error(), tc.errMsg) {
					t.Errorf("Expected error to contain '%s', got '%s'", tc.errMsg, err.Error())
				}
			} else if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestGetBoardReturnsCopy(t *testing.T) {
	mockIO := &testutils.MockIO{}
	board := game.NewBoard(mockIO)

	boardCopy := board.GetBoard()
	boardCopy[0][0] = 'X'

	// Obtener una nueva copia para verificar
	boardData := board.GetBoard()
	if boardData[0][0] != '-' {
		t.Error("GetBoard() should return a copy of the board, not a reference")
	}
}

func TestPrint(t *testing.T) {
	mockIO := &testutils.MockIO{}
	board := game.NewBoard(mockIO)

	// Colocar algunas fichas para probar la salida
	board.PlaceToken('X', testutils.MustNewCoordinate(t, 0, 0))
	board.PlaceToken('O', testutils.MustNewCoordinate(t, 1, 1))

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
