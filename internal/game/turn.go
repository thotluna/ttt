package game

import "github.com/thotluna/ttt/internal/view"

const (
	PlayerXIndex    = 0
	PlayerOIndex    = 1
	NumberOfPlayers = 2
)

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
	t.turn = (t.turn + 1) % NumberOfPlayers
}

func (t *Turn) GetTurn() (int, rune) {
	switch t.turn {
	case PlayerXIndex:
		return PlayerXIndex, 'X'
	case PlayerOIndex:
		return PlayerOIndex, 'O'
	default:
		return PlayerXIndex, 'X'
	}
}

func (t *Turn) GetCurrentPlayer() *Player {
	return t.players[t.turn]
}

func (t *Turn) PrintTurn() {
	_, symbol := t.GetTurn()
	t.io.PrintMessage(FormatPlayerTurn(symbol))
}
