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

	for i := 0; i < game.BoardSize; i++ {
		for j := 0; j < game.BoardSize; j++ {
			coor := testutils.MustNewCoordinate(t, i, j)
			if !board.IsFull() && !board.IsCellOccupiedBy(*coor, game.EmptyCell) {
				t.Errorf("Expected empty cell at position [%d][%d]", i, j)
			}
		}
	}
}

func TestPlaceToken(t *testing.T) {
	mockIO := &testutils.MockIO{}
	board := game.NewBoard(mockIO)
	coor := testutils.MustNewCoordinate(t, 1, 1)

	if board.IsCellOccupiedBy(*coor, game.PlayerX) {
		t.Error("Cell should be empty initially")
	}

	// Colocar ficha y verificar
	err := board.PlaceToken(game.PlayerX, nil, coor)
	if err != nil {
		t.Fatalf("Error placing token: %v", err)
	}

	// Verificar que la celda tiene la X
	if !board.IsCellOccupiedBy(*coor, game.PlayerX) {
		t.Error("Token 'X' not placed correctly at [1][1]")
	}

	// Verificar que otras celdas sigan vacías
	otherCoords := []*game.Coordinate{
		testutils.MustNewCoordinate(t, 0, 0),
		testutils.MustNewCoordinate(t, 2, 2),
	}

	for _, c := range otherCoords {
		if board.IsCellOccupiedBy(*c, game.PlayerX) || board.IsCellOccupiedBy(*c, game.PlayerO) {
			t.Errorf("Cell [%d][%d] should be empty", c.Row(), c.Col())
		}
	}
}

func TestFullBoard(t *testing.T) {

	tests := []struct {
		name   string
		symbol []game.Symbol
		coor   []*game.Coordinate
		full   bool
	}{
		{
			name:   "Empty board",
			symbol: []game.Symbol{},
			coor:   []*game.Coordinate{},
			full:   false,
		},
		{
			name:   "Partially filled board",
			symbol: []game.Symbol{game.PlayerX, game.PlayerO},
			coor: []*game.Coordinate{
				testutils.MustNewCoordinate(t, 0, 0),
				testutils.MustNewCoordinate(t, 1, 1),
			},
			full: false,
		},
		{
			name:   "Full board",
			symbol: []game.Symbol{game.PlayerX, game.PlayerO, game.PlayerX, game.PlayerO, game.PlayerX, game.PlayerO, game.PlayerX, game.PlayerO, game.PlayerX},
			coor: []*game.Coordinate{
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
				board.PlaceToken(tc.symbol[index], nil, coor)
			}
			if got := board.IsFull(); got != tc.full {
				t.Errorf("Expected FullBoard() = %v, got %v", tc.full, got)
			}
		})
	}
}

func TestPlaceToken_Validation(t *testing.T) {
	tests := []struct {
		name        string
		symbol      game.Symbol
		coor        *game.Coordinate
		setup       func(*game.Board) error
		expectError bool
		errMsg      string
	}{
		{
			name:        "valid position",
			symbol:      game.PlayerX,
			coor:        testutils.MustNewCoordinate(t, 1, 1),
			setup:       func(b *game.Board) error { return nil },
			expectError: false,
		},
		{
			name:   "position already taken",
			symbol: game.PlayerX,
			coor:   testutils.MustNewCoordinate(t, 0, 0),
			setup: func(b *game.Board) error {
				return b.PlaceToken(game.PlayerO, nil, testutils.MustNewCoordinate(t, 0, 0))
			},
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

			err := board.PlaceToken(tc.symbol, nil, tc.coor)

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

func TestPrint(t *testing.T) {
	mockIO := &testutils.MockIO{}
	board := game.NewBoard(mockIO)

	// Colocar algunas fichas para probar la salida
	board.PlaceToken('X', nil, testutils.MustNewCoordinate(t, 0, 0))
	board.PlaceToken('O', nil, testutils.MustNewCoordinate(t, 1, 1))

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
