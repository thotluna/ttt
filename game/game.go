package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	turn  Turn
	board *Board
}

func NewGame() Game {
	return Game{
		turn:  NewTurn(),
		board: NewBoard(),
	}
}
func (g *Game) Play() {
	for {
		_, rune := g.turn.GetTurn()
		g.board.PrintBoard()
		fmt.Printf("Player %c turn\n", rune)
		row, col, err := g.readInput()
		if err != nil {
			fmt.Println("Error leyendo la entrada:", err)
			continue
		}

		g.board.PlaceToken(NewToken(rune, row, col))

		if g.board.CheckWin(rune) {
			g.board.PrintBoard()
			fmt.Printf("Player %c wins!\n", rune)
			os.Exit(0)
		}

		if g.board.FullBoard() {
			g.board.PrintBoard()
			fmt.Println("Draw!")
			os.Exit(0)
		}

		g.turn.TurnChange()

	}

}

func (g *Game) readInput() (int, int, error) {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error leyendo la entrada:", err)
			continue
		}

		parts := strings.Split(input, ".")
		if len(parts) != 2 {
			fmt.Println("Entrada invalida. Por favor, ingresa numero de fila, numero de columna valido.")
			continue
		}

		row, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		fmt.Println("row:", parts[0])
		if err != nil {
			fmt.Println("Entrada invalida. Por favor, ingresa numero de filas valido.")
			continue
		}

		col, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		fmt.Println("col:", parts[1])
		if err != nil {
			fmt.Println("Entrada invalida. Por favor, ingresa numero de columnas valido.", err)
			continue
		}

		if row < 0 || row > 2 || col < 0 || col > 2 {
			fmt.Println("Entrada invalida. Por favor, ingresa numero de fila y numero de columna valido dentro de rango.")
			continue
		}

		return row, col, nil
	}
}
