package elements

import "fmt"

type Snake struct {
	start int
	end   int
}

func (s *Snake) GetStart() int {
	return s.start
}

func (s *Snake) GetEnd() int {
	return s.end
}

func GetSnake(start int, end int) *Snake {
	return &Snake{start, end}
}

//IsValid() : Returns if snake is valid or not
func (sn *Snake) IsValid() error {
	if sn.start == sn.end {
		return fmt.Errorf(" start and end same ")
	} else if sn.start < sn.end {
		return fmt.Errorf(" start less than end")

	}
	return nil
}
