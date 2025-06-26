package main

import (
	"os"

	"github.com/thotluna/ttt/game"
	"github.com/thotluna/ttt/view"
)

func main() {
	io := view.NewIOTerminal()
	g := game.NewGame(&io)
	g.Play()

	os.Exit(0)
}
