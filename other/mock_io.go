package other_test

import (
	"strings"

	"github.com/thotluna/ttt/view"
)

type MockIO struct {
	inputQueue []string
	output     []string
}

func NewMockIO(inputs ...string) *MockIO {
	return &MockIO{
		inputQueue: inputs,
		output:     []string{},
	}
}

var _ view.IO = (*MockIO)(nil)

func (m *MockIO) ReadInput() string {
	if len(m.inputQueue) == 0 {
		return ""
	}
	input := m.inputQueue[0]
	m.inputQueue = m.inputQueue[1:]
	return input
}

func (m *MockIO) PrintLine(line string) {
	m.output = append(m.output, line)
}

func (m *MockIO) PrintMessage(message string) {
	m.output = append(m.output, message)
}

func (m *MockIO) PrintBoard(board [3][3]rune) {
	var rows []string
	for _, row := range board {
		var rowStr strings.Builder
		for _, cell := range row {
			if cell == 0 {
				rowStr.WriteString(".")
			} else {
				rowStr.WriteRune(cell)
			}
		}
		rows = append(rows, rowStr.String())
	}
	m.output = append(m.output, strings.Join(rows, "\n"))
}

func (m *MockIO) PrintWin(player rune) {
	m.output = append(m.output, "Player "+string(player)+" wins!")
}

func (m *MockIO) PrintDraw() {
	m.output = append(m.output, "It's a draw!")
}

func (m *MockIO) GetOutput() []string {
	return m.output
}

func (m *MockIO) ContainsOutput(text string) bool {
	for _, output := range m.output {
		if strings.Contains(output, text) {
			return true
		}
	}
	return false
}
