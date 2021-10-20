package opengl

import (
	"errors"

	"github.com/go-gl/gl/v2.1/gl"
)

// Check for error during the compiling
func getError(programm uint32) error {
	var status int32
	gl.GetShaderiv(programm, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var logLenght int32
		gl.GetShaderiv(programm, gl.INFO_LOG_LENGTH, &logLenght)
		log := string(make([]byte, logLenght+1))
		gl.GetShaderInfoLog(programm, logLenght, nil, gl.Str(log))

		return errors.New("Failed to compile shader: \n" + log)
	}

	return nil
}
