package dice_service

type DiceService interface {
	RollDice() (int64, string, error)
}
