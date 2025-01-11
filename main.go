package main

import "github.com/aod6060/golang-game-space-escape/engine/app"
import "github.com/aod6060/golang-game-space-escape/game"

func main() {
	var config app.Config

	game.Setup(&config)

	app.Init(&config)
	app.Update()
	app.Release()
}


