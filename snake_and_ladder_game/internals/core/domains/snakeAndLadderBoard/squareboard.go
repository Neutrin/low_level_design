package snakeandladderboard

import (
	"fmt"
	"hexagnal/internals/core/domains/elements"
	"hexagnal/internals/core/domains/player"
)

type SquareBoard struct {
	size      int
	playerPos map[int]int
	snakesPos map[int]int
	ladderPos map[int]int
	isRunning bool
}

func GetSquareBoard(size int) Board {
	return &SquareBoard{size: size,
		playerPos: make(map[int]int, 0),
		snakesPos: make(map[int]int, 0),
		ladderPos: make(map[int]int),
	}
}

func (s *SquareBoard) GetPlayerPosition(player *player.Player) (int, error) {
	var (
		pos   int
		err   error
		found bool
	)
	if pos, found = s.playerPos[player.GetId()]; !found {
		err = fmt.Errorf(" player not found")
	}
	return pos, err
}

func (s *SquareBoard) AddPlayer(players []*player.Player) error {
	var err error
	if len(players) == 0 {
		err = fmt.Errorf(" add some player")
		return err
	}

	for _, curPlayer := range players {
		s.playerPos[curPlayer.GetId()] = 0
	}
	fmt.Println(" Players intialized")
	return nil
}
func (s *SquareBoard) IsWinningMove(player *player.Player) bool {
	return s.playerPos[player.GetId()] == s.size*s.size
}

func (s *SquareBoard) UpdPos(player *player.Player, moves int) error {
	if s.playerPos[player.GetId()]+moves > s.size*s.size {
		return fmt.Errorf(" try again exceeded limit")
	}
	s.playerPos[player.GetId()] = s.playerPos[player.GetId()] + moves
	return nil
}

func (s *SquareBoard) AddSnakes(snakes []elements.TwoEdgeElement) error {
	var err error
	for _, curSnake := range snakes {
		if _, exist := s.snakesPos[curSnake.GetStart()]; exist {
			err = fmt.Errorf(" snake exist at position  %d", curSnake.GetStart())
			break
		}
		if ladderStart, exists := s.ladderPos[curSnake.GetEnd()]; exists {
			if s.ladderPos[ladderStart] == curSnake.GetStart() {
				err = fmt.Errorf(" forming an infinit loop with ladder start = %d and end = %d",
					ladderStart, s.ladderPos[ladderStart])
				break
			}
		}
		s.snakesPos[curSnake.GetStart()] = curSnake.GetEnd()
	}
	return err
}

func (s *SquareBoard) AddLadders(ladder []elements.TwoEdgeElement) error {
	var err error
	for _, curLadder := range ladder {
		if _, exist := s.ladderPos[curLadder.GetStart()]; exist {
			err = fmt.Errorf(" ladder exist at the position = %d", curLadder.GetStart())
			break
		}
		if snakeStart, exist := s.snakesPos[curLadder.GetEnd()]; exist {
			if curLadder.GetStart() == s.snakesPos[snakeStart] {
				err = fmt.Errorf(" infinite loop for snake start = %d and end = %d",
					snakeStart, s.snakesPos[snakeStart])
				break
			}
		}
		s.ladderPos[curLadder.GetStart()] = curLadder.GetEnd()
	}
	return err
}
