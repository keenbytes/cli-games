package main

import (
	"context"
	"strings"

	"github.com/keenbytes/cli-games/pkg/lettersnake"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type rightWordPane struct {
	game *lettersnake.Game
}

func (w *rightWordPane) Render(_ *termui.Pane) {
}

//nolint:mnd
func (w *rightWordPane) Iterate(pane *termui.Pane) {
	trim := 20 - len(w.game.ConsumedLetters())
	pane.Write(1, 0, w.game.ConsumedLetters()+strings.Repeat(" ", trim))
}

func (w *rightWordPane) HasBackend() bool {
	return false
}

func (w *rightWordPane) Backend(_ context.Context) {
}
