package lettersnake

// SetPlayAreaSize sets size of the play area on which the snake can move.
func (g *Game) SetPlayAreaSize(width int, height int) {
	g.playAreaSize[0] = width
	g.playAreaSize[1] = height
	g.playAreaSizeSet = true
}

// IsPlayAreaSizeSet determines whether the play area size was set.
func (g *Game) IsPlayAreaSizeSet() bool {
	return g.playAreaSizeSet
}
