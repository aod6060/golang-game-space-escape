package main

//import "fmt"
//import "github.com/aod6060/golang-game-space-escape/game"
//import "github.com/aod6060/golang-game-space-escape/game"

//import "github.com/aod6060/golang-game-space-escape/engine/vmath"


import "github.com/aod6060/golang-game-space-escape/engine/app"
import "github.com/aod6060/golang-game-space-escape/game"

func main() {
	var config app.Config

	game.Setup(&config)

	app.Init(&config)
	app.Update()
	app.Release()
}


