package main

import (
	"fmt"

	"github.com/neutrin/snake_and_ladder/internals/src"
	"github.com/neutrin/snake_and_ladder/internals/src/boards"
	"github.com/neutrin/snake_and_ladder/internals/src/dice_service"
	"github.com/neutrin/snake_and_ladder/internals/src/dices"
	"github.com/neutrin/snake_and_ladder/internals/src/elements"
	"github.com/neutrin/snake_and_ladder/internals/src/players"
)

func main() {
	var (
		snakeCount  int
		ladderCount int
		snakeHead   int
		snakeTail   int
		ladderStart int
		ladderEnd   int
		playerCount int
		curSnake    *elements.Snake
		curLadder   *elements.Ladder
		curPlayer   *players.Player
		name        string
		snakes      []*elements.Snake
		ladders     []*elements.Ladder
		playersList []*players.Player
		err         error
		board       boards.Board
	)
	fmt.Println("enter number of snakes")
	fmt.Scan(&snakeCount)
	fmt.Println("enter space seperated lines for snakes")
	for index := 0; index < snakeCount; index++ {
		fmt.Scan(&snakeHead, &snakeTail)
		if curSnake, err = elements.NewSnake(snakeHead, snakeTail); err != nil {
			fmt.Println(" error in snake pls retry = ", err.Error())
			index--
		} else {
			snakes = append(snakes, curSnake)
		}
	}
	fmt.Println("snakes entered")
	fmt.Println("enter number of ladders")
	fmt.Scan(&ladderCount)
	fmt.Println("enter space seperated lines for ladders")
	for index := 0; index < ladderCount; index++ {
		fmt.Scan(&ladderStart, &ladderEnd)
		if curLadder, err = elements.NewLadder(ladderStart, ladderEnd); err != nil {
			fmt.Println(" error in ladder pls retry = ", err.Error())
			index--
		} else {
			ladders = append(ladders, curLadder)
		}
	}
	fmt.Println("enter number of players")
	fmt.Scan(&playerCount)
	fmt.Println("enter player names")
	for index := 0; index < playerCount; index++ {
		fmt.Scan(&name)
		curPlayer = players.NewPlayer(name, players.NewWidget("black"))
		playersList = append(playersList, curPlayer)
	}
	/*
		OPTIONAL REQUIREMENT #2 : DONE
	*/
	rows := 0
	cols := 0
	fmt.Println("please enter number of rows")
	fmt.Scan(&rows)
	fmt.Println("please enter number of cols")
	fmt.Scan(&cols)
	if board, err = boards.NewReactangularBoard(rows, cols, playersList); err != nil {
		fmt.Println("error in board intialization = %s", err.Error())
	}
	if err = board.AddLadders(ladders); err != nil {
		fmt.Println(" error in adding ladders = ", err.Error())
	}
	fmt.Println("ladders added")
	if err = board.AddSnakes(snakes); err != nil {
		fmt.Println(" error in adding snakes ", err.Error())
	}
	fmt.Println("snakes added")
	simpleDice := dices.NewSimpleDice(1, 6)
	gameManager := src.NewGameManager(board, playersList, []dice_service.DiceService{dice_service.NewSimpleDiceService(simpleDice)})
	gameManager.IsMultiWinner()
	if err = gameManager.Play(); err != nil {
		fmt.Println(" error while playing games = ", err.Error())
	}

}
