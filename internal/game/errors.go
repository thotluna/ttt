package game

import "fmt"

// GameError represents a custom error type for game-related errors
type GameError struct {
	Code    ErrorCode
	Message string
	Details string
}

// ErrorCode represents different types of game errors
type ErrorCode int

const (
	// ErrInvalidInput represents invalid user input
	ErrInvalidInput ErrorCode = iota + 1
	// ErrPositionOccupied represents an attempt to place a token in an occupied position
	ErrPositionOccupied
	// ErrOutOfBounds represents an attempt to access a position outside the board
	ErrOutOfBounds
)

// Error returns the error message
func (e *GameError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Details)
	}
	return e.Message
}

// NewGameError creates a new GameError with the given code and details
func NewGameError(code ErrorCode, details string) *GameError {
	err := &GameError{
		Code:    code,
		Details: details,
	}

	switch code {
	case ErrInvalidInput:
		err.Message = "invalid input"
	case ErrPositionOccupied:
		err.Message = "position already occupied"
	case ErrOutOfBounds:
		err.Message = "position out of bounds"
	default:
		err.Message = "unknown error"
	}

	return err
}

// Is checks if the error is of type GameError with the given code
func Is(err error, code ErrorCode) bool {
	if gameErr, ok := err.(*GameError); ok {
		return gameErr.Code == code
	}
	return false
}
