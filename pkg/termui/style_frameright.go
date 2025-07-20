package termui

// FrameRight represents frame only on the right side.
type FrameRight struct{}

// CornerChars contains characters used in a frame.
func (s FrameRight) CornerChars() [8]string {
	return [8]string{"", "", "", "│", "", "", "", ""}
}

// LeftFrameSize returns size of the left frame.
func (s FrameRight) LeftFrameSize() int {
	return 0
}

// RightFrameSize returns size of the right frame.
func (s FrameRight) RightFrameSize() int {
	return 1
}

// TopFrameSize returns size of the top frame.
func (s FrameRight) TopFrameSize() int {
	return 0
}

// BottomFrameSize returns size of the bottom frame.
func (s FrameRight) BottomFrameSize() int {
	return 0
}
