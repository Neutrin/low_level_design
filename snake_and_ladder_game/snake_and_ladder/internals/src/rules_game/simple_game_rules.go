package rules_game

import (
	"fmt"

	"github.com/neutrin/snake_and_ladder/internals/src/dices"
)

//ToDO : this can be made singleton
type SimpleGameRules struct {
}

func NewSimpleGameRules() GameRules {
	return &SimpleGameRules{}
}

func (rules *SimpleGameRules) RollDice(dice dices.Dice) (int64, string, error) {
	var (
		count int64
		msg   string
		err   error
	)
	if dice == nil {
		return count, msg, fmt.Errorf(" dice not intialized")
	}
	if count, err = dice.RollDice(); err == nil {
		msg = fmt.Sprintf(" dice rolled = %d", count)
	}
	return count, msg, err

}
