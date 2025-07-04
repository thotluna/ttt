package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var _ IO = (*IOTerminal)(nil)

type IOTerminal struct {
}

func NewIOTerminal() IOTerminal {
	return IOTerminal{}
}

func (i *IOTerminal) ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (i *IOTerminal) PrintLine(line string) {
	fmt.Println(line)
}

func (i *IOTerminal) PrintMessage(message string) {
	fmt.Print(message + " ")
}

func (i *IOTerminal) PrintBoard(board [3][3]rune) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
}
