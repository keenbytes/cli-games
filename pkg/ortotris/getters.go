package ortotris

// WordListTitle returns the title of the current word list.
func (g *Game) WordListTitle() string {
	return g.wordListTitle
}

// State returns the current state of the game.
func (g *Game) State() int {
	return g.state
}

// LeftLetter returns the first letter the player can choose from.
func (g *Game) LeftLetter() string {
	return g.lettersToChooseFrom[0]
}

// RightLetter returns the second letter the player can choose from.
func (g *Game) RightLetter() string {
	return g.lettersToChooseFrom[1]
}

// CurrentGuess returns the word the player has selected as their answer.
func (g *Game) CurrentGuess() string {
	return g.currentGuess
}

// NumUsedWords returns the number of words used so far in the game.
func (g *Game) NumUsedWords() int {
	return g.numUsedWords
}

// NumWordList returns the total number of words in the input list.
func (g *Game) NumWordList() int {
	return len(g.wordList)
}

// CurrentLine returns the index of the line where the player's current guess is positioned.
func (g *Game) CurrentLine() int {
	return g.currentLine
}

// PreviousLine returns the index of the previous line where the player's current guess word was.
func (g *Game) PreviousLine() int {
	return g.previousLine
}

// NumCorrectGuesses returns the number of correctly guessed words.
func (g *Game) NumCorrectGuesses() int {
	return g.numUsedWords - len(g.incorrectGuesses)
}
