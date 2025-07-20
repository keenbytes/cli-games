package lettersnake

// StopGame sets the game state to GameOver.
func (g *Game) StopGame() {
	g.state = GameOver
}

// StartGame resets all game variables and starts a new game.
func (g *Game) StartGame() {
	g.state = GameOn

	g.currentWordListIndex = 0
	g.currentWord = ""
	g.currentTranslation = ""

	g.numUsedWords = 0
}

// SetDirection changes direction snake is moving.
func (g *Game) SetDirection(direction int) {
	g.direction = direction
}
