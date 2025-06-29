package game

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thotluna/ttt/internal/view"
)

type Player struct {
	symbol rune
	io     view.IO
	board  *Board
}

func NewPlayer(symbol rune, io view.IO, board *Board) *Player {
	return &Player{
		symbol: symbol,
		io:     io,
		board:  board,
	}
}

func (p *Player) Play() bool {
	for {
		row, col := p.readInput()

		coor, err := NewCoordinate(row, col)
		if err != nil {
			p.io.PrintLine(err.Error())
			continue
		}

		err = p.board.PlaceToken(p.symbol, coor)
		if err != nil {
			p.io.PrintLine(err.Error())
			continue
		}
		break
	}
	return p.CheckWin()
}

func (p *Player) CheckWin() bool {
	tokens := p.board.GetTokenBy(p.symbol)
	if len(tokens) < 3 {

		return false
	}

	direction := make(map[Direction]int)

	for i := 0; i < len(tokens); i++ {
		for j := i + 1; j < len(tokens); j++ {
			dir := tokens[i].Direction(tokens[j])
			if dir != None {
				direction[dir]++
			}
		}
	}

	for _, v := range direction {
		if v >= 3 {
			p.io.PrintLine(fmt.Sprintf(MsgPlayerWins, p.symbol))
			return true
		}
	}

	return false
}

func (p *Player) readInput() (int, int) {
	var row, col int
	var err error
	for {
		input := p.io.ReadInput()
		input = strings.TrimSpace(input)
		parts := strings.Split(input, ".")
		if len(parts) != 2 {
			p.io.PrintLine(NewGameError(ErrInvalidInput, MsgInvalidFormat).Error())
			continue
		}

		row, err = strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			p.io.PrintLine(NewGameError(ErrInvalidInput, MsgRowMustBeNumber).Error())
			continue
		}

		col, err = strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			p.io.PrintLine(NewGameError(ErrInvalidInput, MsgColMustBeNumber).Error())
			continue
		}

		if row < 0 || row > 2 || col < 0 || col > 2 {
			p.io.PrintLine(NewGameError(ErrOutOfBounds, MsgOutOfBounds).Error())
			continue
		}
		break
	}

	return row, col
}
