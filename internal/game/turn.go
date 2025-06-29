package game

import "github.com/thotluna/ttt/internal/view"

const (
	PlayerXIndex    = 0
	PlayerOIndex    = 1
	NumberOfPlayers = 2
)

type Turn struct {
	symbol  SymbolPlayerCurrent
	players map[SymbolPlayerCurrent]Player
	io      view.IO
}

func NewTurn(players map[SymbolPlayerCurrent]Player, io view.IO) Turn {
	return Turn{
		symbol:  PlayerX,
		players: players,
		io:      io,
	}
}

func (t *Turn) TurnChange() {
	switch t.symbol {
	case PlayerX:
		t.symbol = PlayerO
	case PlayerO:
		t.symbol = PlayerX
	}
	switch t.symbol {
	case PlayerX:
		t.symbol = PlayerO
	case PlayerO:
		t.symbol = PlayerX
	}
}

func (t *Turn) GetTurn() (int, rune) {
	switch t.symbol {
	case PlayerX:
		return PlayerXIndex, rune(PlayerX)
	case PlayerO:
		return PlayerOIndex, rune(PlayerO)
	switch t.symbol {
	case PlayerX:
		return PlayerXIndex, rune(PlayerX)
	case PlayerO:
		return PlayerOIndex, rune(PlayerO)
	default:
		return PlayerXIndex, rune(PlayerX)
		return PlayerXIndex, rune(PlayerX)
	}
}

func (t *Turn) GetCurrentPlayer() *Player {
	player := t.players[t.symbol]
	return &player
}

func (t *Turn) PrintTurn() {
	_, symbol := t.GetTurn()
	t.io.PrintMessage(FormatPlayerTurn(symbol))
}
