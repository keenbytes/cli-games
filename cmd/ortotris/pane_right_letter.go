package main

import (
	"context"

	"github.com/keenbytes/cli-games/pkg/ortotris"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type rightLetterPane struct {
	game *ortotris.Game
}

func (p *rightLetterPane) Render(pane *termui.Pane) {
	pane.Write(0, 0, "   ->   ")
	pane.Write(0, 1, "   "+p.game.RightLetter()+"    ")
}

func (p *rightLetterPane) Iterate(_ *termui.Pane) {
}

func (p *rightLetterPane) HasBackend() bool {
	return false
}

func (p *rightLetterPane) Backend(_ context.Context) {
}
