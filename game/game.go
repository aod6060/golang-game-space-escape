package game

import sdl "github.com/veandco/go-sdl2/sdl"
import "github.com/go-gl/gl/v4.1-compatibility/gl"
import "github.com/aod6060/golang-game-space-escape/engine/vmath"
import "github.com/aod6060/golang-game-space-escape/engine/app"
//import "github.com/aod6060/golang-game-space-escape/engine/input"
import "github.com/aod6060/golang-game-space-escape/engine/render"

var player Player

func Init() {
	render.TextureCreateFromFile("player", "data/textures/player.png")
	render.TextureCreateFromFile("bg", "data/textures/background.png")

	render.FontCreateFont("font", "data/font/font.ttf", 32)

	PlayerInit(&player)
}

func HandleEvent(e sdl.Event) {
}

func Update(delta float32) {
	PlayerUpdate(&player, delta)
}

func DrawEntity(pos vmath.Vec3, size vmath.Vec3) {
	var t, s, model vmath.Mat4

	t = vmath.TransformTranslate(&pos)
	s = vmath.TransformScale(&size)

	model = vmath.Mat4MulMatrix(&s, &t)

	render.SetModel(&model)

	render.DrawCenter()
}

func DrawText(text string, pos vmath.Vec3) {
	var t, s, model vmath.Mat4

	t = vmath.TransformTranslate(&pos)

	var w, h int = render.FontGetSize("font", text)

	var size vmath.Vec3 = vmath.Vec3Create(float32(w), float32(h), 0.0)

	s = vmath.TransformScale(&size)

	model = vmath.Mat4MulMatrix(&s, &t)

	render.SetModel(&model)
	render.FontDraw("font", text)
}

func RenderBackgroundHelper(offset vmath.Vec3) {
	var s, t, model vmath.Mat4

	var screenSize = app.GetScreenSize()

	var pos = vmath.Vec3Add(&player.pos, &offset)
	t = vmath.TransformTranslate(&pos)

	s = vmath.TransformScale(&screenSize)

	model = vmath.Mat4MulMatrix(&s, &t)

	render.SetModel(&model)

	render.TextureBind("bg", gl.TEXTURE0)
	render.DrawCenter()
	render.TextureUnbind("bg", gl.TEXTURE0)
}

func RenderBackground() {
	RenderBackgroundHelper(vmath.Vec3Create(-320.0, -240.0, 0.0))
	RenderBackgroundHelper(vmath.Vec3Create(320.0, -240.0, 0.0))
	RenderBackgroundHelper(vmath.Vec3Create(-320.0, 240.0, 0.0))
	RenderBackgroundHelper(vmath.Vec3Create(320.0, 240.0, 0.0))
}

func HandleCamera() {
	var r, t, view vmath.Mat4

	//view = vmath.Mat4Identity()
	var screenSize vmath.Vec3 = app.GetScreenSize()
	var halfScreenSize vmath.Vec3 = vmath.Vec3Div(&screenSize, 2.0)
	var offset vmath.Vec3 = vmath.Vec3Sub(&halfScreenSize, &player.pos)

	t = vmath.TransformTranslate(&offset)
	r = vmath.TransformRotationZ(vmath.ToRadian(player.rot))

	view = vmath.Mat4MulMatrix(&r, &t)

	render.SetView(&view)
}

func Render() {
	render.Clear(vmath.Vec4Create(0.0, 0.0, 0.0, 1.0))

	render.Bind()

	var proj vmath.Mat4

	proj = vmath.TransformOrtho2D(0.0, float32(app.GetWidth()), float32(app.GetHeight()), 0.0)
	render.SetProjection(&proj)

	HandleCamera()

	RenderBackground()

	render.EnableBlend()

	PlayerRender(&player)

	render.DisableBlend()

	render.Unbind()

}

func Release() {
	PlayerRelease(&player)
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