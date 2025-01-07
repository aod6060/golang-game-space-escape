package game


//import "fmt"
import sdl "github.com/veandco/go-sdl2/sdl"
//import "github.com/go-gl/gl/v4.1-compatibility/gl"
import "github.com/aod6060/golang-game-space-escape/engine/vmath"
import "github.com/aod6060/golang-game-space-escape/engine/app"
import "github.com/aod6060/golang-game-space-escape/engine/input"
import "github.com/aod6060/golang-game-space-escape/engine/render"

//import "github.com/aod6060/golang-game-space-escape/engine/input"


var count int32 = 0


var pos vmath.Vec3
var speed float32 = 128.0

func Init() {
	pos = vmath.Vec3Create(32.0, 32.0, 0.0)

}

func HandleEvent(e sdl.Event) {
}

func Update(delta float32) {

	var dir vmath.Vec3 = vmath.Vec3Create(
		input.GetKeyPressedAxis(input.KEYS_LEFT, input.KEYS_RIGHT),
		input.GetKeyPressedAxis(input.KEYS_UP, input.KEYS_DOWN),
		0.0)

	var vel = vmath.Vec3Mul(&dir, (delta * speed))

	pos = vmath.Vec3Add(&pos, &vel)
}

func Render() {
	render.Clear(vmath.Vec4Create(1.0, 0.0, 0.0, 1.0))

	render.Bind()

	var proj, view, model vmath.Mat4

	proj = vmath.TransformOrtho2D(0.0, float32(app.GetWidth()), float32(app.GetHeight()), 0.0)
	render.SetProjection(&proj)

	view = vmath.Mat4Identity()
	render.SetView(&view)

	var t, s vmath.Mat4
	var v vmath.Vec3

	v = vmath.Vec3Create(32.0, 32.0, 0.0)
	t = vmath.TransformTranslate(&pos)
	s = vmath.TransformScale(&v)

	model = vmath.Mat4MulMatrix(&s, &t)

	render.SetModel(&model)

	render.DrawCenter()

	render.Unbind()

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