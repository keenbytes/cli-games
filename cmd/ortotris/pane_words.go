package main

import (
	"context"
	"time"

	"github.com/keenbytes/cli-games/pkg/ortotris"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type wordsPane struct {
	g              *ortotris.Game
	previousStatus int
	pane           *termui.Pane
	speed          int
}

func (w *wordsPane) Render(pane *termui.Pane) {
	// Update canvas height so that we now how many lines are available
	w.g.SetAvailableLines(pane.CanvasHeight())
	w.drawInitial(pane)
}

func (w wordsPane) Iterate(pane *termui.Pane) {
	w.drawInitial(pane)
}

func (w wordsPane) HasBackend() bool {
	return true
}

func (w *wordsPane) Backend(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(w.speed) * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if w.g.State() != ortotris.GameOn {
				continue
			}

			if w.g.NumUsedWords() == 0 {
				clearPane(w.pane)
			}

			event := w.g.Iterate()
			switch event {
			case ortotris.TopReached, ortotris.AllWordsUsed:
				w.drawInitial(w.pane)
			case ortotris.CorrectAnswer:
				line := w.g.PrevCurrentLine()
				if line > 0 {
					clearPaneLine(w.pane, line)
				}

				line = w.g.CurrentLine()
				if line > 0 {
					clearPaneLine(w.pane, line)
				}

				clearPaneLine(w.pane, line+1)
			case ortotris.WrongAnswer, ortotris.ContinueGame:
				line := w.g.PrevCurrentLine()
				if line > 0 {
					clearPaneLine(w.pane, line)
				}

				line = w.g.CurrentLine()
				if line > 0 {
					clearPaneLine(w.pane, line)
				}

				w.writeWord(w.pane, w.g.CurrentWord(), line+1)
			default:
			}
		}
	}
}

func (w *wordsPane) drawInitial(pane *termui.Pane) {
	state := w.g.State()
	switch state {
	case ortotris.NotStarted:
		pane.Write(1, 0, "Instructions")
		pane.Write(1, 1, "------------")
		pane.Write(1, 2, "Do you know Tetris? Here only")
		pane.Write(1, 3, "properly written words disappear.")
		pane.Write(1, 4, "Use Arrows to choose the missing")
		pane.Write(1, 5, "letter.")
		pane.Write(1, 6, "Can you get all the words")
		pane.Write(1, 7, "correctly?")
		pane.Write(1, 9, "Press S to start the game.")
		pane.Write(1, 10, "Press X at any time to quit.")
		pane.Write(1, 12, "Selected game")
		pane.Write(1, 13, "-------------")
		pane.Write(1, 14, w.g.Title())

		return
	case ortotris.GameOver:
		pane.Write(2, 0, "** Game over! **")

		return
	default:
	}
}

func (w *wordsPane) writeWord(pane *termui.Pane, word string, l int) {
	pane.Write((pane.CanvasWidth()-len(word))/2, l, word)
}
