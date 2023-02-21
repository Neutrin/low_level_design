package main

import (
	"bufio"
	"fmt"
	"hexagnal/internals/core/domains"
	"hexagnal/internals/core/domains/dice"
	"hexagnal/internals/core/domains/elements"
	"hexagnal/internals/core/domains/player"
	snakeandladderboard "hexagnal/internals/core/domains/snakeAndLadderBoard"
	"os"
	"strconv"
	"strings"
)

func getIntFromString(str string) []int {
	var output = make([]int, 0)
	str = strings.Replace(str, "\n", "", -1)
	strSplit := strings.Split(str, " ")
	for _, curSplit := range strSplit {
		intVal, _ := strconv.Atoi(curSplit)
		output = append(output, intVal)
	}
	return output
}

func getStringFromString(str string) []string {
	var output = make([]string, 0)
	str = strings.Replace(str, "\n", "", -1)
	strSplit := strings.Split(str, " ")
	for _, curSplit := range strSplit {
		output = append(output, curSplit)
	}
	return output
}
func main() {
	var (
		err                                          error
		str                                          string
		snakeCount, ladderCount, playerCount         int
		snakeList                                    = make([]elements.TwoEdgeElement, 0)
		ladderList                                   = make([]elements.TwoEdgeElement, 0)
		curSnake                                     *elements.Snake
		curLadder                                    *elements.Ladder
		playerList                                   = make([]*player.Player, 0)
		snakeStart, snakeEnd, ladderStart, ladderEnd int
	)
	fmt.Println("please enter input")
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	snakeCount = getIntFromString(str)[0]
	for count := 0; count < snakeCount; count++ {
		str, _ = reader.ReadString('\n')
		snakeStart = getIntFromString(str)[0]
		snakeEnd = getIntFromString(str)[1]
		curSnake = elements.GetSnake(snakeStart, snakeEnd)
		if err = curSnake.IsValid(); err != nil {
			fmt.Printf("error = %+v\n ", err)
		}
		snakeList = append(snakeList, curSnake)
	}
	str, _ = reader.ReadString('\n')
	ladderCount = getIntFromString(str)[0]
	for count := 0; count < ladderCount; count++ {
		str, _ = reader.ReadString('\n')
		ladderStart = getIntFromString(str)[0]
		ladderEnd = getIntFromString(str)[1]
		curLadder = elements.GetNewLadder(ladderStart, ladderEnd)
		if err = curLadder.IsValid(); err != nil {
			fmt.Printf("error = %+v\n ", err)
		}
		ladderList = append(ladderList, curLadder)
	}
	str, _ = reader.ReadString('\n')
	playerCount = getIntFromString(str)[0]
	for count := 0; count < playerCount; count++ {
		str, _ = reader.ReadString('\n')
		playerList = append(playerList, player.GetNewPlayer(getStringFromString(str)[0],
			fmt.Sprintf("%s@gmail.com", getStringFromString(str)[0])))
	}
	board := snakeandladderboard.GetSquareBoard(10)
	if err = board.AddPlayer(playerList); err != nil {
		fmt.Printf(" error while adding players = %+v\n", err)
		return
	}
	if err = board.AddSnakes(snakeList); err != nil {
		fmt.Printf(" error while adding snakes = %+v\n", err)
		return
	}
	if err = board.AddLadders(ladderList); err != nil {
		fmt.Printf(" error while adding ladders = %+v\n", err)
		return

	}
	dice := dice.GetSingleDice(6)
	game := domains.GetNewGame(board, playerList, &dice)

	if err = game.Play(); err != nil {
		fmt.Printf(" error came while running game = %+v\n", err)
	}
	fmt.Println("game finished")

}
