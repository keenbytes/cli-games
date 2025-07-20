package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/keenbytes/cli-games/pkg/ortotris"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type scorePane struct {
	game *ortotris.Game
}

func (p *scorePane) Render(_ *termui.Pane) {
}

//nolint:mnd
func (p *scorePane) Iterate(pane *termui.Pane) {
	pane.Write(0, 0, "Correct:")
	pane.Write(1, 1, fmt.Sprintf("%d/%d", p.game.NumCorrectGuesses(), p.game.NumUsedWords()))
	pane.Write(0, 3, "Total:")
	pane.Write(1, 4, strconv.Itoa(p.game.NumWordList()))
}

func (p *scorePane) HasBackend() bool {
	return false
}

func (p *scorePane) Backend(_ context.Context) {
}
