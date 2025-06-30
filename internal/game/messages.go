package game

import "fmt"

const (
	MsgPlayerTurn = "Player %c's turn"
	MsgPlayerWins = "Player %c wins!"
	MsgGameDraw   = "It's a draw!"
)

const (
	MsgInvalidFormat       = "invalid format. Use 'row.col' (e.g. '1.2')"
	MsgRowMustBeNumber     = "row must be a number between 0 and 2"
	MsgColMustBeNumber     = "column must be a number between 0 and 2"
	MsgOutOfBounds         = "position is out of bounds (0-2,0-2)"
	MsgPositionOutOfBounds = "position (%d,%d) is out of bounds"
	MsgPositionTaken       = "position (%d,%d) is already taken"
	MsgInvalidInput        = "invalid input: %s"
	MsgInvalidRange        = "coordinates are out of range"
)

const (
	MsgUnexpectedError   = "Unexpected error: %s"
	MsgInvalidInputError = "invalid input"
	MsgPositionOccupied  = "position is already taken"
	MsgOutOfBoundsError  = "position is out of bounds"
	MsgEmptyCellError    = "position is empty"
	MsgUnknownError      = "unknown error"
)

func FormatPlayerTurn(symbol rune) string {
	return fmt.Sprintf(MsgPlayerTurn, symbol)
}

func FormatPositionTaken(row, col int) string {
	return fmt.Sprintf(MsgPositionTaken, row, col)
}

func FormatInvalidInput(msg string) string {
	return fmt.Sprintf(MsgInvalidInput, msg)
}

func FormatUnexpectedError(err error) string {
	return fmt.Sprintf(MsgUnexpectedError, err.Error())
}

func FormatPositionOutOfBounds(row, col int) string {
	return fmt.Sprintf(MsgPositionOutOfBounds, row, col)
}
