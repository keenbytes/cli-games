package ortotris

import (
	"bufio"
	"io"
	"math/rand"
	"strings"
)

// State
const (
	NotStarted = iota
	GameOn
	GameOver
)

// Iterate result
const (
	_ = iota
	TopReached
	AllWordsUsed
	WrongAnswer
	CorrectAnswer
	ContinueGame
)

type Game struct {
	state int

	title               string
	words               []string
	letters             [2]string
	started             bool
	nextWordIndex       int
	currentWordTemplate string
	currentWord         string
	currentWordCorrect  string
	prevWordLine        int
	nextWordLine        int
	wordsNotGuessed     []string
	lastAvailableLine   int
	wordsGiven          int
	availableLines      int
	iterateToLast       bool
}

func NewGame() *Game {
	return &Game{
		state:           NotStarted,
		words:           []string{},
		letters:         [2]string{"", ""},
		wordsNotGuessed: []string{},
		availableLines:  20,
	}
}

func (g *Game) ReadWords(f io.Reader) {
	// TODO: Validation - for now, code assumes that the file contains correct data
	i := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		i++
		if i == 1 {
			g.title = line
		}
		if i == 2 {
			lineArr := strings.Split(line, ":")
			g.letters = [2]string{lineArr[0], lineArr[1]}
			continue
		}
		g.words = append(g.words, line)
	}
}

func (g *Game) RandomizeWords() {
	rand.Shuffle(len(g.words), func(i, j int) {
		g.words[i], g.words[j] = g.words[j], g.words[i]
	})
}

func (g *Game) Title() string {
	return g.title
}

func (g *Game) State() int {
	return g.state
}

func (g *Game) LeftLetter() string {
	return g.letters[0]
}

func (g *Game) RightLetter() string {
	return g.letters[1]
}

func (g *Game) CurrentWord() string {
	return g.currentWord
}

func (g *Game) CurrentLine() int {
	return g.nextWordLine - 1
}

func (g *Game) PrevCurrentLine() int {
	return g.prevWordLine
}

func (g *Game) NumCorrectAnswers() int {
	return g.wordsGiven - len(g.wordsNotGuessed)
}

func (g *Game) NumUsedWords() int {
	return g.wordsGiven
}

func (g *Game) NumAllWords() int {
	return len(g.words)
}

func (g *Game) StopGame() {
	g.state = GameOver
}

func (g *Game) StartGame() {
	g.state = GameOn

	g.nextWordIndex = 0
	g.currentWord = ""
	g.nextWordLine = 0
	g.wordsNotGuessed = []string{}
	g.wordsGiven = 0
	g.iterateToLast = false
}

func (g *Game) ChooseLeftLetter() {
	g.currentWord = strings.Replace(g.currentWordTemplate, "_", g.letters[0], 1)
}

func (g *Game) ChooseRightLetter() {
	g.currentWord = strings.Replace(g.currentWordTemplate, "_", g.letters[1], 1)
}

func (g *Game) SetAvailableLines(num int) {
	g.availableLines = num
}

func (g *Game) IsCurrentLineLast() bool {
	return g.nextWordLine == g.lastAvailableLine
}

func (g *Game) SetNextLineToLast() {
	g.iterateToLast = true
}

func (g *Game) Iterate() int {
	if g.state != GameOn {
		return NotStarted
	}

	// If there is no word then take the next one
	if g.isCurrentWordEmpty() {
		g.useNewWord()
	}

	// We need a position that is at the very bottom, remembering that all the
	// words which are not guess stay there
	g.lastAvailableLine = g.availableLines - 1 - len(g.wordsNotGuessed)

	if g.lastAvailableLine == 0 || g.nextWordIndex == len(g.words) {
		g.StopGame()
		g.wordsGiven++
		if g.lastAvailableLine == 0 {
			return TopReached
		} else {
			return AllWordsUsed
		}
	}

	// If the word is already in the last line
	if g.IsCurrentLineLast() {
		g.wordsGiven++
		if g.currentWord != g.currentWordCorrect {
			g.wordsNotGuessed = append(g.wordsNotGuessed, g.currentWord)
			g.currentWord = ""
			return WrongAnswer
		} else {
			g.currentWord = ""
			return CorrectAnswer
		}
	}

	// Increment the line for the next iteration
	g.prevWordLine = g.nextWordLine
	if g.iterateToLast && g.nextWordLine < g.lastAvailableLine-1 {
		g.nextWordLine = g.lastAvailableLine - 1
	} else {
		g.nextWordLine++
	}
	g.iterateToLast = false

	// If the word is not at the bottom then just continue
	return ContinueGame
}

func (g *Game) isCurrentWordEmpty() bool {
	return g.currentWord == ""
}

func (g *Game) useNewWord() {
	currentWordArr := strings.Split(g.words[g.nextWordIndex], ":")
	g.currentWordTemplate = currentWordArr[0]
	g.currentWord = g.currentWordTemplate
	g.currentWordCorrect = strings.Replace(g.currentWordTemplate, "_", currentWordArr[1], 1)
	g.nextWordIndex++
	g.nextWordLine = 0
}
