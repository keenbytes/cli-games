package termui

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.pl/mikogs/termui/pkg/term"
)

type TermUI struct {
	stdout        *os.File
	stderr        *os.File
	width         int
	height        int
	pane          *Pane
	iterablePanes []*Pane
	backendPanes  []*Pane
}

func NewTermUI() *TermUI {
	t := &TermUI{
		pane: &Pane{},
	}

	t.pane.ui = t
	return t
}

// GetPane returns initial terminal pane
func (t *TermUI) Pane() *Pane {
	return t.pane
}

// Run clears the terminal and starts program's main loop
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
		go pane.Widget.Backend(ctx)
	}

	done := make(chan struct{}, 1)
	defer close(done)
	go t.loop(ctx, done, backendCancelFuncs)
	<-done

	return 0
}

// Write prints out on the terminal window at a specified position
func (t *TermUI) Write(x int, y int, s string) {
	fmt.Fprintf(t.stdout, "\u001b[1000A\u001b[1000D")
	if x > 0 {
		fmt.Fprintf(t.stdout, "\u001b["+strconv.Itoa(x)+"C")
	}
	if y > 0 {
		fmt.Fprintf(t.stdout, "\u001b["+strconv.Itoa(y)+"B")
	}
	fmt.Fprint(t.stdout, s)
}

// RefreshIterablePanes loops through all the panes and gets the ones that are not a split
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

// loop is the main program loop
func (t *TermUI) loop(ctx context.Context, done chan<- struct{}, backendCancelFuncs []context.CancelFunc) {
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

// refreshSize gets terminal size and caches it
func (t *TermUI) refreshSize() bool {
	w, h, err := term.GetSize()
	if err != nil {
		return false
	}
	if t.width != w || t.height != h {
		t.width = w
		t.height = h
		t.pane.setWidth(w)
		t.pane.setHeight(h)
		return true
	}
	return false
}
