package dice

import (
	"math/rand"
	"time"
)

type SingleDice struct {
	max int
}

var (
	source rand.Source
	r      *rand.Rand
)

func GetSingleDice(max int) SingleDice {
	source = rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
	return SingleDice{max}
}

func (dice *SingleDice) Throw() int {
	return (1 + rand.Intn(dice.max))
}
