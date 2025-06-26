package game

type Token struct {
	token rune
	row   int
	col   int
}

func NewToken(symbol rune, row int, col int) Token {
	return Token{
		token: symbol,
		row:   row,
		col:   col,
	}
}

func (t *Token) GetSymbol() rune {
	return t.token
}
