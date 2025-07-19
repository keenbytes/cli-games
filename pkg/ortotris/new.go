package ortotris

import (
	"bufio"
	"io"
	"math/rand/v2"
	"strings"
)

const (
	defaultNumAvailableLines = 20
	numLineWithTitle         = 1
	numLineWithLetters       = 2
)

// NewGame returns a pointer to a new Game.
func NewGame() *Game {
	return &Game{
		state:               NotStarted,
		wordList:            []string{},
		lettersToChooseFrom: [2]string{"", ""},
		incorrectGuesses:    []string{},
		numAvailableLines:   defaultNumAvailableLines,
	}
}

// ReadWords reads list of words from a text file.
func (g *Game) ReadWords(file io.Reader) {
	lineNumber := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := strings.TrimSpace(scanner.Text())
		if lineText == "" {
			continue
		}

		lineNumber++

		// first line contains title of the word list
		if lineNumber == numLineWithTitle {
			g.wordListTitle = lineText

			continue
		}

		// second line contains letters to choose from, eg. h:ch
		if lineNumber == numLineWithLetters {
			lineLetters := strings.Split(lineText, ":")
			g.lettersToChooseFrom = [2]string{lineLetters[0], lineLetters[1]}

			continue
		}

		g.wordList = append(g.wordList, lineText)
	}
}

// RandomizeWords reorders list of words in a random order.
func (g *Game) RandomizeWords() {
	rand.Shuffle(len(g.wordList), func(i, j int) {
		g.wordList[i], g.wordList[j] = g.wordList[j], g.wordList[i]
	})
}
