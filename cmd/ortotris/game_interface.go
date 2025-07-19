package main

import (
	"context"
	"os"
	"sync"

	"github.com/keenbytes/cli-games/pkg/ortotris"
	"github.com/keenbytes/cli-games/pkg/termui"
)

const (
	letterColumnWidth = 10
	bottomRowHeight   = 4
)

type gameInterface struct {
	tui         *termui.TermUI
	words       *termui.Pane
	leftLetter  *termui.Pane
	rightLetter *termui.Pane
	score       *termui.Pane
	info        *termui.Pane
	game        *ortotris.Game
}

func newGameInterface(game *ortotris.Game, speed int) *gameInterface {
	gui := &gameInterface{}

	gui.game = game
	gui.tui = termui.NewTermUI()
	mainPane := gui.tui.Pane()

	_left, _middleAndRight := mainPane.Split(
		termui.Vertically,
		termui.LeftPane,
		letterColumnWidth,
		termui.Char,
	)
	paneWords, _right := _middleAndRight.Split(
		termui.Vertically,
		termui.RightPane,
		letterColumnWidth,
		termui.Char,
	)
	paneInfo, paneLeftLetter := _left.Split(
		termui.Horizontally,
		termui.RightPane,
		bottomRowHeight,
		termui.Char,
	)
	paneScore, paneRightLetter := _right.Split(
		termui.Horizontally,
		termui.RightPane,
		bottomRowHeight,
		termui.Char,
	)

	paneInfo.Widget = &infoPane{game: game}
	paneLeftLetter.Widget = &leftLetterPane{game: game}
	paneRightLetter.Widget = &rightLetterPane{game: game}
	paneScore.Widget = &scorePane{game: game}
	paneWords.Widget = &wordsPane{game: game, pane: paneWords, speed: speed}

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
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	stopStdio := false

	go func() {
		gui.tui.Run(ctx, os.Stdout, os.Stderr)
		waitGroup.Done()

		stopStdio = true
	}()

	go func() {
		input := make([]byte, 1)

		for !stopStdio {
			_, _ = os.Stdin.Read(input)
			// key press code here
			if string(input) == "t" {
				cancel()

				break
			}

			if string(input) == "g" {
				if gui.game.State() != ortotris.GameOn {
					gui.game.StartGame()
				}

				continue
			}

			if string(input) == "a" {
				gui.game.ChooseLeftLetter()

				continue
			}
			// right arrow pressed
			if string(input) == "d" {
				gui.game.ChooseRightLetter()

				continue
			}
			// down arrow pressed
			if string(input) == "s" {
				gui.game.SetNextLineToLast()
			}

			// increase speed
			if string(input) == "i" {
			}

			// decrease speed
			if string(input) == "k" {
			}
		}

		waitGroup.Done()
	}()

	waitGroup.Wait()
}
