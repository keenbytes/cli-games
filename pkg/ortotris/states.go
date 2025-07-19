package ortotris

// Game state.
const (
	// NotStarted represents the state before the game begins.
	NotStarted = iota

	// GameOn represents the state while the game is in progress.
	GameOn

	// GameOver represents the state after the game has ended.
	GameOver
)

// Iteration result.
const (
	_ = iota

	// TopReached indicates that the screen is full and no space is left for new falling words.
	TopReached

	// AllWordsUsed means that all words from the input file have been used.
	AllWordsUsed

	// WrongAnswer indicates that the word reached the bottom and the selected answer was incorrect.
	WrongAnswer

	// CorrectAnswer indicates that the word reached the bottom and the selected answer was correct.
	CorrectAnswer

	// ContinueGame means no special condition occurred; the word simply moves down one slot.
	ContinueGame

	// JumpToLastLine indicates a jump to the last available line.
	JumpToLastLine
)
