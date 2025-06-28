package game

import (
"fmt"
"strconv"
"strings"

"github.com/thotluna/ttt/internal/view"
)

type Player struct {
symbol rune
io     view.IO
board  *Board
}

func NewPlayer(symbol rune, io view.IO, board *Board) *Player {
return &Player{
symbol: symbol,
io:     io,
board:  board,
}
}

func (p *Player) Put() bool {
for {
row, col, err := p.readInput()
if err != nil {
p.io.PrintLine(err.Error())
continue
}

err = p.board.PlaceToken(p.symbol, Coordinate{row, col})
if err != nil {
p.io.PrintLine(err.Error())
continue
}

return p.CheckWin()
}
}

func (p *Player) CheckWin() bool {
tokens := p.board.GetTokenBy(p.symbol)
if len(tokens) < 3 {
return false
}

dir1 := tokens[0].Direction(tokens[1])
dir2 := tokens[1].Direction(tokens[2])
isWin := dir1 == dir2 && dir1 != None

if isWin {
p.io.PrintLine(fmt.Sprintf(MsgPlayerWins, p.symbol))
}

return isWin
}

func (p *Player) readInput() (int, int, error) {
input := p.io.ReadInput()
input = strings.TrimSpace(input)
parts := strings.Split(input, ".")
if len(parts) != 2 {
return 0, 0, NewGameError(ErrInvalidInput, MsgInvalidFormat)
}

row, err := strconv.Atoi(strings.TrimSpace(parts[0]))
if err != nil {
return 0, 0, NewGameError(ErrInvalidInput, MsgRowMustBeNumber)
}

col, err := strconv.Atoi(strings.TrimSpace(parts[1]))
if err != nil {
return 0, 0, NewGameError(ErrInvalidInput, MsgColMustBeNumber)
}

if row < 0 || row > 2 || col < 0 || col > 2 {
return 0, 0, NewGameError(ErrOutOfBounds, MsgOutOfBounds)
}

return row, col, nil
}
