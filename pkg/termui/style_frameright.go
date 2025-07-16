package termui

type FrameRight struct{}

func (s FrameRight) C() [8]string {
	return [8]string{"", "", "", "│", "", "", "", ""}
}

func (s FrameRight) L() int {
	return 0
}

func (s FrameRight) R() int {
	return 1
}

func (s FrameRight) T() int {
	return 0
}

func (s FrameRight) B() int {
	return 0
}
