package termui

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"gopkg.pl/mikogs/termui/pkg/term"
)

type TermUI struct {
	stdout *os.File
	stderr *os.File
	width int
	height int
	pane *Pane
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
func (t *TermUI) Run(stdout *os.File, stderr *os.File) int {
	t.stdout = stdout
	t.stderr = stderr
	term.InitTTY()
	term.Clear(t.stdout)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan struct{}, 1)

	go t.loop(sigs, done)
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

// loop is the main program loop
func (t *TermUI) loop(sigs <-chan os.Signal, done chan<- struct{}) {
	ticker := time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case <-sigs:
			t.exit()
			done<-struct{}{}
		case <-ticker.C:
			sizeChanged := t.refreshSize()
			if sizeChanged {
				term.Clear(t.stdout)
				t.pane.render()
			}
			t.pane.iterate()
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
