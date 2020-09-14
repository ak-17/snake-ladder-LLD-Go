package engine

import "github.com/ak-17/snake-ladder-lld/model"

func InitEngine(numberOfSnakes, numberOfLadders, boardSize int64) *engine {
	return &engine{
		board:   model.InitBoard(0,boardSize,numberOfSnakes,numberOfLadders),
		dice:    model.InitDice(6),
		players: []*model.Player{},
	}
}

