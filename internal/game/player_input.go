package game

import (
	"github.com/thotluna/ttt/internal/view"
	"github.com/thotluna/ttt/internal/validator"
	"strconv"
	"strings"
)

type PlayerInput struct {
	io        view.IO
	validator *validator.InputCoordinateValidator
}

func NewPlayerInput(io view.IO) *PlayerInput {
	return &PlayerInput{
		io:        io,
		validator: validator.NewInputCoordinateValidator(nil),
	}
}

func (p *PlayerInput) GetMove() (Coordinate, error) {
	for {
		input := p.io.ReadInput()
		input = strings.TrimSpace(input)

		if err := p.validator.Validate(input); err != nil {
			p.io.PrintLine(err.Error())
			continue
		}

		row, col, err := p.parseCoordinate(input)
		if err != nil {
			p.io.PrintLine(err.Error())
			continue
		}

		coord, err := NewCoordinate(row, col)
		if err != nil {
			p.io.PrintLine(MsgOutOfBounds)
			continue
		}

		return coord, nil
	}
}

func (p *PlayerInput) parseCoordinate(input string) (int, int, error) {
	parts := strings.Split(input, ".")
	if len(parts) != 2 {
		return 0, 0, NewGameError(ErrInvalidInput, MsgInvalidFormat)
	}

	row, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, NewGameError(ErrInvalidInput, MsgRowMustBeNumber)
	}

	col, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, NewGameError(ErrInvalidInput, MsgColMustBeNumber)
	}

	return row, col, nil
}
