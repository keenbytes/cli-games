package main

import (
	"strings"

	"github.com/keenbytes/cli-games/pkg/termui"
)

func clearPane(pane *termui.Pane) {
	for y := 0; y < pane.CanvasHeight(); y++ {
		clearPaneLine(pane, y)
	}
}

func clearPaneLine(pane *termui.Pane, y int) {
	pane.Write(0, y, strings.Repeat(" ", pane.CanvasWidth()))
}
