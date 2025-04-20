package termui

const (
	NW = iota
	N
	NE
	E
	SE
	S
	SW
	W
)

type FrameStyle interface{
	C() [8]string
	L() int
	R() int
	T() int
	B() int
}

func (t *TermUI) SetFrame(style FrameStyle, panes ...*Pane) {
	for _, pane := range panes {
		pane.frame = style
	}
}
