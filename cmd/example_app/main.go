package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

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
	_contentTop, _contentBottom := _rightCol1.Split(t.Horizontally, t.Top, 50, t.Percent)
	itemHeader, itemContent := _rightCol2.Split(t.Horizontally, t.Top, 5, t.Char)

	box1, box2 := _contentTop.Split(t.Vertically, t.Left, 50, t.Percent)
	box3, box4 := _contentBottom.Split(t.Vertically, t.Left, 50, t.Percent)

	tui.SetFrame(&t.NoFrame{}, itemContent)
	tui.SetFrame(&t.FrameRight{}, title, menu)
	tui.SetFrame(&t.Frame{}, itemHeader, box1, box2, box3, box4)

	box1.Widget = &t.WidgetTime{}
	box2.Widget = &t.WidgetBackend{}
	box3.Widget = &t.WidgetGraph{}

	ctx, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		tui.Run(ctx, os.Stdout, os.Stderr)
		wg.Done()
	}()
	go func() {
		ticker := time.NewTicker(1000 * time.Millisecond)
		timeout := time.NewTimer(time.Duration(3600 * time.Second))
		widgetTime := widgetTime{}
		for {
			select {
			case <-timeout.C:
				cancel()
				wg.Done()
			case <-sigs:
				cancel()
				wg.Done()
			case <-ticker.C:
				now := time.Now()
				itemHeader.Write(0, 0, now.Format("15:04:05"))
				widgetTime.iterate(itemContent)
			}
		}
	}()
	wg.Wait()
}
