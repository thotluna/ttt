package game_test

import (
	"strings"
	"testing"

	"github.com/thotluna/ttt/internal/game"
	"github.com/thotluna/ttt/testutils"
)

func mustNewCoordinate(row, col int) game.Coordinate {
	coord, err := game.NewCoordinate(row, col)
	if err != nil {
		panic(err)
	}
	return coord
}

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
	coor, _ := game.NewCoordinate(1, 1)
	board.PlaceToken('X', coor)
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
				mustNewCoordinate(0, 0),
				mustNewCoordinate(1, 1),
			},
			full: false,
		},
		{
			name:   "Full board",
			symbol: []rune{'X', 'O', 'X', 'O', 'X', 'O', 'X', 'O', 'X'},
			coor: []game.Coordinate{
				mustNewCoordinate(0, 0),
				mustNewCoordinate(0, 1),
				mustNewCoordinate(0, 2),
				mustNewCoordinate(1, 0),
				mustNewCoordinate(1, 1),
				mustNewCoordinate(1, 2),
				mustNewCoordinate(2, 0),
				mustNewCoordinate(2, 1),
				mustNewCoordinate(2, 2),
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
			coor:        mustNewCoordinate(1, 1),
			setup:       func(b *game.Board) error { return nil },
			expectError: false,
		},
		{
			name:        "position already taken",
			symbol:      'X',
			coor:        mustNewCoordinate(0, 0),
			setup:       func(b *game.Board) error { return b.PlaceToken('O', mustNewCoordinate(0, 0)) },
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

func TestCheckWin(t *testing.T) {
	tests := []struct {
		name   string
		symbol []rune
		coor   []game.Coordinate
		win    bool
	}{
		{
			name:   "No win",
			symbol: []rune{'X', 'O'},
			coor: []game.Coordinate{
				mustNewCoordinate(0, 0),
				mustNewCoordinate(1, 1),
			},
			win: false,
		},
		{
			name:   "Row win",
			symbol: []rune{'X', 'X', 'X'},
			coor: []game.Coordinate{
				mustNewCoordinate(0, 0),
				mustNewCoordinate(0, 1),
				mustNewCoordinate(0, 2),
			},
			win: true,
		},
		{
			name:   "Column win",
			symbol: []rune{'O', 'O', 'O'},
			coor: []game.Coordinate{
				mustNewCoordinate(0, 1),
				mustNewCoordinate(1, 1),
				mustNewCoordinate(2, 1),
			},
			win: true,
		},
		{
			name:   "Diagonal 1 win",
			symbol: []rune{'X', 'X', 'X'},
			coor: []game.Coordinate{
				mustNewCoordinate(0, 0),
				mustNewCoordinate(1, 1),
				mustNewCoordinate(2, 2),
			},
			win: true,
		},
		{
			name:   "Diagonal 2 win",
			symbol: []rune{'O', 'O', 'O'},
			coor: []game.Coordinate{
				mustNewCoordinate(0, 2),
				mustNewCoordinate(1, 1),
				mustNewCoordinate(2, 0),
			},
			win: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockIO := &testutils.MockIO{}
			board := game.NewBoard(mockIO)
			for _, coor := range tc.coor {
				board.PlaceToken(tc.symbol[0], coor)
			}
			if got := board.CheckWin(tc.symbol[0]); got != tc.win {
				t.Errorf("CheckWin() = %v, want %v", got, tc.win)
			}
		})
	}
}

func TestGetBoardReturnsCopy(t *testing.T) {
	mockIO := &testutils.MockIO{}
	board := game.NewBoard(mockIO)

	boardCopy := board.GetBoard()

	boardCopy[0][0] = 'X'

	if board.GetBoard()[0][0] != '-' {
		t.Error("GetBoard() should return a copy of the board, not a reference")
	}
}

func TestPrint(t *testing.T) {
	mockIO := &testutils.MockIO{}
	board := game.NewBoard(mockIO)

	// Colocar algunas fichas para probar la salida
	board.PlaceToken('X', mustNewCoordinate(0, 0))
	board.PlaceToken('O', mustNewCoordinate(1, 1))

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
