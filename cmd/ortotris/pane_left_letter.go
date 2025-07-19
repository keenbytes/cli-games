package main

import (
	"context"

	"github.com/keenbytes/cli-games/pkg/ortotris"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type leftLetterPane struct {
	game *ortotris.Game
}

func (p *leftLetterPane) Render(pane *termui.Pane) {
	pane.Write(0, 0, "   <-   ")
	pane.Write(0, 1, "    "+p.game.LeftLetter()+"   ")
}

func (p *leftLetterPane) Iterate(pane *termui.Pane) {
}

func (p *leftLetterPane) HasBackend() bool {
	return false
}

func (p *leftLetterPane) Backend(ctx context.Context) {
}
