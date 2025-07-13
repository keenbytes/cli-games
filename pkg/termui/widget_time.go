package termui

import (
	"context"
	"time"
)

type WidgetTime struct{}

func (w *WidgetTime) Render(pane *Pane) {
}

func (w WidgetTime) Iterate(pane *Pane) {
	now := time.Now()
	pane.Write(0, 0, now.Format("15:04:05"))
}

func (w WidgetTime) HasBackend() bool {
	return false
}

func (w *WidgetTime) Backend(ctx context.Context) {
}
