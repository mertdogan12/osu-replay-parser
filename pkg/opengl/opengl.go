package opengl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func Init(winWidth int, winHeight int) {
	// Initialisates glfw
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	defer glfw.Terminate()

	window, err := glfw.CreateWindow(winWidth, winHeight, "OpenGL", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	defer window.Destroy()

	// OpenGL init
	err = gl.Init()
	if err != nil {
		panic(err)
	}

	// Vertex shader
	vertexShader, err := compileShader("pkg/opengl/shaders/default.vert", gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	// Fragment shader
	fragmentShader, err := compileShader("pkg/opengl/shaders/default.frag", gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	// Link shaders in a programm
	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)

	err = getError(shaderProgram, true)
	if err != nil {
		panic(err)
	}

	// Delete shaders
	gl.DeleteShader(fragmentShader)
	gl.DeleteShader(vertexShader)

	// Vertices
	vertices := []float32{
		0.5, 0.5, 0.0,
		0.5, -0.5, 0.0,
		-0.5, 0.5, 0.0,
		-0.5, -0.5, 0.0}

	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)

	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.BindVertexArray(0)

	for !window.ShouldClose() {
		gl.ClearColor(0.0, 0.0, 0.0, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.UseProgram(shaderProgram)
		gl.BindVertexArray(VAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
