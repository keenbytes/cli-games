package ortotris

import "strings"

// ChooseLeftLetter selects the first letter as the player's answer and inserts it into the current word.
func (g *Game) ChooseLeftLetter() {
	g.currentGuess = strings.Replace(g.currentWordWithPlaceholder, "_", g.lettersToChooseFrom[0], 1)
}

// ChooseRightLetter selects the second letter as the player's answer and inserts it into the current word.
func (g *Game) ChooseRightLetter() {
	g.currentGuess = strings.Replace(g.currentWordWithPlaceholder, "_", g.lettersToChooseFrom[1], 1)
}

// StopGame sets the game state to GameOver.
func (g *Game) StopGame() {
	g.state = GameOver
}

// StartGame resets all game variables and starts a new game.
func (g *Game) StartGame() {
	g.state = GameOn

	// Reset game state
	g.currentGuess = ""
	g.currentWordListIndex = 0
	g.currentLine = 0

	g.incorrectGuesses = []string{}
	g.numUsedWords = 0
}

// SetNextLineToLast jumps to the last line in next iteration.
func (g *Game) SetNextLineToLast() {
	g.iterateToLastLine = true
}
