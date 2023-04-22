package dices

import (
	"fmt"
	"math/rand"
	"time"
)

type SimpleDice struct {
	min  int
	max  int
	rand *rand.Rand
}

//Validation exlcuded as of now
func NewSimpleDice(min int, max int) *SimpleDice {
	source := rand.NewSource(time.Now().Unix())
	return &SimpleDice{
		min:  min,
		max:  max,
		rand: rand.New(source),
	}
}

func (s *SimpleDice) RollDice() (int64, error) {
	if s.rand == nil {
		return 0, fmt.Errorf(" souce not intialized")
	}
	return int64(s.rand.Intn(s.max-s.min) + s.min), nil
}
