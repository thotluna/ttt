package game

import (
	"fmt"
	"strconv"
)

const (
	MsgPlayerTurn = "Player %c turn"
	MsgPlayerWins = "Player %c wins!"
	MsgGameDraw   = "It's a draw!"
)

const (
	MsgInvalidFormat       = "invalid format. Please use 'row.col' (e.g., '1.2')"
	MsgRowMustBeNumber     = "row must be a number between 0 and 2"
	MsgColMustBeNumber     = "column must be a number between 0 and 2"
	MsgOutOfBounds         = "position is out of bounds (0-2,0-2)"
	MsgPositionOutOfBounds = "position (%d,%d) is out of bounds"
	MsgPositionTaken       = "position (%d,%d) is already taken"
	MsgInvalidInput        = "invalid input: %s"
)

const (
	MsgUnexpectedError   = "Unexpected error: %s"
	MsgInvalidInputError = "invalid input"
	MsgPositionOccupied  = "position already occupied"
	MsgOutOfBoundsError  = "position out of bounds"
	MsgUnknownError      = "unknown error"
)

func FormatPlayerTurn(symbol rune) string {
	return "Player " + string(symbol) + " turn"
}

func FormatPositionTaken(row, col int) string {
	return "position already taken: position (" + strconv.Itoa(row) + "," + strconv.Itoa(col) + ") is already taken"
}

func FormatInvalidInput(msg string) string {
	return "invalid input: " + msg
}

func FormatUnexpectedError(err error) string {
	return "Unexpected error: " + err.Error()
}

func FormatPositionOutOfBounds(row, col int) string {
	return fmt.Sprintf(MsgPositionOutOfBounds, row, col)
}
