package dices

import (
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

func (s *SimpleDice) MinFace() int {
	return s.min
}

func (s *SimpleDice) MaxFace() int {
	return s.max
}

func (s *SimpleDice) Source() *rand.Rand {
	return s.rand
}
