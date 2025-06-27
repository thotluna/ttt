package game

import (
	"github.com/thotluna/ttt/internal/view"
)

type Game struct {
	turn    Turn
	board   *Board
	Players []*Player
	io      view.IO
}

func NewGame(io view.IO) Game {
	board := NewBoard(io)
	players := []*Player{
		NewPlayer('X', io, board),
		NewPlayer('O', io, board),
	}
	return Game{
		turn:    NewTurn(players, io),
		board:   board,
		Players: players,
		io:      io,
	}
}

func (g *Game) Play() {

	for {
		g.board.Print()
		g.turn.PrintTurn()

		currentPlayer := g.turn.GetCurrentPlayer()
		isWin := currentPlayer.Put()

		if isWin {
			return
		}

		if g.board.FullBoard() {
			g.board.Print()
			g.io.PrintLine(MsgGameDraw)
			return
		}

		g.turn.TurnChange()

	}

}
