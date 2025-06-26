package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	turn  int
	board [3][3]rune
}

func NewGame() Game {
	return Game{
		turn: 0,
		board: [3][3]rune{
			{'-', '-', '-'},
			{'-', '-', '-'},
			{'-', '-', '-'},
		},
	}
}

func (g *Game) getTurn() rune {
	switch g.turn {
	case 0:
		return 'X'
	case 1:
		return 'O'
	default:
		return 'X'
	}
}

func (g *Game) Play() {
	for {
		turn := g.getTurn()
		g.printBoard()
		fmt.Printf("Player %c's turn\n", turn)
		row, col, err := g.readInput()
		if err != nil {
			fmt.Println("Error leyendo la entrada:", err)
			return
		}

		g.board[row][col] = turn

		if g.checkWin(turn) {
			g.printBoard()
			fmt.Printf("Player %c wins!\n", turn)
			os.Exit(0)
		}

		if g.fullBoard() {
			g.printBoard()
			fmt.Println("Draw!")
			os.Exit(0)
		}

		g.TurnChange()

	}

}

func (g *Game) fullBoard() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board[i][j] == '-' {
				return false
			}
		}
	}
	return true
}

func (g *Game) checkWin(turn rune) bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if g.board[i][0] == g.board[i][1] && g.board[i][1] == g.board[i][2] && g.board[i][0] == turn {
			return true
		}
	}

	// Check columns
	for i := 0; i < 3; i++ {
		if g.board[0][i] == g.board[1][i] && g.board[1][i] == g.board[2][i] && g.board[0][i] == turn {
			return true
		}
	}

	// Check diagonals
	if g.board[0][0] == g.board[1][1] && g.board[1][1] == g.board[2][2] && g.board[0][0] == turn {
		return true
	}
	if g.board[0][2] == g.board[1][1] && g.board[1][1] == g.board[2][0] && g.board[0][2] == turn {
		return true
	}

	return false
}

func (g *Game) TurnChange() {
	g.turn = (g.turn + 1) % 2
}

func (g *Game) printBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%c ", g.board[i][j])
		}
		fmt.Println()
	}
}

func (g *Game) readInput() (int, int, error) {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return 0, 0, err
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
