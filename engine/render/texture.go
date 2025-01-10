package render


import "github.com/go-gl/gl/v4.1-compatibility/gl"
import img "github.com/veandco/go-sdl2/img"
import sdl "github.com/veandco/go-sdl2/sdl"

type Texture2D struct {
	Id uint32
	Width int32
	Height int32
}


var textures map[string]Texture2D

func TextureInit() {
	textures = make(map[string]Texture2D)
}

func TextureCreateFromFile(name string, path string) {
	var surf *sdl.Surface
	var err error

	surf, err = img.Load(path)

	if err != nil {
		panic(err)
	}

	var temp Texture2D

	temp.Width = surf.W
	temp.Height = surf.H

	gl.GenTextures(1, &temp.Id)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, temp.Id)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		temp.Width,
		temp.Height,
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(surf.Pixels()))

	
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, 0)

	surf.Free()

	textures[name] = temp
}

func TextureBind(name string, text uint32) {
	gl.ActiveTexture(text)
	gl.BindTexture(gl.TEXTURE_2D, textures[name].Id)

}

func TextureUnbind(name string, text uint32) {
	gl.ActiveTexture(text)
	gl.BindTexture(gl.TEXTURE_2D, textures[name].Id)
}

func TextureRelease() {
	for _, val := range textures {
		gl.DeleteTextures(1, &val.Id)
	}
}