package engine

import "github.com/ak-17/snake-ladder-lld/model"

type engine struct {
	board *model.Board
	dice *model.Dice
	players []*model.Player
}
