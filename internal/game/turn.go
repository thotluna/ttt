package game

type Turn struct {
	turn int
}

func NewTurn() Turn {
	return Turn{
		turn: 0,
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
