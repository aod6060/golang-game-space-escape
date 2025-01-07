package game


import "fmt"
import "github.com/aod6060/golang-game-space-escape/engine/app"
import sdl "github.com/veandco/go-sdl2/sdl"
import "github.com/go-gl/gl/v4.1-compatibility/gl"

func Init() {

}

func HandleEvent(e sdl.Event) {
}

func Update(delta float32) {
	fmt.Println("Delta: ", delta)
}

func Render() {
	gl.ClearColor(1.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

}

func Release() {

}

func Setup(config *app.Config) {
	config.Caption = "Golang Space Escape"
	config.Width = 640
	config.Height = 480

	config.InitCB = Init
	config.HandleEventCB = HandleEvent
	config.UpdateCB = Update
	config.RenderCB = Render
	config.ReleaseCB = Release
}