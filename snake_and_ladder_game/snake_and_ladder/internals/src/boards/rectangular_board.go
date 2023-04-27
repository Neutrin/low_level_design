package boards

import (
	"fmt"
	"sync/atomic"

	"github.com/neutrin/snake_and_ladder/internals/src/elements"
	"github.com/neutrin/snake_and_ladder/internals/src/players"
)

var boardId uint64

type ReactangularBoard struct {
	id        int
	rows      int
	cols      int
	snakesPos map[int]int
	ladderPos map[int]int
	playerPos map[int]int
}

const startingPoint = 0

/*
TODO
how can i apply validation in other object oriented design way
This constructor is too big how can i reduce size of this
*/
func NewReactangularBoard(rows int, cols int, players []*players.Player) (Board, error) {
	board := &ReactangularBoard{}
	if rows < 0 || cols < 0 {
		return board, fmt.Errorf(" rows and cols cannot be less than zero")
	}
	board.rows = rows
	board.cols = cols

	board.playerPos = make(map[int]int)
	for _, curPlayer := range players {

		board.playerPos[curPlayer.GetId()] = startingPoint
	}
	atomic.AddUint64(&boardId, 1)
	board.id = int(boardId)
	return board, nil
}

func (r *ReactangularBoard) AddLadders(ladders []*elements.Ladder) error {
	r.ladderPos = make(map[int]int)
	for _, curLadder := range ladders {
		if _, ladderExist := r.ladderPos[curLadder.GetLower()]; ladderExist {
			return fmt.Errorf("two ladders with same start")
		}
		r.ladderPos[curLadder.GetLower()] = curLadder.GetUpper()
	}
	return nil
}

func (r *ReactangularBoard) AddSnakes(snakes []*elements.Snake) error {
	r.snakesPos = make(map[int]int)
	for _, curSnake := range snakes {
		if _, snakeExist := r.snakesPos[curSnake.GeyHead()]; snakeExist {
			fmt.Errorf("two snakes at same position")
		}
		r.snakesPos[curSnake.GeyHead()] = curSnake.GetTail()
	}
	return nil
}

func (r *ReactangularBoard) MovePlayer(playerId int, diceCount int) (error, string) {
	var (
		curPos      int
		msg         string
		playerExist bool
	)
	if curPos, playerExist = r.playerPos[playerId]; !playerExist {
		return fmt.Errorf("player does not exists"), msg
	}
	if curPos+diceCount <= r.rows*r.cols {
		msg = fmt.Sprintf(" moving from %d to %d \n", curPos, curPos+diceCount)
		r.playerPos[playerId] = curPos + diceCount
		nextPos := curPos + diceCount
		if snakeTail, snakeExists := r.snakesPos[nextPos]; snakeExists {
			msg += fmt.Sprintf(" snake found moving from %d to %d \n", nextPos, snakeTail)
			r.playerPos[playerId] = snakeTail
		}
		if ladderHead, ladderExists := r.ladderPos[nextPos]; ladderExists {
			msg += fmt.Sprintf(" ladder found moving from %d to %d \n", nextPos, ladderHead)
			r.playerPos[playerId] = ladderHead
		}

	} else {
		msg = fmt.Sprintln(" not moving as out of bound")
	}
	return nil, msg

}

func (r *ReactangularBoard) IsWinning(playerId int) (bool, error) {
	var (
		curPos      int
		playerExist bool
	)
	if curPos, playerExist = r.playerPos[playerId]; !playerExist {
		return false, fmt.Errorf("player does not exists")
	}
	if curPos == r.rows*r.cols {
		return true, nil
	}
	return false, nil
}
