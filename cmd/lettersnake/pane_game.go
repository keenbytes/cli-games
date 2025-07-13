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
	g              *lettersnake.Game
	pane           *termui.Pane
	speed          int
}

func (w *gamePane) Render(pane *termui.Pane) {
	w.drawInitial(pane)
}

func (w gamePane) Iterate(pane *termui.Pane) {
	w.drawInitial(pane)
}

func (w gamePane) HasBackend() bool {
	return true
}

func (w *gamePane) Backend(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(w.speed) * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if w.g.State() != lettersnake.GameOn {
				continue
			}

			if w.g.NumUsedWords() == 0 {
				clearPane(w.pane)
			}

			event := w.g.Iterate()
			switch event {
			case lettersnake.AteItself, lettersnake.EdgeHit, lettersnake.AllWordsUsed:
				w.drawInitial(w.pane)
			default:
				letters := w.g.Letters()
				for i := 0; i < len(letters); i++ {
					w.pane.Write(letters[i].X, letters[i].Y, w.wrapInRandomColour(letters[i].L))
				}
				w.drawSnake()
			}
		}
	}
}

func (w *gamePane) drawInitial(pane *termui.Pane) {
	if !w.g.SizeSet() {
		w.g.SetSize(w.pane.CanvasWidth(), w.pane.CanvasHeight())
	}
	
	state := w.g.State()
	switch state {
	case lettersnake.NotStarted:
		pane.Write(1, 0, "Instructions")
		pane.Write(1, 1, "------------")
		pane.Write(1, 2, "Do you know Snake? Here only")
		pane.Write(1, 3, "properly written words disappear.")
		pane.Write(1, 4, "Use Arrows to steer the snake.")
		pane.Write(1, 6, "Can you eat all the letters")
		pane.Write(1, 7, "in a correct order?")
		pane.Write(1, 9, "Press S to start the game.")
		pane.Write(1, 10, "Press X at any time to quit.")
		pane.Write(1, 12, "Selected game")
		pane.Write(1, 13, "-------------")
		pane.Write(1, 14, w.g.Title())
		return
	case lettersnake.GameOver:
		pane.Write(2, 0, "** Game over! **")
		return
	default:
	}
}

func (w *gamePane) drawSnake() {
	snake := w.g.Snake()
	for i := 0; i < len(snake); i++ {
		w.pane.Write(snake[i].X, snake[i].Y, w.getSnakeSegment(i))
	}
	remove := w.g.Remove()
	if remove != nil {
		w.pane.Write(remove.X, remove.Y, " ")
	}
}

func (w *gamePane) getSnakeSegment(i int) string {
	// 125-159
	c := 125 + i
	s := "▓"
	if i > 0 {
		s = "▒"
	}
	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", c, s)
}

func (w *gamePane) wrapInRandomColour(s string) string {
	colours := []string{"\033[1;93m", "\033[1;92m", "\033[1;95m", "\033[1;96m"}
	reset := "\033[0m"
	return colours[rand.Intn(len(colours))] + s + reset
}
