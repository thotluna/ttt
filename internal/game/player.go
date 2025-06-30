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
	move   int
}

func NewPlayer(symbol Symbol, io view.IO, board *Board) *Player {
	return &Player{
		symbol: symbol,
		board:  board,
		input:  NewPlayerInput(io),
		move:   0,
	}
}

func (p *Player) Play() bool {
	for {

		var origin *Coordinate
		var err error
		if p.move >= 3 {
			p.input.io.PrintLine(MsgMoveTokenPrompt)

			originRaw, err := p.input.GetMove()
			origin = &originRaw
			if err != nil {
				continue
			}

		} else {
			origin = nil
		}

		p.input.io.PrintLine(MsgPlaceTokenPrompt)
		destination, err := p.input.GetMove()
		if err != nil {
			continue
		}

		err = p.board.PlaceToken(p.symbol, origin, &destination)
		if err != nil {
			p.input.io.PrintLine(err.Error())
			continue
		}
		break
	}
	p.move++
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

func (p *Player) MoveCount() int {
	return p.move
}
