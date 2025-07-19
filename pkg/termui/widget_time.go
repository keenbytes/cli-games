package termui

import (
	"context"
	"time"
)

// WidgetTime shows current time.
type WidgetTime struct{}

// Render draws the widget content onto the given pane.
func (w *WidgetTime) Render(_ *Pane) {
}

// Iterate is called on every termui iteration to update the widget's state or content.
func (w *WidgetTime) Iterate(pane *Pane) {
	now := time.Now()
	pane.Write(0, 0, now.Format("15:04:05"))
}

// HasBackend indicates whether the widget has a background process for computation or data retrieval.
func (w *WidgetTime) HasBackend() bool {
	return false
}

// Backend is a background function that performs processing or data fetching.
func (w *WidgetTime) Backend(_ context.Context) {
}
