package game

import (
	"strconv"
	"strings"

	"github.com/thotluna/ttt/internal/constans"
	"github.com/thotluna/ttt/internal/view"
)

type Game struct {
	turn  Turn
	board *Board
	io    view.IO
}

func NewGame(io view.IO) Game {
	return Game{
		turn:  NewTurn(),
		board: NewBoard(),
		io:    io,
	}
}
func (g *Game) Play() {
	for {
		_, player := g.turn.GetTurn()
		g.io.PrintBoard(g.board.GetBoard())
		g.io.PrintMessage(constans.FormatPlayerTurn(player))

		row, col, err := g.readInput()
		if err != nil {
			g.io.PrintLine("Error: " + err.Error())
			continue
		}

		token := NewToken(player, row, col)
		if err := g.board.PlaceToken(token); err != nil {
			if Is(err, ErrPositionOccupied) || Is(err, ErrOutOfBounds) {
				g.io.PrintLine("Error: " + err.Error())
			} else {
				g.io.PrintLine(constans.FormatUnexpectedError(err))
			}
			continue
		}

		if g.board.CheckWin(player) {
			g.io.PrintBoard(g.board.GetBoard())
			g.io.PrintWin(player)
			return
		}

		if g.board.FullBoard() {
			g.io.PrintBoard(g.board.GetBoard())
			g.io.PrintDraw()
			return
		}

		g.turn.TurnChange()

	}

}

func (g *Game) readInput() (int, int, error) {
	input := g.io.ReadInput()
	input = strings.TrimSpace(input)
	parts := strings.Split(input, ".")
	if len(parts) != 2 {
		return 0, 0, NewGameError(ErrInvalidInput, constans.MsgInvalidFormat)
	}

	row, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, NewGameError(ErrInvalidInput, constans.MsgRowMustBeNumber)
	}

	col, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, NewGameError(ErrInvalidInput, constans.MsgColMustBeNumber)
	}

	if row < 0 || row > 2 || col < 0 || col > 2 {
		return 0, 0, NewGameError(ErrOutOfBounds, constans.MsgOutOfBounds)
	}

	return row, col, nil
}
