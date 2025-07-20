package lettersnake

// WordListTitle returns title of the word list.
func (g *Game) WordListTitle() string {
	return g.wordListTitle
}

// State returns current state of the game.
func (g *Game) State() int {
	return g.state
}

// CurrentWord returns the current word that is being guessed by the player.
func (g *Game) CurrentWord() string {
	return g.currentWord
}

// CurrentTranslation returns the translation of the current word.
func (g *Game) CurrentTranslation() string {
	return g.currentTranslation
}

// ConsumedLetters returns letters that have been consumed by the snake so far.
func (g *Game) ConsumedLetters() string {
	return g.consumedLetters
}

// Letters returns two-dimensional map of letter positions on the play area.
func (g *Game) Letters() *map[int]map[int]rune {
	return g.letters
}

// NumUsedWords returns the number of words used so far in the game.
func (g *Game) NumUsedWords() int {
	return g.numUsedWords
}

// NumWordList returns the total number of words in the input list.
func (g *Game) NumWordList() int {
	return len(g.wordList)
}

// NumCorrectGuesses returns the number of correctly guessed words.
func (g *Game) NumCorrectGuesses() int {
	return len(g.correctGuesses)
}

// Direction returns direction snake is moving.
func (g *Game) Direction() int {
	return g.direction
}

// Tail returns last Segment of the snake.
func (g *Game) Tail() *Segment {
	return g.tail
}

// Snake returns snake's Segments.
func (g *Game) Snake() []Segment {
	return g.snake
}
