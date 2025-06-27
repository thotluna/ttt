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
	return Game{
		turn:  NewTurn(),
		board: board,
		Players: []*Player{
			NewPlayer('X', io, board),
			NewPlayer('O', io, board),
		},
		io: io,
	}
}

func (g *Game) getPlayer(indexPlayer int) *Player {
	return g.Players[indexPlayer]
}

func (g *Game) Play() {
	var isWin bool
	for {
		g.board.Print()
		indexPlayer, symbolPlayer := g.turn.GetTurn()
		g.io.PrintMessage(FormatPlayerTurn(symbolPlayer))

		isWin = g.getPlayer(indexPlayer).Put()

		if isWin {
			return
		}

		if g.board.FullBoard() {
			g.io.PrintBoard(g.board.GetBoard())
			g.io.PrintLine(MsgGameDraw)
			return
		}

		g.turn.TurnChange()

	}

}
