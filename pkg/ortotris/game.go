package ortotris

import (
	"strings"
)

// Game holds the state and data of the game.
type Game struct {
	// state is an actual state of the game
	state int

	// wordListTitle is title of the word list
	wordListTitle string

	// wordList represents words that have to be guessed
	wordList []string

	// letters are two letters to choose from
	lettersToChooseFrom [2]string

	// currentGuess represents word with a letter selected by the player
	currentGuess string

	// currentWordListIndex indicates which word from the list is currently being guessed (falling)
	currentWordListIndex int

	// currentWordWithPlaceholder is word with missing letter, just like it is from the input
	currentWordWithPlaceholder string

	// currentWordCorrect is the correct answer, so word from the input with correct letter put in it
	currentWordCorrect string

	// numAvailableLines is number of lines available on the screen
	numAvailableLines int

	// currentLine represents number of the line that the current guess word is
	currentLine int

	// previousLine represents number of the line that the current guess word moved from
	previousLine int

	// incorrectGuesses contains words that haven't been guessed
	incorrectGuesses []string

	// numUsedWords represents number of words that have been taken from the list so far
	numUsedWords int

	// lastAvailableLine represents the position of the last line available for falling words.
	// Incorrectly guessed words accumulate at the bottom, reducing the number of available lines.
	lastAvailableLine int

	// iterateToLastLine signals that the word should skip to the last available line immediately.
	iterateToLastLine bool
}

// IsCurrentLineLast returns true if the player's word is on the last available line.
func (g *Game) IsCurrentLineLast() bool {
	return g.currentLine == g.lastAvailableLine
}

// Iterate runs one iteration of the game.
func (g *Game) Iterate() int {
	if g.state != GameOn {
		return NotStarted
	}

	// take new word from the list if current guess (player's word) is empty
	if g.shouldTakeNewWord() {
		more := g.useNewWordFromTheList()
		if !more {
			g.StopGame()

			return AllWordsUsed
		}
	}

	// get the last available position, remembering that number of available lines shrink whenever there is an
	// incorrent guess (word stays at the bottom)
	g.lastAvailableLine = g.numAvailableLines - len(g.incorrectGuesses) - 1
	if g.lastAvailableLine == 0 {
		g.StopGame()

		g.numUsedWords++

		return TopReached
	}

	// word is already in the last available line
	if g.IsCurrentLineLast() {
		g.numUsedWords++

		// checking if guess is correct
		if g.currentGuess != g.currentWordCorrect {
			g.incorrectGuesses = append(g.incorrectGuesses, g.currentGuess)
			g.currentWordWithPlaceholder = ""

			return WrongAnswer
		}

		g.currentWordWithPlaceholder = ""

		return CorrectAnswer
	}

	if g.iterateToLastLine && g.currentLine != g.lastAvailableLine {
		g.previousLine = g.currentLine
		g.currentLine = g.lastAvailableLine

		g.iterateToLastLine = false

		return JumpToLastLine
	}

	g.previousLine = g.currentLine
	g.currentLine++

	return ContinueGame
}

func (g *Game) shouldTakeNewWord() bool {
	return g.currentWordWithPlaceholder == ""
}

// useNewWordFromTheList takes the new word from the input list if available and returns true.
// Otherwise returns false.
func (g *Game) useNewWordFromTheList() bool {
	if g.currentWordListIndex == len(g.wordList) {
		return false
	}

	nextWord := g.wordList[g.currentWordListIndex]

	nextWordTrimmed := strings.TrimSpace(nextWord)
	if nextWordTrimmed == "" {
		return false
	}

	nextWordArray := strings.Split(nextWordTrimmed, ":")
	if len(nextWordArray) != 2 || nextWordArray[0] == "" || nextWordArray[1] == "" {
		return false
	}

	g.currentWordWithPlaceholder = nextWordArray[0]
	g.currentWordCorrect = strings.Replace(g.currentWordWithPlaceholder, "_", nextWordArray[1], 1)

	g.currentGuess = nextWordArray[0]

	g.currentWordListIndex++
	g.currentLine = 0
	g.previousLine = -1

	return true
}
