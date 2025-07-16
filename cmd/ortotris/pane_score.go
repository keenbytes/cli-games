package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/keenbytes/cli-games/pkg/ortotris"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type scorePane struct {
	g *ortotris.Game
}

func (w *scorePane) Render(pane *termui.Pane) {
}

func (w scorePane) Iterate(pane *termui.Pane) {
	pane.Write(0, 0, "Correct:")
	pane.Write(1, 1, fmt.Sprintf("%d/%d", w.g.NumCorrectAnswers(), w.g.NumUsedWords()))
	pane.Write(0, 3, "Total:")
	pane.Write(1, 4, strconv.Itoa(w.g.NumAllWords()))
}

func (w scorePane) HasBackend() bool {
	return false
}

func (w *scorePane) Backend(ctx context.Context) {
}
