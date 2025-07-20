package main

import (
	"context"
	"fmt"

	"github.com/keenbytes/cli-games/pkg/lettersnake"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type scorePane struct {
	game *lettersnake.Game
}

func (w *scorePane) Render(_ *termui.Pane) {
}

func (w *scorePane) Iterate(pane *termui.Pane) {
	pane.Write(
		1,
		0,
		fmt.Sprintf("Correct: %d/%d", w.game.NumCorrectGuesses(), w.game.NumUsedWords()),
	)
}

func (w *scorePane) HasBackend() bool {
	return false
}

func (w *scorePane) Backend(_ context.Context) {
}
