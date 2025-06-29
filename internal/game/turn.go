package game

import "github.com/thotluna/ttt/internal/view"

const (
	PlayerXIndex    = 0
	PlayerOIndex    = 1
	NumberOfPlayers = 2
)

type Turn struct {
	symbol  SymbolPlayerCurrent
	players map[SymbolPlayerCurrent]*Player
	io      view.IO
}

func NewTurn(players []*Player, io view.IO) Turn {
	return Turn{
		symbol: PlayerX,
		players: map[SymbolPlayerCurrent]*Player{
			PlayerX: players[0],
			PlayerO: players[1],
		},
		io: io,
	}
}

func (t *Turn) TurnChange() {
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
	default:
		return PlayerXIndex, rune(PlayerX)
	}
}

func (t *Turn) GetCurrentPlayer() *Player {
	return t.players[t.symbol]
}

func (t *Turn) PrintTurn() {
	_, symbol := t.GetTurn()
	t.io.PrintMessage(FormatPlayerTurn(symbol))
}
