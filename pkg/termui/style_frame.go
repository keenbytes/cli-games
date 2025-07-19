package termui

// Frame represents one character size frame.
type Frame struct{}

// CornerChars contains characters used in a frame.
func (s Frame) CornerChars() [8]string {
	return [8]string{"┌", "─", "┐", "│", "┘", "─", "└", "│"}
}

// LeftFrameSize returns size of the left frame.
func (s Frame) LeftFrameSize() int {
	return 1
}

// RightFrameSize returns size of the right frame.
func (s Frame) RightFrameSize() int {
	return 1
}

// TopFrameSize returns size of the top frame.
func (s Frame) TopFrameSize() int {
	return 1
}

// BottomFrameSize returns size of the bottom frame.
func (s Frame) BottomFrameSize() int {
	return 1
}
