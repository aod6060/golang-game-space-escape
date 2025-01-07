package render

type Texture2D struct {
	Id uint32
	Width uint32
	Height uint32
}


var textures map[string]Texture2D

func TextureCreateFromFile(name string, path string) {

}

func TextureBind(name string, text uint32) {

}

func TextureUnbind(name string, text uint32) {

}

func TextureRelease() {
	
}