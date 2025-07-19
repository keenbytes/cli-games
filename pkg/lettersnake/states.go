package lettersnake

// Game state.
const (
	// NotStarted indicates that the game has not started yet.
	NotStarted = iota

	// GameOn indicates that the game is currently in progress.
	GameOn

	// GameOver indicates that the game has ended.
	GameOver
)

// Direction represents the direction in which the snake is moving.
const (
	// Down indicates that the snake is moving downward.
	Down = iota

	// Up indicates that the snake is moving upward.
	Up

	// Left indicates that the snake is moving to the left.
	Left

	// Right indicates that the snake is moving to the right.
	Right
)

// Game iteration result.
const (
	_ = iota

	// EdgeHit indicates that the snake has hit the edge of the screen.
	EdgeHit

	// AteItself indicates that the snake has collided with its own body.
	AteItself

	// AllWordsUsed means all words in the list have already been used.
	AllWordsUsed

	// ContinueGame indicates that the game should continue running.
	ContinueGame
)
