package game

// IO define la interfaz para la entrada/salida del juego
type IO interface {
	ReadInput() string
	PrintLine(line string)
	PrintMessage(message string)
	PrintBoard(board [3][3]rune)
	PrintWin(player rune)
	PrintDraw()
}
