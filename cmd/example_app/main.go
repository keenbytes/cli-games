package main

import (
	"os"

	t "gopkg.pl/mikogs/termui"
)

func main() {
	tui := t.NewTermUI()
	main := tui.Pane()

	/*
	+----------+---------------------------------+
	|+--------+|+---------------+---------------+|
	||title   |||+-------------+|+-------------+||
	||        ||||contentTop   |||itemHeader   |||
  |+--------+|||             |||             |||
	||menu    |||+-------------+|+-------------+||
	||        ||||contentBottom|||itemContent  |||
	||        ||||             |||             |||
	||        |||+-------------+|+-------------+||
	|+--------+|+---------------+---------------+|
	+----------+---------------------------------+
	*/

	_left, _right := main.Split(t.Vertically, t.Left, 10, t.Char)
	_rightCol1, _rightCol2 := _right.Split(t.Vertically, t.Left, 70, t.Percent)

	title, menu := _left.Split(t.Horizontally, t.Top, 20, t.Char)
	contentTop, contentBottom := _rightCol1.Split(t.Horizontally, t.Top, 50, t.Percent)
	itemHeader, itemContent := _rightCol2.Split(t.Horizontally, t.Top, 5, t.Char)

	tui.SetFrame(&t.NoFrame{}, contentTop, contentBottom, itemContent)
	tui.SetFrame(&t.FrameRight{}, title, menu)
	tui.SetFrame(&t.Frame{}, itemHeader)

	tui.Run(os.Stdout, os.Stderr)
}
