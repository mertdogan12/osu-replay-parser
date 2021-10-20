package opengl

import (
	"os"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// Compiles an shader from an a shader file.
// Then returns the shader
func compileShader(shaderFile string, shaderType uint32) (uint32, error) {
	shaderSource, err := os.ReadFile(shaderFile)
	if err != nil {
		return 0, err
	}

	shader := gl.CreateShader(shaderType)
	csourse, free := gl.Strs(string(shaderSource))
	gl.ShaderSource(shader, 1, csourse, nil)
	free()
	gl.CompileShader(shader)

	err = getError(shader)
	if err != nil {
		return 0, err
	}

	return shader, nil
}
