package view

type IO interface {
	ReadInput() string
	PrintLine(string)
	PrintMessage(string)
	PrintBoard(board [3][3]rune)
	PrintWin(player rune)
	PrintDraw()
}
