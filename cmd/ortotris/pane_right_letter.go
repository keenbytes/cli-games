package main

import (
	"context"

	"github.com/keenbytes/cli-games/pkg/ortotris"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type rightLetterPane struct {
	g *ortotris.Game
}

func (w *rightLetterPane) Render(pane *termui.Pane) {
	pane.Write(0, 0, "   ->   ")
	pane.Write(0, 1, "   "+w.g.RightLetter()+"    ")
}

func (w rightLetterPane) Iterate(pane *termui.Pane) {
}

func (w rightLetterPane) HasBackend() bool {
	return false
}

func (w *rightLetterPane) Backend(ctx context.Context) {
}
