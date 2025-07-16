package termui

import (
	"context"
	"time"
)

type WidgetBackend struct {
	cachedValue string
}

func (w *WidgetBackend) Render(pane *Pane) {
}

func (w WidgetBackend) Iterate(pane *Pane) {
	pane.Write(0, 0, "Time: "+w.cachedValue)
}

func (w WidgetBackend) HasBackend() bool {
	return true
}

func (w *WidgetBackend) Backend(ctx context.Context) {
	ticker := time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			w.cachedValue = time.Now().Format("15:04:05")
		}
	}
}
