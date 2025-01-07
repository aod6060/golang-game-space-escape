package app



import sdl "github.com/veandco/go-sdl2/sdl"
import "github.com/go-gl/gl/v4.1-compatibility/gl"

type Config struct {
	Caption string
	Width int32
	Height int32

	InitCB func()
	HandleEventCB func(e sdl.Event)
	UpdateCB func(delta float32)
	RenderCB func()
	ReleaseCB func()
}


var conf *Config
var isRunning bool = true
var window *sdl.Window = nil
var context sdl.GLContext


func Init(_conf *Config) {
	conf = _conf

	var err error

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic("Failed to initialize SDL")
	}

	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 4)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1)
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)

	window, err = sdl.CreateWindow(
		conf.Caption,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		conf.Width,
		conf.Height,
		sdl.WINDOW_OPENGL)

	if err != nil {
		panic(err)
	}

	context, err = window.GLCreateContext()
	if err != nil {
		panic(err)
	}

	gl.Init()

	conf.InitCB()
}

func Update() {
	var event sdl.Event
	var preTime uint32 = sdl.GetTicks()
	var currTime uint32
	var delta float32

	for isRunning {
		currTime = sdl.GetTicks()
		delta = (float32(currTime) - float32(preTime)) / 1000.0
		preTime = currTime

		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			if event.GetType() == sdl.QUIT {
				Exit()
			}
			conf.HandleEventCB(event)
		}

		conf.UpdateCB(delta)
		conf.RenderCB()

		window.GLSwap()
	}
}

func Release() {
	conf.ReleaseCB()

	sdl.GLDeleteContext(context)
	window.Destroy()
	sdl.Quit()
}

func GetWidth() int32 {
	return conf.Width
}

func GetHeight() int32 {
	return conf.Height
}

func GetAspect() float32 {
	return float32(GetWidth()) / float32(GetHeight())
}

func Exit() {
	isRunning = false
}

