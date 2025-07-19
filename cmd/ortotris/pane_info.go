package main

import (
	"context"

	"github.com/keenbytes/cli-games/pkg/ortotris"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type infoPane struct {
	game *ortotris.Game
}

func (p *infoPane) Render(pane *termui.Pane) {
	pane.Write(0, 0, " ")
}

func (p *infoPane) Iterate(pane *termui.Pane) {
}

func (p *infoPane) HasBackend() bool {
	return false
}

func (p *infoPane) Backend(ctx context.Context) {
}
