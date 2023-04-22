package elements

import "fmt"

type Snake struct {
	head int
	tail int
}

func NewSnake(head int, tail int) (*Snake, error) {
	if head <= tail {
		return nil, fmt.Errorf("head should be greater than tail ")
	}
	return &Snake{
		head: head,
		tail: tail,
	}, nil
}

func (s *Snake) GeyHead() int {
	return s.head
}

func (s *Snake) GetTail() int {
	return s.tail
}
