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
	game      *termui.Pane
	score     *termui.Pane
	leftWord  *termui.Pane
	rightWord *termui.Pane
	g         *lettersnake.Game
}

func newGameInterface(g *lettersnake.Game, speed int) *gameInterface {
	gui := &gameInterface{}

	gui.g = g
	gui.tui = termui.NewTermUI()
	mainPane := gui.tui.Pane()

	paneScore, _bottom := mainPane.Split(termui.Horizontally, termui.Left, 3, termui.Char)
	paneGame, _bottomBottom := _bottom.Split(termui.Horizontally, termui.Right, 3, termui.Char)

	paneLeftWord, paneRightWord := _bottomBottom.Split(
		termui.Vertically,
		termui.Right,
		50,
		termui.Percent,
	)

	paneScore.Widget = &scorePane{g: g}
	paneLeftWord.Widget = &leftWordPane{g: g}
	paneRightWord.Widget = &rightWordPane{g: g}
	paneGame.Widget = &gamePane{g: g, pane: paneGame, speed: speed}

	gui.game = paneGame
	gui.score = paneScore
	gui.leftWord = paneLeftWord
	gui.rightWord = paneRightWord

	gui.tui.SetFrame(&termui.Frame{}, paneGame, paneScore, paneLeftWord, paneRightWord)

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
				if gui.g.State() != lettersnake.GameOn {
					gui.g.StartGame()
				}

				continue
			}
			// TODO: Keys should be handled differently, maybe in raw mode
			// left arrow pressed
			if string(b) == "D" {
				if gui.g.Direction() != lettersnake.Right {
					gui.g.SetDirection(lettersnake.Left)
				}

				continue
			}
			// right arrow pressed
			if string(b) == "C" {
				if gui.g.Direction() != lettersnake.Left {
					gui.g.SetDirection(lettersnake.Right)
				}

				continue
			}
			// down arrow pressed
			if string(b) == "B" {
				if gui.g.Direction() != lettersnake.Up {
					gui.g.SetDirection(lettersnake.Down)
				}

				continue
			}
			// up arrow pressed
			if string(b) == "A" {
				if gui.g.Direction() != lettersnake.Down {
					gui.g.SetDirection(lettersnake.Up)
				}

				continue
			}
		}

		wg.Done()
	}()

	wg.Wait()
}
