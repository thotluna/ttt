package game

import (
	"strconv"
	"strings"

	"github.com/thotluna/ttt/view"
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
		g.io.PrintBoard(g.board.board)
		g.io.PrintMessage("Player " + string(player) + " turn")

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
				g.io.PrintLine("Unexpected error: " + err.Error())
			}
			continue
		}

		if g.board.CheckWin(player) {
			g.io.PrintBoard(g.board.board)
			g.io.PrintWin(player)
			return
		}

		if g.board.FullBoard() {
			g.io.PrintBoard(g.board.board)
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
		return 0, 0, NewGameError(ErrInvalidInput, "invalid format. Please use 'row.col' (e.g., '1.2')")
	}

	row, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, NewGameError(ErrInvalidInput, "row must be a number between 0 and 2")
	}

	col, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, NewGameError(ErrInvalidInput, "column must be a number between 0 and 2")
	}

	if row < 0 || row > 2 || col < 0 || col > 2 {
		return 0, 0, NewGameError(ErrOutOfBounds, "position is out of bounds (0-2,0-2)")
	}

	return row, col, nil
}
