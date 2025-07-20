package termui

import "context"

// Widget represents a UI component that can be rendered inside a pane.
type Widget interface {
	// Render draws the widget content onto the given pane.
	Render(pane *Pane)

	// Iterate is called on every termui iteration to update the widget's state or content.
	Iterate(pane *Pane)

	// HasBackend indicates whether the widget has a background process for computation or data retrieval.
	HasBackend() bool

	// Backend is a background function that performs processing or data fetching.
	Backend(ctx context.Context)
}
