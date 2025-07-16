package main

import (
	"context"
	"os"
	"sync"

	"github.com/keenbytes/cli-games/pkg/ortotris"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type gameInterface struct {
	tui         *termui.TermUI
	words       *termui.Pane
	leftLetter  *termui.Pane
	rightLetter *termui.Pane
	score       *termui.Pane
	info        *termui.Pane
	g           *ortotris.Game
}

func newGameInterface(g *ortotris.Game, speed int) *gameInterface {
	gui := &gameInterface{}

	gui.g = g
	gui.tui = termui.NewTermUI()
	mainPane := gui.tui.Pane()

	_left, _middleAndRight := mainPane.Split(termui.Vertically, termui.Left, 10, termui.Char)
	paneWords, _right := _middleAndRight.Split(termui.Vertically, termui.Right, 10, termui.Char)
	paneInfo, paneLeftLetter := _left.Split(termui.Horizontally, termui.Right, 4, termui.Char)
	paneScore, paneRightLetter := _right.Split(termui.Horizontally, termui.Right, 4, termui.Char)

	paneInfo.Widget = &infoPane{g: g}
	paneLeftLetter.Widget = &leftLetterPane{g: g}
	paneRightLetter.Widget = &rightLetterPane{g: g}
	paneScore.Widget = &scorePane{g: g}
	paneWords.Widget = &wordsPane{g: g, pane: paneWords, speed: speed}

	gui.words = paneWords
	gui.leftLetter = paneLeftLetter
	gui.rightLetter = paneRightLetter
	gui.score = paneScore
	gui.info = paneInfo

	gui.tui.SetFrame(
		&termui.Frame{},
		paneWords,
		paneLeftLetter,
		paneRightLetter,
		paneScore,
		paneInfo,
	)

	return gui
}

func (gui *gameInterface) run(ctx context.Context, cancel func()) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	stopStdio := false

	go func() {
		gui.tui.Run(ctx, os.Stdout, os.Stderr)
		wg.Done()

		stopStdio = true
	}()

	go func() {
		b := make([]byte, 1)

		for !stopStdio {

			os.Stdin.Read(b)
			// key press code here
			if string(b) == "x" {
				cancel()

				break
			}

			if string(b) == "s" {
				if gui.g.State() != ortotris.GameOn {
					gui.g.StartGame()
				}

				continue
			}
			// TODO: Keys should be handled differently, maybe in raw mode
			// left arrow pressed
			if string(b) == "D" {
				gui.g.ChooseLeftLetter()

				continue
			}
			// right arrow pressed
			if string(b) == "C" {
				gui.g.ChooseRightLetter()

				continue
			}
			// down arrow pressed
			if string(b) == "B" {
				gui.g.SetNextLineToLast()
			}
		}

		wg.Done()
	}()

	wg.Wait()
}
