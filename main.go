package main

import "github.com/ak-17/snake-ladder-lld/engine"

func main() {

	gameEngine := engine.InitEngine(10, 10, 200)
	gameEngine.AddPlayer("Naruto")
	gameEngine.AddPlayer("Sasuke")
	gameEngine.AddPlayer("Sakura")
	gameEngine.AddPlayer("Negi")
	gameEngine.Play()

}
