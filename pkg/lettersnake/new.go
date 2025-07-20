package lettersnake

import (
	"bufio"
	"io"
	"math/rand/v2"
	"strings"
)

const (
	numLineWithTitle = 1
)

// NewGame returns new Game instance.
//
//nolint:mnd
func NewGame() *Game {
	return &Game{
		wordList:  []string{},
		direction: MovingDown,
		snake: []Segment{
			{PositionX: 3, PositionY: 5},
			{PositionX: 3, PositionY: 4},
			{PositionX: 3, PositionY: 3},
			{PositionX: 3, PositionY: 2},
			{PositionX: 3, PositionY: 1},
		},
	}
}

// ReadWords reads list of words from a text file.
func (g *Game) ReadWords(f io.Reader) {
	lineNumber := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineText := strings.TrimSpace(scanner.Text())
		if lineText == "" {
			continue
		}

		lineNumber++

		if lineNumber == numLineWithTitle {
			g.wordListTitle = lineText
		}

		g.wordList = append(g.wordList, lineText)
	}
}

// RandomizeWords reorders the word list in a random order.
func (g *Game) RandomizeWords() {
	rand.Shuffle(len(g.wordList), func(i, j int) {
		g.wordList[i], g.wordList[j] = g.wordList[j], g.wordList[i]
	})
}
