package termui

type NoFrame struct {}
func (s NoFrame) C() [8]string {
	return [8]string{"", "", "", "", "", "", "", ""}
}
func (s NoFrame) L() int {
	return 0
}
func (s NoFrame) R() int {
	return 0
}
func (s NoFrame) T() int {
	return 0
}
func (s NoFrame) B() int {
	return 0
}
