package lettersnake

import (
	"bufio"
	"io"
	"math/rand/v2"
	"strings"
)

// State
const (
	NotStarted = iota
	GameOn
	GameOver
)

// Direction
const (
	Down = iota
	Up
	Left
	Right
)

// Iterate result
const (
	_ = iota
	EdgeHit
	AteItself
	AllWordsUsed
	ContinueGame
)

type Letter struct {
	X int
	Y int
	L string
}

type Segment struct {
	X int
	Y int
}

type Game struct {
	state int
	title string
	words              []string
	nextWordIndex      int
	currentWord        string
	currentTranslation string
	wordsGiven         int
	wordsCorrect       int
	letters            []Letter
	size               [2]int
	direction          int
	snake              []Segment
	remove             *Segment
	consumedLetters    string
	sizeSet bool
}

func NewGame() *Game {
	return &Game{
		words:     []string{},
		direction: Down,
		snake: []Segment{
			{X: 3, Y: 5},
			{X: 3, Y: 4},
			{X: 3, Y: 3},
			{X: 3, Y: 2},
			{X: 3, Y: 1},
		},
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

func (g *Game) CurrentWord() string {
	return g.currentWord
}

func (g *Game) CurrentTranslation() string {
	return g.currentTranslation
}

func (g *Game) ConsumedLetters() string {
	return g.consumedLetters
}

func (g *Game) Letters() []Letter {
	return g.letters
}

func (g *Game) Remove() *Segment {
	return g.remove
}

func (g *Game) Snake() []Segment {
	return g.snake
}

func (g *Game) NumUsedWords() int {
	return g.wordsGiven
}

func (g *Game) NumCorrectWords() int {
	return g.wordsCorrect
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
	g.currentTranslation = ""
	g.wordsGiven = 0
}

func (g *Game) Direction() int {
	return g.direction
}

func (g *Game) SetDirection(direction int) {
	g.direction = direction
}

func (g *Game) SetSize(w int, h int) {
	g.size[0] = w
	g.size[1] = h
	g.sizeSet = true
}

func (g *Game) SizeSet() bool {
	return g.sizeSet
}

func (g *Game) Iterate() int {
	if g.state != GameOn {
		return NotStarted
	}

	// If there is no word then take the next one
	if g.isCurrentWordEmpty() {
		g.useNewWord()
	}

	// Check hitting edges and eating itself (going backwards)
	switch g.direction {
	case Down:
		if g.snake[0].X == g.snake[1].X && g.snake[0].Y == g.snake[1].Y {
			g.StopGame()
			return AteItself
		}
		if g.snake[0].Y == g.size[1]-1 {
			g.StopGame()
			return EdgeHit
		}
	case Up:
		if g.snake[0].X == g.snake[1].X && g.snake[1].Y == g.snake[0].Y {
			g.StopGame()
			return AteItself
		}
		if g.snake[0].Y == 0 {
			g.StopGame()
			return EdgeHit
		}
	case Left:
		if g.snake[0].Y == g.snake[1].Y && g.snake[0].X == g.snake[1].X {
			g.StopGame()
			return AteItself
		}
		if g.snake[0].X == 0 {
			g.StopGame()
			return EdgeHit
		}
	case Right:
		if g.snake[0].Y == g.snake[1].Y && g.snake[1].X == g.snake[0].X {
			g.StopGame()
			return AteItself
		}
		if g.snake[0].X == g.size[0]-1 {
			g.StopGame()
			return EdgeHit
		}
	}

	consumedLetter := false
	newLetters := []Letter{}
	var addTail *Segment
	for _, l := range g.letters {
		if g.snake[0].X == l.X && g.snake[0].Y == l.Y {
			g.consumedLetters += l.L
			consumedLetter = true
		} else {
			newLetters = append(newLetters, l)
		}
	}
	g.letters = newLetters

	if !consumedLetter {
		g.remove = &Segment{
			X: g.snake[len(g.snake)-1].X,
			Y: g.snake[len(g.snake)-1].Y,
		}
	} else {
		g.remove = nil
		addTail = &Segment{
			X: g.snake[len(g.snake)-1].X,
			Y: g.snake[len(g.snake)-1].Y,
		}
	}

	for i := len(g.snake) - 1; i > 0; i-- {
		g.snake[i].X = g.snake[i-1].X
		g.snake[i].Y = g.snake[i-1].Y
	}
	if addTail != nil {
		g.snake = append(g.snake, *addTail)
	}

	switch g.direction {
	case Down:
		g.snake[0].Y++
	case Up:
		g.snake[0].Y--
	case Left:
		g.snake[0].X--
	case Right:
		g.snake[0].X++
	}

	if len(g.letters) == 0 {
		if g.currentWord == g.consumedLetters {
			g.wordsCorrect++
		}
		if g.nextWordIndex == len(g.words) {
			g.StopGame()
			return AllWordsUsed
		}
		g.useNewWord()
	}

	return ContinueGame
}

func (g *Game) isCurrentWordEmpty() bool {
	return g.currentWord == ""
}

func (g *Game) useNewWord() {
	curWordArr := strings.Split(g.words[g.nextWordIndex], ":")
	g.currentWord = curWordArr[1]
	g.currentTranslation = curWordArr[0]
	g.setLettersFromCurrentWord()
	g.nextWordIndex++
	g.wordsGiven++
	g.consumedLetters = ""
}

func (g *Game) setLettersFromCurrentWord() {
	g.letters = make([]Letter, 0)
	for i := 0; i < len(g.currentWord); i++ {
		g.letters = append(g.letters, Letter{
			X: rand.IntN(g.size[0]-2) + 1,
			Y: rand.IntN(g.size[1]-2) + 1,
			L: string(g.currentWord[i]),
		})
	}
}
