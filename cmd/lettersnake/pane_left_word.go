package main

import (
	"context"
	"strings"

	"github.com/keenbytes/cli-games/pkg/lettersnake"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type leftWordPane struct {
	game *lettersnake.Game
}

func (w *leftWordPane) Render(_ *termui.Pane) {
}

//nolint:mnd
func (w *leftWordPane) Iterate(pane *termui.Pane) {
	trim := 20 - len(w.game.CurrentTranslation())
	pane.Write(1, 0, w.game.CurrentTranslation()+strings.Repeat(" ", trim))
}

func (w *leftWordPane) HasBackend() bool {
	return false
}

func (w *leftWordPane) Backend(_ context.Context) {
}
