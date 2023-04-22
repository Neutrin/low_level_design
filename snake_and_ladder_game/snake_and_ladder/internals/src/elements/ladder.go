package elements

import "fmt"

type Ladder struct {
	lower int
	upper int
}

/*
TODO
Could have used chain of responsibility also here to create this object
need to validate what is better */
func NewLadder(lower int, upper int) (*Ladder, error) {
	if lower >= upper {
		return nil, fmt.Errorf(" lower should be less than upper ")
	}
	return &Ladder{lower: lower, upper: upper}, nil
}

func (l *Ladder) GetLower() int {
	return l.lower
}

func (l *Ladder) GetUpper() int {
	return l.upper
}
