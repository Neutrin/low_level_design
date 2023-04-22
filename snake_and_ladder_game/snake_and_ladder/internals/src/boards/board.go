package boards

import "github.com/neutrin/snake_and_ladder/internals/src/elements"

type Board interface {
	AddSnakes(snakes []*elements.Snake) error
	AddLadders(ladders []*elements.Ladder) error
	MovePlayer(playerId int, diceCount int) (error, string)
	IsWinning(playerId int) (bool, error)
}
