package main

import "github.com/aod6060/golang-game-space-escape/engine/app"
import "github.com/aod6060/golang-game-space-escape/game"

func main() {
	/*
	var config app.Config

	game.Setup(&config)

	app.Init(&config)
	app.Update()
	app.Release()
	*/

	// Redoing the vmath library

	var gameApp game.GameApp
	var config app.Config

	config.Caption = "Space Escape"
	config.Width = 640
	config.Height = 480
	config.App = &gameApp

	app.Init(&config)
	app.Update()
	app.Release()
	
}


