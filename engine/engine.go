package engine

import (
	"fmt"
	"github.com/ak-17/snake-ladder-lld/model"
)

func (engine *engine) AddPlayer(name string) {
	p := model.InitPlayer(name)
	engine.players = append(engine.players, p)
}

func (engine *engine) Play()  {
	if len(engine.players) == 0 {
		fmt.Printf("no players joined.")
		return
	}

	for {
		currentPlayer := engine.players[0]
		roll := engine.dice.Roll()

		newPosition := currentPlayer.Position + roll
		if newPosition > engine.board.GetEndValue() {
			fmt.Printf("Player: %s, diceRoll: %d\n", currentPlayer, roll)
			engine.players = append(engine.players[1:],engine.players[0])
			continue
		}

		currentPlayer.Position = engine.board.GetNewPosition(currentPlayer.Position + roll)
		fmt.Printf("Player: %s, diceRoll: %d\n", currentPlayer, roll)
		if currentPlayer.Position == engine.board.GetEndValue() {
			currentPlayer.Won = true
			fmt.Printf("%s Won!!!\n",currentPlayer)
			engine.players = engine.players[1:]
		} else {
			engine.players = append(engine.players[1:],engine.players[0])
		}

		if len(engine.players) < 2 {
			fmt.Printf("Game finished!!\n")
			return
		}
	}
}