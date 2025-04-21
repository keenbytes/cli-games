package termui

import "context"

type Widget interface {
	Render(pane *Pane)
	Iterate(pane *Pane)
	HasBackend() bool
	Backend(ctx context.Context)
}
