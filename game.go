package poker

import "io"

// Game manages the state of a game.
type Game interface {
	Start(numberOfPlayers int,alertdestination io.Writer)
	Finish(winner string)
}
