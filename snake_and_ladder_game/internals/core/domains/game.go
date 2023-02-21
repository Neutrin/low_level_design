package domains

import (
	"fmt"
	"hexagnal/internals/core/domains/dice"
	"hexagnal/internals/core/domains/player"
	snakeandladderboard "hexagnal/internals/core/domains/snakeAndLadderBoard"
)

type Game struct {
	board     snakeandladderboard.Board
	players   []*player.Player
	dice      dice.Dice
	isRunning bool
}

func GetNewGame(board snakeandladderboard.Board, players []*player.Player, dice dice.Dice) *Game {
	return &Game{
		board:     board,
		players:   players,
		dice:      dice,
		isRunning: true,
	}
}

func PrintMove(player *player.Player, startPos int, endPos int) {
	fmt.Println(fmt.Sprintf("Player %s moved from %d to %d", player.GetName(), startPos, endPos))
}

func PrintWinner(player *player.Player) {
	fmt.Println(fmt.Sprintf("Player %s won the game ", player.GetName()))
}

func (game *Game) Play() error {
	var (
		err               error
		curPlayer         *player.Player
		diceCount, curPos int
	)
	for {
		curPlayer = game.players[0]
		diceCount = game.dice.Throw()
		//fmt.Printf("here the dice count is = %d\n", diceCount)
		if curPos, err = game.board.GetPlayerPosition(curPlayer); err != nil {
			return err
		}
		//fmt.Printf(" cur pos = %d for player = %s with id = %d", curPos, curPlayer.GetName(), curPlayer.GetId())
		game.board.UpdPos(curPlayer, diceCount)
		PrintMove(curPlayer, curPos, curPos+diceCount)
		if game.board.IsWinningMove(curPlayer) {
			PrintWinner(curPlayer)
			break
		}
		game.players = game.players[1:]
		game.players = append(game.players, curPlayer)
	}
	return err
}
