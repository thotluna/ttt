package game

import (
	"fmt"

	"github.com/thotluna/ttt/internal/view"
)

type Game struct {
	turn    Turn
	board   *Board
	Players []*Player
	io      view.IO
}

func NewGame(io view.IO) Game {
	return Game{
		turn:  NewTurn(),
		board: NewBoard(),
		Players: []*Player{
			NewPlayer('X', io),
			NewPlayer('O', io),
		},
		io: io,
	}
}

func (g *Game) getPlayer(indexPlayer int) *Player {
	return g.Players[indexPlayer]
}

func (g *Game) Play() {
	for {
		g.io.PrintBoard(g.board.GetBoard())
		indexPlayer, symbolPlayer := g.turn.GetTurn()
		g.io.PrintMessage(FormatPlayerTurn(symbolPlayer))

		playerCurrent := g.getPlayer(indexPlayer)
		playerCurrent.Put(g.board)

		if g.board.CheckWin(symbolPlayer) {
			g.io.PrintBoard(g.board.GetBoard())
			g.io.PrintLine(fmt.Sprintf(MsgPlayerWins, symbolPlayer))
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
