package rules_game

import "github.com/neutrin/snake_and_ladder/internals/src/dices"

type GameRules interface {
	RollDice(dice dices.Dice) (int64, string, error)
}
