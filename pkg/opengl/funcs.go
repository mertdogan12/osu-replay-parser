package opengl

import (
	"errors"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// Check for error during the compiling
func getError(programm uint32, programmitiv bool) error {
	var status int32

	if programmitiv {
		gl.GetProgramiv(programm, gl.LINK_STATUS, &status)
	} else {
		gl.GetShaderiv(programm, gl.COMPILE_STATUS, &status)
	}

	if status == gl.FALSE {
		var logLenght int32
		gl.GetShaderiv(programm, gl.INFO_LOG_LENGTH, &logLenght)
		log := string(make([]byte, logLenght+1))
		gl.GetShaderInfoLog(programm, logLenght, nil, gl.Str(log))

		return errors.New(log)
	}

	return nil
}
