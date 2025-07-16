package termui

import (
	"context"
	"math/rand"
	"strconv"
	"time"
)

// TODO:.
type WidgetGraph struct {
	cachedValue int
}

func (w *WidgetGraph) Render(pane *Pane) {
}

func (w WidgetGraph) Iterate(pane *Pane) {
	pane.Write(0, 0, strconv.Itoa(w.cachedValue))
}

func (w WidgetGraph) HasBackend() bool {
	return true
}

func (w *WidgetGraph) Backend(ctx context.Context) {
	ticker := time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			v := rand.Intn(5)
			w.cachedValue = v
		}
	}
}
