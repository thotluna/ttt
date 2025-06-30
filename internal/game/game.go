package game

import (
	"github.com/thotluna/ttt/internal/view"
)

type Game struct {
	turn  Turn
	board *Board
	io    view.IO
}

func NewGame(io view.IO) Game {
	board := NewBoard(io)
	players := map[Symbol]Player{
		PlayerX: *NewPlayer(PlayerX, io, board),
		PlayerO: *NewPlayer(PlayerO, io, board),
	}
	return Game{
		turn:  NewTurn(players, io),
		board: board,
		io:    io,
	}
}

func (g *Game) Play() {

	for {
		g.board.Print()
		g.turn.PrintTurn()

		currentPlayer := g.turn.GetCurrentPlayer()
		isWin := currentPlayer.Play()

		if isWin {
			return
		}

		if g.board.IsFull() {
			g.board.Print()
			g.io.PrintLine(MsgGameDraw)
			return
		}

		g.turn.TurnChange()

	}

}
