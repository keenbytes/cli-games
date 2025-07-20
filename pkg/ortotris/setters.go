package ortotris

// SetAvailableLines sets how many lines a word can fall through.
func (g *Game) SetAvailableLines(num int) {
	g.numAvailableLines = num
}
