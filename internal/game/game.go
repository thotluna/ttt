package game

import (
	"fmt"
	"strconv"
	"strings"

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
		_, symbolPlayer := g.turn.GetTurn()
		g.io.PrintBoard(g.board.GetBoard())
		g.io.PrintMessage(FormatPlayerTurn(symbolPlayer))

		row, col, err := g.readInput()
		if err != nil {
			g.io.PrintLine("Error: " + err.Error())
			continue
		}

		token := NewToken(symbolPlayer, row, col)
		if err := g.board.PlaceToken(token); err != nil {
			if Is(err, ErrPositionOccupied) || Is(err, ErrOutOfBounds) {
				g.io.PrintLine("Error: " + err.Error())
			} else {
				g.io.PrintLine(FormatUnexpectedError(err))
			}
			continue
		}

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

func (g *Game) readInput() (int, int, error) {
	input := g.io.ReadInput()
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
