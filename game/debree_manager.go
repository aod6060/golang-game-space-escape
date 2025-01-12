package game

/*
import "github.com/go-gl/gl/v4.1-compatibility/gl"
import "github.com/aod6060/golang-game-space-escape/engine/vmath"
import "github.com/aod6060/golang-game-space-escape/engine/render"

type DebreeType int32
const (
	DT_DEBREE_01 DebreeType = iota
	DT_DEBREE_02
	DT_DEBREE_03
	DT_DEBREE_04
	DT_DEBREE_05
	DT_DEBREE_MAX_SIZE
)

type Debree struct {
	debreeType DebreeType
	pos vmath.Vec3
	rot float32
	size vmath.Vec3
}

var debreeList []Debree


func DebreeInit(amount int) {
	for i := 0; i < amount; i++ {
		var temp Debree

		var x, y, r float32

		x = MRand().Float32() * 1024.0
		y = MRand().Float32() * 1024.0

		r = MRand().Float32() * 360.0

		temp.pos = vmath.Vec3Create(x, y, 0.0)
		temp.rot = r
		temp.size = vmath.Vec3Create(64.0, 64.0, 0.0)

		temp.debreeType = DebreeType(MRand().Int31n(int32(DT_DEBREE_MAX_SIZE)))

		debreeList = append(debreeList, temp)
	}
}

func DebreeUpdate(delta float32) {
	// I'll implement this later
}

func toTexture(t DebreeType) string {
	var tex string = "debree1"

	if t == DT_DEBREE_01 {
		tex = "debree1"
	} else if t == DT_DEBREE_02 {
		tex = "debree2"
	} else if t == DT_DEBREE_03 {
		tex = "debree3"
	} else if t == DT_DEBREE_04 {
		tex = "debree4"
	} else if t == DT_DEBREE_05 {
		tex = "debree5"
	} else {
		tex = "debree1"
	}

	return tex
}

func DebreeRender() {
	for i := 0; i < len(debreeList); i++ {
		var tex string = toTexture(debreeList[i].debreeType)

		var t, r, s, tr, model vmath.Mat4

		t = vmath.TransformTranslate(&debreeList[i].pos)
		r = vmath.TransformRotationZ(vmath.ToRadian(debreeList[i].rot))
		s = vmath.TransformScale(&debreeList[i].size)

		tr = vmath.Mat4MulMatrix(&r, &t)
		model = vmath.Mat4MulMatrix(&s, &tr)

		render.SetModel(&model)

		render.TextureBind(tex, gl.TEXTURE0)
		render.DrawCenter()
		render.TextureUnbind(tex, gl.TEXTURE0)
	}
}

func DebreeRelease() {
	debreeList = debreeList[:0]
}
*/