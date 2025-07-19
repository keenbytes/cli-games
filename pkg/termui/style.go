package termui

const (
	// TopLeftChar is a character in the top left corner of the pane frame.
	TopLeftChar = iota

	// TopChar is a character used for the top of the pane frame.
	TopChar

	// TopRightChar is a character in the top right corner of the pane frame.
	TopRightChar

	// RightChar is a character used for the right of the pane frame.
	RightChar

	// BottomRightChar is a character in the bottom right corner of the pane frame.
	BottomRightChar

	// BottomChar is a character used for the bottom of the pane frame.
	BottomChar

	// BottomLeftChar is a character in the bottom left corner of the pane frame.
	BottomLeftChar

	// LeftChar is a character used for the left of the pane frame.
	LeftChar
)

// FrameStyle represents frame of the pane.
type FrameStyle interface {
	// CornerChars contains characters used in a frame.
	CornerChars() [8]string
	// LeftFrameSize returns size of the left frame.
	LeftFrameSize() int
	// RightFrameSize returns size of the right frame.
	RightFrameSize() int
	// TopFrameSize returns size of the top frame.
	TopFrameSize() int
	// BottomFrameSize returns size of the bottom frame.
	BottomFrameSize() int
}

// SetFrame sets specific frame style on the pane.
func (t *TermUI) SetFrame(style FrameStyle, panes ...*Pane) {
	for _, pane := range panes {
		pane.frame = style
	}
}
