// Package termui is designed to simplify output to a terminal window by allowing the specification of panes with
// static or dynamic content.
package termui

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/keenbytes/cli-games/pkg/term"
)

// TermUI represents the UI in the terminal.
type TermUI struct {
	stdout        *os.File
	stderr        *os.File
	width         int
	height        int
	pane          *Pane
	iterablePanes []*Pane
	backendPanes  []*Pane
	mutex         sync.Mutex
}

// NewTermUI returns new TermUI instance.
func NewTermUI() *TermUI {
	termUI := &TermUI{
		pane: &Pane{},
	}

	termUI.pane.ui = termUI

	return termUI
}

// Pane returns initial terminal pane.
func (t *TermUI) Pane() *Pane {
	return t.pane
}

// Run clears the terminal and starts program's main loop.
func (t *TermUI) Run(ctx context.Context, stdout *os.File, stderr *os.File) int {
	t.stdout = stdout
	t.stderr = stderr

	term.InitTTY()
	term.Clear(t.stdout)

	t.getIterablePanes(nil)
	backendCancelFuncs := make([]context.CancelFunc, 0, len(t.backendPanes))

	for _, pane := range t.backendPanes {
		ctx, cancel := context.WithCancel(context.Background())
		backendCancelFuncs = append(backendCancelFuncs, cancel)

		//nolint:contextcheck
		go pane.Widget.Backend(ctx)
	}

	done := make(chan struct{}, 1)
	go t.loop(ctx, done, backendCancelFuncs)

	<-done

	return 0
}

// Write prints out on the terminal window at a specified position.
//
//nolint:errcheck
func (t *TermUI) Write(positionX int, positionY int, str string) {
	t.mutex.Lock()
	fmt.Fprintf(t.stdout, "\u001b[1000A\u001b[1000D")

	if positionX > 0 {
		fmt.Fprintf(t.stdout, "\u001b[%sC", strconv.Itoa(positionX))
	}

	if positionY > 0 {
		fmt.Fprintf(t.stdout, "\u001b[%sB", strconv.Itoa(positionY))
	}

	fmt.Fprint(t.stdout, str)
	t.mutex.Unlock()
}

// RefreshIterablePanes loops through all the panes and gets the ones that are not a split.
func (t *TermUI) getIterablePanes(pane *Pane) {
	if pane == nil {
		t.iterablePanes = make([]*Pane, 0)
		t.backendPanes = make([]*Pane, 0)
		t.getIterablePanes(t.pane)

		return
	}

	switch pane.splitType {
	case Horizontally, Vertically:
		t.getIterablePanes(pane.panes[0])
		t.getIterablePanes(pane.panes[1])

	default:
		t.iterablePanes = append(t.iterablePanes, pane)
		if pane.Widget != nil && pane.Widget.HasBackend() {
			t.backendPanes = append(t.backendPanes, pane)
		}
	}
}

// loop is the main program loop.
func (t *TermUI) loop(
	ctx context.Context,
	done chan<- struct{},
	backendCancelFuncs []context.CancelFunc,
) {
	//nolint:mnd
	ticker := time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			for _, fn := range backendCancelFuncs {
				fn()
			}

			t.exit()

			done <- struct{}{}
		case <-ticker.C:
			sizeChanged := t.refreshSize()
			if sizeChanged {
				term.Clear(t.stdout)
				t.pane.render()
			}

			if len(t.iterablePanes) > 0 {
				for _, pane := range t.iterablePanes {
					pane.iterate()
				}
			}
		}
	}
}

func (t *TermUI) exit() {
	term.Clear(t.stdout)
}

// refreshSize gets terminal size and caches it.
func (t *TermUI) refreshSize() bool {
	width, height, err := term.GetSize()
	if err != nil {
		return false
	}

	if t.width != width || t.height != height {
		t.width = width
		t.height = height
		t.pane.setWidth(width)
		t.pane.setHeight(height)

		return true
	}

	return false
}
