package game

type Token struct {
	symbol rune
	row    int
	col    int
}

func NewToken(symbol rune, row, col int) Token {
	if symbol == 0 {
		return Token{}
	}

	if row < 0 || row > 2 || col < 0 || col > 2 {
		return Token{}
	}

	return Token{
		symbol: symbol,
		row:    row,
		col:    col,
	}
}

func (t *Token) GetSymbol() rune {
	return t.symbol
}

func (t *Token) String() string {
	return string(t.symbol)
}
