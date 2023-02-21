package elements

import "fmt"

type Ladder struct {
	start int
	end   int
}

func GetNewLadder(start, end int) *Ladder {
	return &Ladder{start: start, end: end}
}

func (l *Ladder) GetStart() int {
	return l.start
}

func (l *Ladder) GetEnd() int {
	return l.end
}

func (l *Ladder) IsValid() error {
	var err error
	if l.start == l.end {
		err = fmt.Errorf(" start and end same")
	} else if l.start > l.end {
		err = fmt.Errorf(" start greater than end")
	}
	return err
}
