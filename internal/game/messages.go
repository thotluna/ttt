package game

import "fmt"

const (
	// Game flow messages
	MsgPlayerTurn = "Player %c's turn"
	MsgPlayerWins = "Player %c wins!"
	MsgGameDraw   = "Game ended in a draw!"
)

const (
	// Input validation messages
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
	// Error messages
	MsgInvalidInputError = "invalid input"
	MsgPositionOccupied  = "position is already taken"
	MsgOutOfBoundsError  = "position is out of bounds"
	MsgEmptyCellError    = "position is empty"
	MsgUnknownError      = "unknown error"
	
	// Move token messages
	MsgMoveTokenPrompt    = "Enter coordinates to move from (row.col): "
	MsgPlaceTokenPrompt   = "Enter coordinates to place token (row.col): "
	MsgNotYourToken       = "the token at this position is not yours"
	MsgDestinationOccupied = "the destination position is already occupied"
	MsgNoTokenAtPosition  = "there is no token at the specified position"
	MsgSamePosition       = "cannot move to the same position"
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
	return fmt.Sprintf("%s: %s", MsgUnknownError, err.Error())
}

func FormatPositionOutOfBounds(row, col int) string {
	return fmt.Sprintf(MsgPositionOutOfBounds, row, col)
}
