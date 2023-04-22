package dices

type Dice interface {
	RollDice() (int64, error)
}
