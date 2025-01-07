package render

import "fmt"
import "unsafe"
import "github.com/go-gl/gl/v4.1-compatibility/gl"
import "github.com/aod6060/golang-game-space-escape/engine/vmath"

// Shaders
var mainVertexShader uint32 = 0
var mainFragmentShader uint32 = 0

// Program
var mainProgram uint32 = 0

// VertexArray
var mainVertexArray uint32 = 0

// Uniforms
var uProj int32 = 0
var uView int32 = 0
var uModel int32 = 0

// Attributes
const A_VERTICES uint32 = 0

type vertexBuffer struct {
	list []float32
	id uint32	
}

var vertices vertexBuffer

var vertCent vertexBuffer

// Uniforms
func Init() {
	gl.Disable(gl.DEPTH_TEST)

	//gl.CreateShader(gl.VERTEX_SHADER)
	//gl.GetUniformLocation(0, strconv.ParseUint("proj"))
	// Yeah! Its works :)
	//gl.GetUniformLocation(0, toGLString("proj"))

	var err error

	mainVertexShader, err = createShader(gl.VERTEX_SHADER, "data/shaders/main.vert.glsl")

	if err != nil {
		//panic(err)
		fmt.Println(err)
	}

	mainFragmentShader, err = createShader(gl.FRAGMENT_SHADER, "data/shaders/main.frag.glsl")

	if err != nil {
		//panic(err)
		fmt.Println(err)
	}


	mainProgram, err = createProgram(mainVertexShader, mainFragmentShader)

	if err != nil {
		//panic(err)
		fmt.Println(err)
	}

	gl.GenVertexArrays(1, &mainVertexArray)

	Bind()
	// Uniforms
	uProj = gl.GetUniformLocation(mainProgram, toGLString("proj"))
	uView = gl.GetUniformLocation(mainProgram, toGLString("view"))
	uModel = gl.GetUniformLocation(mainProgram, toGLString("model"))

	fmt.Println(uProj, uView, uModel)

	// Attributes
	gl.BindVertexArray(mainVertexArray)
	gl.EnableVertexAttribArray(A_VERTICES)
	gl.BindVertexArray(0)

	Unbind()


	vbInit(&vertices)

	vbAdd3f(&vertices, 0.0, 0.0, 0.0)
	vbAdd3f(&vertices, 1.0, 0.0, 0.0)
	vbAdd3f(&vertices, 0.0, 1.0, 0.0)

	vbAdd3f(&vertices, 0.0, 1.0, 0.0)
	vbAdd3f(&vertices, 1.0, 0.0, 0.0)
	vbAdd3f(&vertices, 1.0, 1.0, 0.0)

	vbUpdate(&vertices)

	fmt.Println("Size: ", len(vertices.list))

	vbInit(&vertCent)

	vbAdd3f(&vertCent, -0.5, -0.5, 0.0)
	vbAdd3f(&vertCent, 0.5, -0.5, 0.0)
	vbAdd3f(&vertCent, -0.5, 0.5, 0.0)

	vbAdd3f(&vertCent, -0.5, 0.5, 0.0)
	vbAdd3f(&vertCent, 0.5, -0.5, 0.0)
	vbAdd3f(&vertCent, 0.5, 0.5, 0.0)

	vbUpdate(&vertCent)
}

func Release() {
	vbRelease(&vertCent)
	vbRelease(&vertices)

	gl.DeleteVertexArrays(1, &mainVertexArray)
	deleteProgram(mainProgram, mainVertexShader, mainFragmentShader)
	gl.DeleteShader(mainFragmentShader)
	gl.DeleteShader(mainVertexShader)
}

func Clear(color vmath.Vec4) {
	gl.ClearColor(color.X, color.Y, color.Z, color.W)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func Bind() {
	gl.UseProgram(mainProgram)
}

func Unbind() {
	gl.UseProgram(0)
}

func SetProjection(m *vmath.Mat4) {
	var cm []float32 = vmath.Mat4ToArray(m)
	gl.UniformMatrix4fv(uProj, 1, false, &cm[0])
}

func SetView(m *vmath.Mat4) {
	var cm []float32 = vmath.Mat4ToArray(m)
	gl.UniformMatrix4fv(uView, 1, false, &cm[0])
}

func SetModel(m *vmath.Mat4) {
	var cm []float32 = vmath.Mat4ToArray(m)
	gl.UniformMatrix4fv(uModel, 1, false, &cm[0])
}

func Draw() {
	gl.BindVertexArray(mainVertexArray)

	vbBind(&vertices)
	gl.VertexAttribPointer(A_VERTICES, 3, gl.FLOAT, false, 0, nil)
	vbUnbind(&vertices)

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices.list) / 3))

	gl.BindVertexArray(0)
}

func DrawCenter() {
	gl.BindVertexArray(mainVertexArray)

	vbBind(&vertCent)
	gl.VertexAttribPointer(A_VERTICES, 3, gl.FLOAT, false, 0, nil)
	vbUnbind(&vertCent)

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices.list) / 3))

	gl.BindVertexArray(0)
}

// VertexBuffer
func vbInit(vb *vertexBuffer) {
	gl.GenBuffers(1, &vb.id)
}

func vbRelease(vb *vertexBuffer) {
	gl.DeleteBuffers(1, &vb.id)
}

func vbBind(vb *vertexBuffer) {
	gl.BindBuffer(gl.ARRAY_BUFFER, vb.id)
}

func vbUnbind(vb *vertexBuffer) {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func vbUpdate(vb *vertexBuffer) {
	vbBind(vb)
	gl.BufferData(gl.ARRAY_BUFFER, len(vb.list) * int(unsafe.Sizeof(float32(0))), gl.Ptr(vb.list), gl.DYNAMIC_DRAW)
	vbUnbind(vb)
}

func vbClear(vb *vertexBuffer) {
	vb.list = vb.list[:0]
}

func vbAdd1f(vb *vertexBuffer, x float32) {
	vb.list = append(vb.list, x)
}

func vbAdd2f(vb *vertexBuffer, x float32, y float32) {
	vbAdd1f(vb, x)
	vbAdd1f(vb, y)
}

func vbAdd3f(vb *vertexBuffer, x float32, y float32, z float32) {
	vbAdd1f(vb, x)
	vbAdd1f(vb, y)
	vbAdd1f(vb, z)
}

func vbAdd4f(vb *vertexBuffer, x float32, y float32, z float32, w float32) {
	vbAdd1f(vb, x)
	vbAdd1f(vb, y)
	vbAdd1f(vb, z)
	vbAdd1f(vb, w)	
}
