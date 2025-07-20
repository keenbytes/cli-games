package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/keenbytes/cli-games/pkg/lettersnake"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type gamePane struct {
	game  *lettersnake.Game
	pane  *termui.Pane
	speed int
}

func (w *gamePane) Render(pane *termui.Pane) {
	w.drawInitial(pane)
}

func (w *gamePane) Iterate(pane *termui.Pane) {
	w.drawInitial(pane)
}

func (w *gamePane) HasBackend() bool {
	return true
}

func (w *gamePane) Backend(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(w.speed) * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if w.game.State() != lettersnake.GameOn {
				continue
			}

			if w.game.NumUsedWords() == 0 {
				clearPane(w.pane)
			}

			event := w.game.Iterate()
			switch event {
			case lettersnake.AteItself, lettersnake.EdgeHit, lettersnake.AllWordsUsed:
				w.drawInitial(w.pane)
			default:
				letters := w.game.Letters()
				for positionX, mapY := range *letters {
					for positionY, letter := range mapY {
						w.pane.Write(positionX, positionY, w.wrapInRandomColour(string(letter)))
					}
				}

				w.drawSnake()
			}
		}
	}
}

//nolint:mnd
func (w *gamePane) drawInitial(pane *termui.Pane) {
	if !w.game.IsPlayAreaSizeSet() {
		w.game.SetPlayAreaSize(w.pane.CanvasWidth(), w.pane.CanvasHeight())
	}

	state := w.game.State()
	switch state {
	case lettersnake.NotStarted:
		pane.Write(1, 0, "Instructions")
		pane.Write(1, 1, "------------")
		pane.Write(1, 2, "Do you know Snake? Here only")
		pane.Write(1, 3, "properly written words disappear.")
		pane.Write(1, 4, "Use ASDW to steer the snake.")
		pane.Write(1, 6, "Can you eat all the letters")
		pane.Write(1, 7, "in a correct order?")
		pane.Write(1, 9, "Press G to start the game.")
		pane.Write(1, 10, "Press T at any time to quit.")
		pane.Write(1, 12, "Selected game")
		pane.Write(1, 13, "-------------")
		pane.Write(1, 14, w.game.WordListTitle())

		return
	case lettersnake.GameOver:
		pane.Write(2, 0, "** Game over! **")

		return
	default:
	}
}

func (w *gamePane) drawSnake() {
	snake := w.game.Snake()
	for i := range snake {
		w.pane.Write(snake[i].PositionX, snake[i].PositionY, w.getSnakeSegment(i))
	}

	remove := w.game.Tail()
	if remove != nil {
		w.pane.Write(remove.PositionX, remove.PositionY, " ")
	}
}

//nolint:mnd
func (w *gamePane) getSnakeSegment(index int) string {
	// 125-159
	colour := 125 + index

	char := "▓"
	if index > 0 {
		char = "▒"
	}

	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", colour, char)
}

func (w *gamePane) wrapInRandomColour(text string) string {
	colours := []string{"\033[1;93m", "\033[1;92m", "\033[1;95m", "\033[1;96m"}
	reset := "\033[0m"

	//nolint:gosec
	return colours[rand.Intn(len(colours))] + text + reset
}
