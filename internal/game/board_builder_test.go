package game_test

import (
	"github.com/thotluna/ttt/internal/game"
)

type boardBuilder struct {
	board *game.Board
}

// newBoardBuilder crea un nuevo constructor de tablero para tests
func newBoardBuilder(board *game.Board) *boardBuilder {
	return &boardBuilder{
		board: board,
	}
}

// WithSymbol coloca un símbolo en la posición especificada
func (b *boardBuilder) WithSymbol(symbol game.Symbol, row, col int) *boardBuilder {
	coor, _ := game.NewCoordinate(row, col)
	_ = b.board.PlaceToken(symbol, nil, &coor) // Ignoramos el error a propósito en tests
	return b
}

// WithRow configura una fila completa del tablero
func (b *boardBuilder) WithRow(row int, symbols ...game.Symbol) *boardBuilder {
	for col, symbol := range symbols {
		if col >= game.BoardSize {
			break
		}
		b.WithSymbol(symbol, row, col)
	}
	return b
}

// Build devuelve el tablero construido
func (b *boardBuilder) Build() *game.Board {
	return b.board
}
