package src

import (
	"fmt"
	"time"

	queue "github.com/enriquebris/goconcurrentqueue"
	"github.com/neutrin/snake_and_ladder/internals/src/boards"
	"github.com/neutrin/snake_and_ladder/internals/src/dice_service"
	"github.com/neutrin/snake_and_ladder/internals/src/players"
)

type GameManager struct {
	board       boards.Board
	players     *queue.FIFO
	playerRanks []*players.Player
	isRunning   bool
	/* IMPROVEMENT
	#1 option requirement have a slice of dices
	DONE
	*/
	Dices []dice_service.DiceService
	/*
		This field is for the optional requirement number 3

	*/
	isMultiWinner bool
}

func NewGameManager(board boards.Board, players []*players.Player, dice []dice_service.DiceService) *GameManager {
	fifo := queue.NewFIFO()
	for _, curPlayer := range players {
		fifo.Enqueue(curPlayer)
	}
	return &GameManager{
		board:   board,
		players: fifo,
		//playerRanks: curPlayerList,
		isRunning:     true,
		Dices:         dice,
		isMultiWinner: false,
	}

}

func (m *GameManager) IsMultiWinner() {
	m.isMultiWinner = true
}

func (m *GameManager) Play() error {
	var (
		err       error
		curPlayer *players.Player
		diceFace  int64
		msg       string
		isWinning bool
	)
	if !m.isRunning {
		return fmt.Errorf("game not running")
	}
	if m.players.GetLen() == 0 {
		return fmt.Errorf("game empty")
	}
	m.playerRanks = make([]*players.Player, 0)

	for m.isRunning {
		curEle, _ := m.players.Dequeue()
		curPlayer, _ = curEle.(*players.Player)
		/*IMPROVEMENT

		if dice was a slice here it would have been better
		to just run a loop and print their messages accordingly
		now this will also be based on the requirements if they wanted this or not

		4> OPTIONAL REQUIREMENT COULD BE EASILY ADDED HERE IF THERE WERE CHANGES IN
		HTHE RULES OF THE GAME IE THAT 3 TIMES MORE THAN 6 COULD RESUME IN NOTHING
		HAPPENING

		*/
		for _, curDice := range m.Dices {
			if diceFace, msg, err = curDice.RollDice(); err != nil {
				fmt.Println(" error = ", err.Error())
				m.isRunning = false
				break
			}
		}
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
		if err, msg = m.board.MovePlayer(curPlayer.GetId(), int(diceFace)); err != nil {
			fmt.Println(" error = ", err.Error())
			m.isRunning = false
			break
		}
		fmt.Printf("player %s %s", curPlayer.GetName(), msg)
		time.Sleep(1 * time.Second)
		if isWinning, err = m.board.IsWinning(curPlayer.GetId()); err != nil {
			fmt.Println(" error = ", err.Error())
			m.isRunning = false
			break
		} else if isWinning {
			fmt.Printf(" player = %s won\n", curPlayer.GetName())
			m.playerRanks = append(m.playerRanks, curPlayer)

			/*IMPROVEMENT
			#OPTIONAL REQUIREMENT 3 : DONE
			*/
			if !m.isMultiWinner || (m.isMultiWinner && m.players.GetLen() == 1) {
				m.isRunning = false
				break
			}
			continue

		}
		m.players.Enqueue(curPlayer)
	}
	return err
}
