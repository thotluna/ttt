package game

import "github.com/thotluna/ttt/internal/view"

var (
	boardInterval = NewGameInterval()
)

const (
	BoardSize = 3
)

type Board struct {
	board [BoardSize][BoardSize]SymbolPlayerCurrent
	io    view.IO
}

func NewBoard(io view.IO) *Board {
	return &Board{
		io: io,
		board: [BoardSize][BoardSize]SymbolPlayerCurrent{
			{EmptyCell, EmptyCell, EmptyCell},
			{EmptyCell, EmptyCell, EmptyCell},
			{EmptyCell, EmptyCell, EmptyCell},
		},
	}
}

func (b *Board) GetBoard() [BoardSize][BoardSize]SymbolPlayerCurrent {
	return b.board
}

func (b *Board) SetBoard(board [BoardSize][BoardSize]SymbolPlayerCurrent) {
	b.board = board
}

func (b *Board) PlaceToken(symbol SymbolPlayerCurrent, coor Coordinate) error {

	if !b.isEmptyCell(coor) {
		return NewGameError(ErrPositionOccupied,
			FormatPositionTaken(coor.row, coor.col))
	}

	b.board[coor.row][coor.col] = symbol
	return nil
}

func (b *Board) GetTokenBy(symbol SymbolPlayerCurrent) []Coordinate {
	var coordinates []Coordinate
	for i := boardInterval.Min(); i <= boardInterval.Max(); i++ {
		for j := boardInterval.Min(); j <= boardInterval.Max(); j++ {
			coor, _ := NewCoordinate(i, j)
			if b.IsOccupiedCellBy(coor, symbol) {
				coordinates = append(coordinates, coor)
			}
		}
	}
	return coordinates
}

func (b *Board) FullBoard() bool {
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

func (b *Board) IsOccupiedCellBy(coor Coordinate, symbol SymbolPlayerCurrent) bool {
	return b.board[coor.row][coor.col] == symbol
}

func (b *Board) isEmptyCell(coor Coordinate) bool {
	return b.IsOccupiedCellBy(coor, EmptyCell)
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
