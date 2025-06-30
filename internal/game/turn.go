package game

import "github.com/thotluna/ttt/internal/view"

const (
	PlayerXIndex    = 0
	PlayerOIndex    = 1
	NumberOfPlayers = 2
)

type Turn struct {
	symbolCurrent Symbol
	players       map[Symbol]Player
	io            view.IO
}

func NewTurn(players map[Symbol]Player, io view.IO) Turn {
	return Turn{
		symbolCurrent: PlayerX,
		players:       players,
		io:            io,
	}
}

func (t *Turn) TurnChange() {
	switch t.symbolCurrent {
	case PlayerX:
		t.symbolCurrent = PlayerO
	case PlayerO:
		t.symbolCurrent = PlayerX
	}
}

func (t *Turn) GetTurn() Symbol {
	switch t.symbolCurrent {
	case PlayerX:
		return PlayerX
	case PlayerO:
		return PlayerO
	default:
		return PlayerX
	}
}

func (t *Turn) GetCurrentPlayer() *Player {
	player := t.players[t.symbolCurrent]
	return &player
}

func (t *Turn) PrintTurn() {
	t.io.PrintMessage(FormatPlayerTurn(rune(t.GetTurn())))
}
