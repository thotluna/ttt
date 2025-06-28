package game_test

import (
	"testing"

	"github.com/thotluna/ttt/internal/game"
	"github.com/thotluna/ttt/testutils"
)

func TestNewCoordinate(t *testing.T) {
	tests := []struct {
		name    string
		row     int
		col     int
		hasErr  bool
		errCode game.ErrorCode
		wantRow int
		wantCol int
	}{
		{"valid coordinate", 0, 0, false, 0, 0, 0},
		{"valid coordinate middle", 1, 1, false, 0, 1, 1},
		{"valid coordinate edge", 2, 2, false, 0, 2, 2},
		{"invalid row negative", -1, 0, true, game.ErrOutOfBounds, 0, 0},
		{"invalid row too large", 3, 0, true, game.ErrOutOfBounds, 0, 0},
		{"invalid col negative", 0, -1, true, game.ErrOutOfBounds, 0, 0},
		{"invalid col too large", 0, 3, true, game.ErrOutOfBounds, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := game.NewCoordinate(tt.row, tt.col)
			if tt.hasErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if gameErr, ok := err.(*game.GameError); !ok || gameErr.Code != tt.errCode {
					t.Errorf("expected error code %v, got %v", tt.errCode, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got.Row() != tt.wantRow || got.Col() != tt.wantCol {
				t.Errorf("expected (%d,%d), got (%d,%d)", tt.wantRow, tt.wantCol, got.Row(), got.Col())
			}
		})
	}
}

func TestCoordinate_IsValid(t *testing.T) {
	tests := []struct {
		name string
		row  int
		col  int
		want bool
	}{
		{"valid coordinate", 0, 0, true},
		{"valid coordinate middle", 1, 1, true},
		{"valid coordinate edge", 2, 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coord, err := game.NewCoordinate(tt.row, tt.col)
			if err != nil {
				t.Fatalf("Failed to create coordinate: %v", err)
			}
			if got := coord.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}

	invalidTests := []struct {
		name string
		row  int
		col  int
	}{
		{"invalid row negative", -1, 0},
		{"invalid row too large", 3, 0},
		{"invalid col negative", 0, -1},
		{"invalid col too large", 0, 3},
	}

	for _, tt := range invalidTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := game.NewCoordinate(tt.row, tt.col)
			if err == nil {
				t.Error("Expected an error for invalid coordinate, but got none")
			}
		})
	}
}

func TestCoordinate_Equals(t *testing.T) {
	coord1 := testutils.MustNewCoordinate(t, 1, 1)

	tests := []struct {
		name string
		c1   game.Coordinate
		c2   game.Coordinate
		want bool
	}{
		{"equal coordinates", coord1, testutils.MustNewCoordinate(t, 1, 1), true},
		{"different row", coord1, testutils.MustNewCoordinate(t, 0, 1), false},
		{"different col", coord1, testutils.MustNewCoordinate(t, 1, 0), false},
		{"completely different", coord1, testutils.MustNewCoordinate(t, 0, 0), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c1.Equals(tt.c2); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoordinate_Directions(t *testing.T) {
	center := testutils.MustNewCoordinate(t, 1, 1)

	tests := []struct {
		name      string
		c1        game.Coordinate
		c2        game.Coordinate
		isVert    bool
		isHoriz   bool
		isDiag    bool
		isInvert  bool
		direction game.Direction
	}{
		{
			"vertical up",
			center,
			testutils.MustNewCoordinate(t, 0, 1),
			true, false, false, false,
			game.Vertical,
		},
		{
			"vertical down",
			center,
			testutils.MustNewCoordinate(t, 2, 1),
			true, false, false, false,
			game.Vertical,
		},
		{
			"horizontal left",
			center,
			testutils.MustNewCoordinate(t, 1, 0),
			false, true, false, false,
			game.Horizontal,
		},
		{
			"horizontal right",
			center,
			testutils.MustNewCoordinate(t, 1, 2),
			false, true, false, false,
			game.Horizontal,
		},
		{
			"diagonal top-left",
			testutils.MustNewCoordinate(t, 0, 0),
			testutils.MustNewCoordinate(t, 2, 2),
			false, false, true, false,
			game.Diagonal,
		},
		{
			"inverter top-right",
			testutils.MustNewCoordinate(t, 0, 2),
			testutils.MustNewCoordinate(t, 2, 0),
			false, false, false, true,
			game.Inverter,
		},
		{
			"no direction",
			testutils.MustNewCoordinate(t, 0, 0),
			testutils.MustNewCoordinate(t, 2, 1),
			false, false, false, false,
			game.None,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c1.IsVerticalTo(tt.c2); got != tt.isVert {
				t.Errorf("IsVerticalTo() = %v, want %v", got, tt.isVert)
			}

			if got := tt.c1.IsHorizontalTo(tt.c2); got != tt.isHoriz {
				t.Errorf("IsHorizontalTo() = %v, want %v", got, tt.isHoriz)
			}

			if got := tt.c1.IsDiagonalTo(tt.c2); got != tt.isDiag {
				t.Errorf("IsDiagonalTo() = %v, want %v", got, tt.isDiag)
			}

			if got := tt.c1.IsInverterTo(tt.c2); got != tt.isInvert {
				t.Errorf("IsInverterTo() = %v, want %v", got, tt.isInvert)
			}

			if got := tt.c1.Direction(tt.c2); got != tt.direction {
				t.Errorf("Direction() = %v, want %v", got, tt.direction)
			}
		})
	}
}
