package lettersnake

import (
	"math/rand/v2"
	"strings"
)

// Game holds the state and data of the game.
type Game struct {
	// state is an actual state of the game
	state int

	// direction indicates which direction snake is moving
	direction int

	// wordListTitle is title of the word list
	wordListTitle string

	// wordList represents words that have to be guessed
	wordList []string

	// currentWordListIndex indicates which word from the list is currently being guessed (falling)
	currentWordListIndex int

	// currentWord contains current word from the input
	currentWord string

	// currentTranslation contains translation for the current word
	currentTranslation string

	// letters contains positions of letters on the screen
	letters *map[int]map[int]rune

	// numLettersLeft contains number of letters left on the play area
	lettersLeft int

	// consumedLetters contain letters that have been eaten by the snake
	consumedLetters string

	// correctGuesses contains words that have been guessed
	correctGuesses []string

	// numUsedWords represents number of words that have been taken from the list so far
	numUsedWords int

	// playAreaSize represents size of playable area within snake can move.
	playAreaSize [2]int

	// playAreaSizeSet indicates that the playAreaSize has been set.
	playAreaSizeSet bool

	// snake contains Segments that connected are the snake.
	snake []Segment

	// tail references the last Segment of the snake.
	tail *Segment
}

// Iterate runs one iteration of the game.
//
//nolint:funlen
func (g *Game) Iterate() int {
	if g.state != GameOn {
		return NotStarted
	}

	// take new word from the list
	if g.shouldTakeNewWord() {
		more := g.useNewWordFromTheList()
		if !more {
			g.StopGame()

			return AllWordsUsed
		}
	}

	// checking if snake is not biting themselves by moving backwards to the same position
	if g.snake[0].PositionX == g.snake[1].PositionX &&
		g.snake[0].PositionY == g.snake[1].PositionY {
		g.StopGame()

		return AteItself
	}

	// checking if snake is not hitting the edge
	switch g.direction {
	case MovingDown:
		if g.snake[0].PositionY == g.playAreaSize[1]-1 {
			g.StopGame()

			return EdgeHit
		}
	case MovingUp:
		if g.snake[0].PositionY == 0 {
			g.StopGame()

			return EdgeHit
		}
	case MovingLeft:
		if g.snake[0].PositionX == 0 {
			g.StopGame()

			return EdgeHit
		}
	case MovingRight:
		if g.snake[0].PositionX == g.playAreaSize[0]-1 {
			g.StopGame()

			return EdgeHit
		}
	}

	letterHasBeenConsumed := false

	letters := *g.letters

	_, posXExists := letters[g.snake[0].PositionX]
	if posXExists {
		foundLetter, posYExists := letters[g.snake[0].PositionX][g.snake[0].PositionY]
		if posYExists {
			letterHasBeenConsumed = true

			delete(letters[g.snake[0].PositionX], g.snake[0].PositionY)

			g.lettersLeft--
			g.consumedLetters += string(foundLetter)
		}
	}

	var tailToAdd *Segment

	if !letterHasBeenConsumed {
		// tail needs to be removed
		g.tail = &Segment{
			PositionX: g.snake[len(g.snake)-1].PositionX,
			PositionY: g.snake[len(g.snake)-1].PositionY,
		}
	} else {
		// nothing should be removed, snake gets longer
		g.tail = nil

		tailToAdd = &Segment{
			PositionX: g.snake[len(g.snake)-1].PositionX,
			PositionY: g.snake[len(g.snake)-1].PositionY,
		}
	}

	for i := len(g.snake) - 1; i > 0; i-- {
		g.snake[i].PositionX = g.snake[i-1].PositionX
		g.snake[i].PositionY = g.snake[i-1].PositionY
	}

	if tailToAdd != nil {
		g.snake = append(g.snake, *tailToAdd)
	}

	switch g.direction {
	case MovingDown:
		g.snake[0].PositionY++
	case MovingUp:
		g.snake[0].PositionY--
	case MovingLeft:
		g.snake[0].PositionX--
	case MovingRight:
		g.snake[0].PositionX++
	}

	if g.lettersLeft > 0 {
		return ContinueGame
	}

	if g.currentWord == g.consumedLetters {
		g.correctGuesses = append(g.correctGuesses, g.currentWord)
	}

	return ContinueGame
}

func (g *Game) shouldTakeNewWord() bool {
	return g.lettersLeft == 0
}

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

	g.currentWord = nextWordArray[0]
	g.currentTranslation = nextWordArray[1]

	g.placeLettersFromCurrentWordRandomlyOnThePlayArea()

	g.currentWordListIndex++
	g.consumedLetters = ""

	return true
}

//nolint:gosec,mnd
func (g *Game) placeLettersFromCurrentWordRandomlyOnThePlayArea() {
	lettersMap := map[int]map[int]rune{}

	for _, wordLetter := range g.currentWord {
		positionX := rand.IntN(g.playAreaSize[0]-2) + 1
		positionY := rand.IntN(g.playAreaSize[1]-2) + 1

		if lettersMap[positionX] == nil {
			lettersMap[positionX] = map[int]rune{}
		}

		lettersMap[positionX][positionY] = wordLetter
	}

	g.letters = &lettersMap
	g.lettersLeft = len(g.currentWord)
}
