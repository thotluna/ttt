package game

import "github.com/thotluna/ttt/internal/view"

type Turn struct {
	turn    int
	players []*Player
	io      view.IO
}

func NewTurn(players []*Player, io view.IO) Turn {
	return Turn{
		turn:    0,
		players: players,
		io:      io,
	}
}

func (t *Turn) TurnChange() {
	t.turn = (t.turn + 1) % 2
}

func (t *Turn) GetTurn() (int, rune) {
	switch t.turn {
	case 0:
		return 0, 'X'
	case 1:
		return 1, 'O'
	default:
		return 0, 'X'
	}
}

func (t *Turn) GetCurrentPlayer() *Player {
	return t.players[t.turn]
}

func (t *Turn) PrintTurn() {
	_, symbol := t.GetTurn()
	t.io.PrintMessage(FormatPlayerTurn(symbol))
}
