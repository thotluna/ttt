package testutils

import (
	"testing"

	"github.com/thotluna/ttt/internal/game"
)

// MustNewCoordinate creates a new coordinate and fails the test if there is an error.
// It's a helper function for use in tests.
func MustNewCoordinate(t testing.TB, row, col int) game.Coordinate {
	t.Helper()
	coord, err := game.NewCoordinate(row, col)
	if err != nil {
		t.Fatalf("failed to create coordinate (%d,%d): %v", row, col, err)
	}
	return coord
}
