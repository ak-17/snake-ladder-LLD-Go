package main

import engine2 "github.com/ak-17/snake-ladder-lld/engine"

func main() {

	gameEngine := engine2.InitEngine(10, 10, 200)
	gameEngine.AddPlayer("Akshay")
	gameEngine.AddPlayer("Sasuke")
	gameEngine.Play()

}
