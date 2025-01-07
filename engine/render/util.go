package render

import "fmt"
import "strings"
import "io/ioutil"
import "github.com/go-gl/gl/v4.1-compatibility/gl"

func toGLString(str string) *uint8 {
	return gl.Str(str + "\x00")
}

func toGLStrings(str string) (**uint8, func()) {
	return gl.Strs(str + "\x00")
}

func getFileContents(path string) string {
	var content, err = ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func createShader(shaderType uint32, path string) (uint32, error) {
	var temp uint32 = gl.CreateShader(shaderType)

	var source = getFileContents(path)

	fmt.Println(source)

	//var c_source = toGLString(source)
	// forgot the "\x00"
	//c_source, free := gl.Strs(source + "\x00")
	c_source, free := toGLStrings(source)

	gl.ShaderSource(temp, 1, c_source, nil)
	free()

	gl.CompileShader(temp)

	var status int32
	gl.GetShaderiv(temp, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var len int32
		gl.GetShaderiv(temp, gl.INFO_LOG_LENGTH, &len)
		log := strings.Repeat("\x00", int(len+1))
		gl.GetShaderInfoLog(temp, len, nil, gl.Str(log))
		return 0, fmt.Errorf("Failed to compile %v: %v", source, log)
	}
	return temp, nil
}

func createProgram(shaders ...uint32) (uint32, error) {
	var temp uint32 = gl.CreateProgram()

	for _, shader := range shaders {
		gl.AttachShader(temp, shader)
	}

	gl.LinkProgram(temp)
	var status int32
	gl.GetProgramiv(temp, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var len int32
		gl.GetProgramiv(temp, gl.INFO_LOG_LENGTH, &len)
		log := strings.Repeat("\x00", int(len + 1))
		gl.GetProgramInfoLog(temp, len, nil, gl.Str(log))
		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	return temp, nil
}

func deleteProgram(program uint32, shaders ...uint32) {
	for _, shader := range shaders {
		gl.DetachShader(program, shader)
	}
	gl.DeleteProgram(program)
}
