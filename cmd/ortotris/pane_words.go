package main

import (
	"context"
	"time"

	"github.com/keenbytes/cli-games/pkg/ortotris"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type wordsPane struct {
	game  *ortotris.Game
	pane  *termui.Pane
	speed int
}

func (p *wordsPane) Render(pane *termui.Pane) {
	// update canvas height so that we now how many lines are available
	p.game.SetAvailableLines(pane.CanvasHeight())
	p.drawInitial(pane)
}

func (p *wordsPane) Iterate(pane *termui.Pane) {
	p.drawInitial(pane)
}

func (p *wordsPane) HasBackend() bool {
	return true
}

func (p *wordsPane) Backend(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(p.speed) * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if p.game.State() != ortotris.GameOn {
				continue
			}

			if p.game.NumUsedWords() == 0 {
				clearPane(p.pane)
			}

			event := p.game.Iterate()
			switch event {
			case ortotris.TopReached, ortotris.AllWordsUsed:
				p.drawInitial(p.pane)

			case ortotris.CorrectAnswer:
				line := p.game.CurrentLine()
				if line > 0 {
					clearPaneLine(p.pane, line)
				}

			case ortotris.ContinueGame, ortotris.JumpToLastLine:
				line := p.game.CurrentLine()
				if line > 1 {
					clearPaneLine(p.pane, p.game.PreviousLine())
				}

				p.writeWord(p.pane, p.game.CurrentGuess(), line)

			case ortotris.WrongAnswer:
				line := p.game.CurrentLine()
				p.writeWord(p.pane, p.game.CurrentGuess(), line)
			}
		}
	}
}

//nolint:mnd
func (p *wordsPane) drawInitial(pane *termui.Pane) {
	state := p.game.State()
	switch state {
	case ortotris.NotStarted:
		pane.Write(1, 0, "Instructions")
		pane.Write(1, 1, "------------")
		pane.Write(1, 2, "Do you know Tetris? Here only")
		pane.Write(1, 3, "properly written words disappear.")
		pane.Write(1, 4, "Use A and D to choose the missing")
		pane.Write(1, 5, "letter.")
		pane.Write(1, 6, "Can you get all the words")
		pane.Write(1, 7, "correctly?")
		pane.Write(1, 9, "Press G to start the game.")
		pane.Write(1, 10, "Press T at any time to quit.")
		pane.Write(1, 12, "Selected game")
		pane.Write(1, 13, "-------------")
		pane.Write(1, 14, p.game.WordListTitle())

		return
	case ortotris.GameOver:
		pane.Write(2, 0, "** Game over! **")

		return
	default:
	}
}

//nolint:mnd
func (p *wordsPane) writeWord(pane *termui.Pane, word string, line int) {
	pane.Write((pane.CanvasWidth()-len(word))/2, line, word)
}
