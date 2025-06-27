package game

import (
	"strconv"
	"strings"

	"github.com/thotluna/ttt/internal/view"
)

type Player struct {
	symbol rune
	io     view.IO
}

func NewPlayer(symbol rune, io view.IO) *Player {
	return &Player{
		symbol: symbol,
		io:     io,
	}
}

func (p *Player) Put(board *Board) {
	for {
		row, col, err := p.readInput()
		if err != nil {
			p.io.PrintLine(err.Error())
			continue
		}

		err = board.PlaceToken(NewToken(p.symbol, row, col))
		if err != nil {
			p.io.PrintLine(err.Error())
			continue
		}

		break
	}
}

func (p *Player) readInput() (int, int, error) {
	input := p.io.ReadInput()
	input = strings.TrimSpace(input)
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

	if row < 0 || row > 2 || col < 0 || col > 2 {
		return 0, 0, NewGameError(ErrOutOfBounds, MsgOutOfBounds)
	}

	return row, col, nil
}
