package snakeandladderboard

import (
	"hexagnal/internals/core/domains/elements"
	"hexagnal/internals/core/domains/player"
)

type Board interface {
	AddPlayer(players []*player.Player) error
	IsWinningMove(player *player.Player) bool
	UpdPos(player *player.Player, moves int) error
	AddSnakes(snakes []elements.TwoEdgeElement) error
	AddLadders(ladder []elements.TwoEdgeElement) error
	GetPlayerPosition(player *player.Player) (int, error)
}
