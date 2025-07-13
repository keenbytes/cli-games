package main

import (
	"context"

	"github.com/keenbytes/cli-games/pkg/ortotris"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type infoPane struct {
	g *ortotris.Game
}

func (w *infoPane) Render(pane *termui.Pane) {
	pane.Write(0, 0, " ")
}

func (w infoPane) Iterate(pane *termui.Pane) {
}

func (w infoPane) HasBackend() bool {
	return false
}

func (w *infoPane) Backend(ctx context.Context) {
}
