package termui

type Frame struct{}

func (s Frame) C() [8]string {
	return [8]string{"┌", "─", "┐", "│", "┘", "─", "└", "│"}
}

func (s Frame) L() int {
	return 1
}

func (s Frame) R() int {
	return 1
}

func (s Frame) T() int {
	return 1
}

func (s Frame) B() int {
	return 1
}
