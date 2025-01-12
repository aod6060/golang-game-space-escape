package game

/*
import "math"
import "github.com/go-gl/gl/v4.1-compatibility/gl"
import "github.com/aod6060/golang-game-space-escape/engine/vmath"
import "github.com/aod6060/golang-game-space-escape/engine/input"
import "github.com/aod6060/golang-game-space-escape/engine/render"


type Player struct {
	pos vmath.Vec3
	rot float32
	size vmath.Vec3
}

func PlayerInit(player *Player) {
	player.pos = vmath.Vec3Create(0.0, 0.0, 0.0)
	player.rot = 0.0
	player.size = vmath.Vec3Create(32.0, 32.0, 0.0)
}

func PlayerUpdate(player* Player, delta float32) {
	var rd = input.GetKeyPressedAxis(input.KEYS_LEFT, input.KEYS_RIGHT)
	player.rot += rd * 32.0 * delta

	//var move = input.GetKeyPressedAxis(input.KEYS_DOWN, input.KEYS_UP)

	var rad = vmath.ToRadian(player.rot)

	if(input.IsKeyPressed(input.KEYS_UP)) {
		player.pos.X -= float32(math.Sin(float64(rad)))
		player.pos.Y += float32(math.Cos(float64(rad)))
	}
}

func PlayerRender(player *Player) {
	var t, r, s, tr, model vmath.Mat4

	t = vmath.TransformTranslate(&player.pos)
	r = vmath.TransformRotationZ(vmath.ToRadian(player.rot))
	s = vmath.TransformScale(&player.size)

	tr = vmath.Mat4MulMatrix(&r, &t)
	model = vmath.Mat4MulMatrix(&s, &tr)

	render.SetModel(&model)

	render.TextureBind("player", gl.TEXTURE0)
	render.DrawCenter()
	render.TextureUnbind("player", gl.TEXTURE0)
}

func PlayerRelease(player *Player) {

}
*/