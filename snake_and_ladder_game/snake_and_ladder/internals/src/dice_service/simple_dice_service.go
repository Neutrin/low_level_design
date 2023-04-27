package dice_service

import (
	"fmt"

	"github.com/neutrin/snake_and_ladder/internals/src/dices"
)

type SimpleDiceService struct {
	dice *dices.SimpleDice
}

func NewSimpleDiceService(dice *dices.SimpleDice) DiceService {
	return &SimpleDiceService{dice: dice}
}
func (simple *SimpleDiceService) RollDice() (int64, string, error) {
	var (
		rollCount int64
		err       error
		msg       string
	)
	random := simple.dice.Source()
	maxRange := simple.dice.MaxFace() - simple.dice.MinFace()
	if maxRange > 0 {
		randomNumber := random.Int63n(int64(maxRange + 1))
		rollCount = int64(simple.dice.MinFace()) + randomNumber
		msg = fmt.Sprintf(" dice roll = %d\n", rollCount)

	} else {
		err = fmt.Errorf(" dice not intialised correctly")
	}
	return rollCount, msg, err
}
