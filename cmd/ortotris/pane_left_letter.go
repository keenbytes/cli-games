package main

import (
	"context"

	"github.com/keenbytes/cli-games/pkg/ortotris"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type leftLetterPane struct {
	g *ortotris.Game
}

func (w *leftLetterPane) Render(pane *termui.Pane) {
	pane.Write(0, 0, "   <-   ")
	pane.Write(0, 1, "    "+w.g.LeftLetter()+"   ")
}

func (w leftLetterPane) Iterate(pane *termui.Pane) {
}

func (w leftLetterPane) HasBackend() bool {
	return false
}

func (w *leftLetterPane) Backend(ctx context.Context) {
}
