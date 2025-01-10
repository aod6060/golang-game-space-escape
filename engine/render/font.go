package render


import "github.com/go-gl/gl/v4.1-compatibility/gl"
import ttf "github.com/veandco/go-sdl2/ttf"
import sdl "github.com/veandco/go-sdl2/sdl"

type Font struct {
	font *ttf.Font
	id uint32
	width int32
	height int32
}

var fonts map[string]Font

func FontInit() {
	ttf.Init()
	fonts = make(map[string]Font)
}

func FontRelease() {
	for _, val := range fonts {
		gl.DeleteTextures(1, &val.id)
		val.font.Close()
	}
	ttf.Quit()
}

func FontCreateFont(name string, path string, size int) {
	var temp Font
	var err error

	temp.font, err = ttf.OpenFont(path, size)

	if err != nil {
		panic(err)
	}

	gl.GenTextures(1, &temp.id)

	fonts[name] = temp
}

func FontGetSize(name string, text string) (int, int) {
	var w, h int
	var err error
	w, h, err = fonts[name].font.SizeUTF8(text)
	if err != nil {
		panic(err)
	}
	return w, h
}

func FontDraw(name string, text string) {
	var color sdl.Color
	color.R = 255
	color.G = 255
	color.B = 255
	color.A = 255

	var surf *sdl.Surface
	var err error

	surf, err = fonts[name].font.RenderUTF8Blended(text, color)

	if err != nil {
		panic(err)
	}

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, fonts[name].id)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		surf.W,
		surf.H,
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(surf.Pixels()))

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	
	Draw()

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, 0)

	surf.Free()
}
