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
		var row, col int
		var err error
		for {
			row, col, err = g.readInput()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			break
		}

		token := NewToken(rune, row, col)
		if err := g.board.PlaceToken(token); err != nil {
			if Is(err, ErrPositionOccupied) || Is(err, ErrOutOfBounds) {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Unexpected error:", err)
			}
			continue
		}

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
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your move (row.col, e.g. 1.2): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, NewGameError(ErrInvalidInput, "failed to read input")
	}

	input = strings.TrimSpace(input)
	parts := strings.Split(input, ".")
	if len(parts) != 2 {
		return 0, 0, NewGameError(ErrInvalidInput, "invalid format. Please use 'row.col' (e.g., '1.2')")
	}

	row, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, NewGameError(ErrInvalidInput, "row must be a number between 0 and 2")
	}

	col, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, NewGameError(ErrInvalidInput, "column must be a number between 0 and 2")
	}

	if row < 0 || row > 2 || col < 0 || col > 2 {
		return 0, 0, NewGameError(ErrInvalidInput, fmt.Sprintf("position (%d,%d) is out of bounds", row, col))
	}

	return row, col, nil

}
