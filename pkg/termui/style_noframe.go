package termui

// NoFrame represents no frame.
type NoFrame struct{}

// CornerChars contains characters used in a frame.
func (s NoFrame) CornerChars() [8]string {
	return [8]string{"", "", "", "", "", "", "", ""}
}

// LeftFrameSize returns size of the left frame.
func (s NoFrame) LeftFrameSize() int {
	return 0
}

// RightFrameSize returns size of the right frame.
func (s NoFrame) RightFrameSize() int {
	return 0
}

// TopFrameSize returns size of the top frame.
func (s NoFrame) TopFrameSize() int {
	return 0
}

// BottomFrameSize returns size of the bottom frame.
func (s NoFrame) BottomFrameSize() int {
	return 0
}
