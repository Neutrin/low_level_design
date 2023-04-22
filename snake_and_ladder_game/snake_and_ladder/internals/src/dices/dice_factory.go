package dices

/*
This can be altered later to new kind of factory
design which will be much more robust
*/

func NewDiceByType(diceTyp string) Dice {
	var dice Dice
	switch diceTyp {
	case "default":
		dice = NewSimpleDice(1, 6)
	}
	return dice
}
