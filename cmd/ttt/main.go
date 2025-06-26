package main

import (
	"os"

	"github.com/thotluna/ttt/internal/game"
	"github.com/thotluna/ttt/internal/view"
)

func main() {
	io := view.NewIOTerminal()
	g := game.NewGame(&io)
	g.Play()

	os.Exit(0)
}
