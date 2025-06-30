package game

import (
	"fmt"

	"github.com/thotluna/ttt/internal/view"
)

const (
	MinTokensForWin = 3
)

type Player struct {
	symbol Symbol
	board  *Board
	input  *PlayerInput
}

func NewPlayer(symbol Symbol, io view.IO, board *Board) *Player {
	return &Player{
		symbol: symbol,
		board:  board,
		input:  NewPlayerInput(io),
	}
}

func (p *Player) Play() bool {
	for {
		if p.board.IsFull() {
			p.input.io.PrintLine(MsgGameDraw)
			return false
		}

		coord, err := p.input.GetMove()
		if err != nil {
			continue
		}

		err = p.board.PlaceToken(p.symbol, coord)
		if err != nil {
			p.input.io.PrintLine(err.Error())
			continue
		}
		break
	}
	return p.CheckWin()
}

func (p *Player) CheckWin() bool {
	tokens := p.board.GetTokenBy(p.symbol)
	if len(tokens) < MinTokensForWin {

		return false
	}

	return p.hasWinningLine(tokens)

}

func (p *Player) hasWinningLine(tokens []Coordinate) bool {
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
		if v >= MinTokensForWin {
			p.input.io.PrintLine(fmt.Sprintf(MsgPlayerWins, p.symbol))
			return true
		}
	}

	return false
}
