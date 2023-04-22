package src

import (
	"fmt"
	"time"

	"github.com/neutrin/snake_and_ladder/internals/src/boards"
	"github.com/neutrin/snake_and_ladder/internals/src/dices"
	"github.com/neutrin/snake_and_ladder/internals/src/players"
	"github.com/neutrin/snake_and_ladder/internals/src/rules_game"
)

type GameManager struct {
	board       boards.Board
	players     []*players.Player
	playerRanks []*players.Player
	isRunning   bool
	Dice        dices.Dice
	rules       rules_game.GameRules
}

func NewGameManager(board boards.Board, players []*players.Player, dice dices.Dice,
	rules rules_game.GameRules) *GameManager {
	return &GameManager{
		board:   board,
		players: players,
		//playerRanks: curPlayerList,
		isRunning: true,
		Dice:      dice,
		rules:     rules,
	}

}

func (m *GameManager) Play() error {
	var (
		err           error
		playerPointer int
		curPlayer     *players.Player
		diceFace      int64
		msg           string
		isWinning     bool
	)
	if !m.isRunning {
		return fmt.Errorf("game not running")
	}
	if len(m.players) == 0 {
		return fmt.Errorf("game empty")
	}
	m.playerRanks = make([]*players.Player, 0)

	for m.isRunning {
		curPlayer = m.players[playerPointer]

		if diceFace, msg, err = m.rules.RollDice(m.Dice); err != nil {
			fmt.Println(" error = ", err.Error())
			m.isRunning = false
			break
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
			m.isRunning = false
			break
		}
		playerPointer = (playerPointer + 1) % len(m.players)
	}
	return err
}
