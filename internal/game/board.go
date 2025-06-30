package game

import (
	"github.com/thotluna/ttt/internal/view"
)

var (
	boardInterval = NewGameInterval()
)

const (
	BoardSize = 3
)

type Board struct {
	board [BoardSize][BoardSize]Symbol
	io    view.IO
}

func NewBoard(io view.IO) *Board {
	return &Board{
		io: io,
		board: [BoardSize][BoardSize]Symbol{
			{EmptyCell, EmptyCell, EmptyCell},
			{EmptyCell, EmptyCell, EmptyCell},
			{EmptyCell, EmptyCell, EmptyCell},
		},
	}
}

func (b *Board) PlaceToken(symbol Symbol, origin, destination *Coordinate) error {
	if origin != nil {
		if b.isEmptyCell(*origin) {
			return NewGameError(ErrNoTokenAtPosition, MsgNoTokenAtPosition)
		}

		if !b.IsCellOccupiedBy(*origin, symbol) {
			return NewGameError(ErrNotYourToken, MsgNotYourToken)
		}

		if origin.Equals(*destination) {
			return NewGameError(ErrSamePosition, MsgSamePosition)
		}

		b.board[origin.row][origin.col] = EmptyCell
	}

	if !b.isEmptyCell(*destination) {
		return NewGameError(ErrPositionOccupied, MsgDestinationOccupied)
	}

	b.board[destination.row][destination.col] = symbol
	return nil
}

func (b *Board) GetTokenBy(symbol Symbol) []Coordinate {
	var coordinates []Coordinate
	for i := boardInterval.Min(); i <= boardInterval.Max(); i++ {
		for j := boardInterval.Min(); j <= boardInterval.Max(); j++ {
			coor, _ := NewCoordinate(i, j)
			if b.IsCellOccupiedBy(coor, symbol) {
				coordinates = append(coordinates, coor)
			}
		}
	}
	return coordinates
}

func (b *Board) IsFull() bool {
	for i := boardInterval.Min(); i <= boardInterval.Max(); i++ {
		for j := boardInterval.Min(); j <= boardInterval.Max(); j++ {
			coor, _ := NewCoordinate(i, j)
			if b.isEmptyCell(coor) {
				return false
			}
		}
	}
	return true
}

func (b *Board) IsCellOccupiedBy(coor Coordinate, symbol Symbol) bool {
	return b.board[coor.row][coor.col] == symbol
}

func (b *Board) isEmptyCell(coor Coordinate) bool {
	return b.board[coor.row][coor.col] == EmptyCell
}

func (b *Board) Print() {
	var runeBoard [BoardSize][BoardSize]rune
	for i, row := range b.board {
		for j, cell := range row {
			runeBoard[i][j] = rune(cell)
		}
	}
	b.io.PrintBoard(runeBoard)
}
