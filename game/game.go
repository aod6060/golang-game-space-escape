package game

import sdl "github.com/veandco/go-sdl2/sdl"

import "github.com/aod6060/golang-game-space-escape/engine/vmath"
import "github.com/aod6060/golang-game-space-escape/engine/render"

type GameApp struct {
	// I'll Add stuff here in abit
}

func (this *GameApp) Init() {

}

func (this *GameApp) HandleEvent(e sdl.Event) {

}

func (this *GameApp) Update(delta float32) {

}

func (this *GameApp) Render() {
	render.Clear(vmath.CreateVec4(1.0, 0.0, 0.0, 1.0))

	render.Bind()
	
	render.Unbind()
}

func (this *GameApp) Release() {

}
