package game

type Token struct {
	symbol rune
	row    int
	col    int
}

func NewToken(symbol rune, row int, col int) Token {
	return Token{
		symbol: symbol,
		row:    row,
		col:    col,
	}
}

func (t *Token) GetSymbol() rune {
	return t.symbol
}
