package termui

import (
	"context"
	"time"
)

// WidgetBackend is a sample widget with a backend.
type WidgetBackend struct {
	cachedValue string
}

// Render draws on the pane.
func (w *WidgetBackend) Render(_ *Pane) {
}

// Iterate is called on every termui iteration to update the widget's state or content.
func (w *WidgetBackend) Iterate(pane *Pane) {
	pane.Write(0, 0, "Time: "+w.cachedValue)
}

// HasBackend indicates whether the widget has a background process for computation or data retrieval.
func (w *WidgetBackend) HasBackend() bool {
	return true
}

// Backend is a background function that performs processing or data fetching.
func (w *WidgetBackend) Backend(ctx context.Context) {
	//nolint:mnd
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
