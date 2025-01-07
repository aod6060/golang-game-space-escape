package render

import "github.com/go-gl/gl/v4.1-compatibility/gl"

import "github.com/aod6060/golang-game-space-escape/engine/vmath"


//
func Init() {
	gl.Disable(gl.DEPTH_TEST)
}

func Release() {
}

func Clear(color vmath.Vec4) {
	gl.ClearColor(color.X, color.Y, color.Z, color.W)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func Bind() {
}

func Unbind() {
}

func SetProjection(m *vmath.Mat4) {

}

func SetView(m *vmath.Mat4) {

}

func SetModel(m *vmath.Mat4) {

}

func Draw() {

}