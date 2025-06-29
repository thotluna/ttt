package game

import "github.com/thotluna/ttt/internal/validator"

const (
	minBoardPosition = 0
	maxBoardPosition = 2
)

type GameInterval struct {
	*validator.Interval
}

func NewGameInterval() *GameInterval {
	return &GameInterval{
		Interval: validator.NewInterval(minBoardPosition, maxBoardPosition),
	}
}
