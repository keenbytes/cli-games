package main

import (
	"context"
	"os"
	"sync"

	"github.com/keenbytes/cli-games/pkg/lettersnake"
	"github.com/keenbytes/cli-games/pkg/termui"
)

type gameInterface struct {
	tui       *termui.TermUI
	playarea  *termui.Pane
	score     *termui.Pane
	leftWord  *termui.Pane
	rightWord *termui.Pane
	game      *lettersnake.Game
}

//nolint:mnd
func newGameInterface(game *lettersnake.Game, speed int) *gameInterface {
	gui := &gameInterface{}

	gui.game = game
	gui.tui = termui.NewTermUI()
	mainPane := gui.tui.Pane()

	paneScore, _bottom := mainPane.Split(termui.Horizontally, termui.LeftPane, 3, termui.Char)
	paneGame, _bottomBottom := _bottom.Split(termui.Horizontally, termui.RightPane, 3, termui.Char)

	paneLeftWord, paneRightWord := _bottomBottom.Split(
		termui.Vertically,
		termui.RightPane,
		50,
		termui.Percent,
	)

	paneScore.Widget = &scorePane{game: game}
	paneLeftWord.Widget = &leftWordPane{game: game}
	paneRightWord.Widget = &rightWordPane{game: game}
	paneGame.Widget = &gamePane{game: game, pane: paneGame, speed: speed}

	gui.playarea = paneGame
	gui.score = paneScore
	gui.leftWord = paneLeftWord
	gui.rightWord = paneRightWord

	gui.tui.SetFrame(&termui.Frame{}, paneGame, paneScore, paneLeftWord, paneRightWord)

	return gui
}

//nolint:gocognit,mnd
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

			if string(input) == "t" {
				cancel()

				break
			}

			if string(input) == "g" {
				if gui.game.State() != lettersnake.GameOn {
					gui.game.StartGame()
				}

				continue
			}

			if string(input) == "a" {
				if gui.game.Direction() != lettersnake.MovingRight {
					gui.game.SetDirection(lettersnake.MovingLeft)
				}

				continue
			}

			if string(input) == "d" {
				if gui.game.Direction() != lettersnake.MovingLeft {
					gui.game.SetDirection(lettersnake.MovingRight)
				}

				continue
			}

			if string(input) == "s" {
				if gui.game.Direction() != lettersnake.MovingUp {
					gui.game.SetDirection(lettersnake.MovingDown)
				}

				continue
			}

			if string(input) == "w" {
				if gui.game.Direction() != lettersnake.MovingDown {
					gui.game.SetDirection(lettersnake.MovingUp)
				}

				continue
			}
		}

		waitGroup.Done()
	}()

	waitGroup.Wait()
}
