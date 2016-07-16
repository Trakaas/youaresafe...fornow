package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v3.3-compatibility/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
	runtime.LockOSThread()
}

func animate(y, h, w float32) (float32, float32) {
	//x1 y1 x2 y2
	gl.Rectf(w/2-5,
		(h/2-25)+y,
		w/2+5,
		(h/2+25)+y)
	gl.Color4f(1.0, 1.0, 1.0, 1.0)

	return (h/2 - 25) + y, (h/2 + 25) + y
}

func bounce(top, bottom, height float32) float32 {
	if top < 0 || bottom > height {
		return -1.0
	}
	return 1.0
}

func main() {

	err := glfw.Init() // Initialize GLFW
	if err != nil {    // check errors
		fmt.Println(err)
		fmt.Println("a")
		return
	}
	defer glfw.Terminate() //terminate at end of main

	// Initialize the Window
	window, err := glfw.CreateWindow(1600, 900, "testing Window", nil, nil)
	if err != nil { // check for errors
		fmt.Println(err)
		return
	}

	glfw.WindowHint(glfw.Resizable, gl.FALSE)    // Force no resize
	glfw.WindowHint(glfw.ContextVersionMajor, 3) // set GLFW major version
	glfw.WindowHint(glfw.ContextVersionMinor, 1) // set GLFW minor version
	window.MakeContextCurrent()                  // set all calls in main to bind to the "window"

	if err := gl.Init(); err != nil { // Initialize GL and check for errors
		fmt.Println(err)
		return
	}

	//	gl.Enable(gl.TEXTURE_2D)                           // set 2d texture mapping
	gl.Enable(gl.BLEND)                                // set alpha blending on
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA) // set alpha blending type
	gl.ClearColor(0, 0, 0, 1)                          // setting background alpha

	// Define Callbacks to event functions
	window.SetKeyCallback(func(window *glfw.Window, key glfw.Key, scancode int,
		action glfw.Action, mods glfw.ModifierKey) {
		fmt.Printf("keypress: key %v, scancode %v, action %v, mods %v\n", key,
			scancode, action, mods)
	},
	)

	window.SetMouseButtonCallback(
		func(window *glfw.Window, button glfw.MouseButton, action glfw.Action,
			mods glfw.ModifierKey) {
			fmt.Printf("button: button %v, action %v, mods %v\n", button,
				action, mods)
		},
	)

	window.SetCursorPosCallback(
		func(window *glfw.Window, xpos, ypos float64) {
			fmt.Printf("position: %.2f %.2f\n", xpos, ypos)
		},
	)

	w1, h1 := window.GetFramebufferSize()
	w, h := float32(w1), float32(h1)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(w), float64(h), 0, 1, -1)
	gl.MatrixMode(gl.MODELVIEW)

	gl.Color4f(1.0, 1.0, 1.0, 1.0)
	var yoff float32 = 1.0

	// for !window.ShouldClose() {
	// 	gl.Clear(gl.COLOR_BUFFER_BIT)
	// 	top := (h/2 - 25) + yoff
	// 	bottom := (h/2 + 25) + yoff
	// 	dir := bounce(top, bottom, h)
	//
	// 	gl.Rectf(w/2-5,
	// 		(h/2-25)+(yoff*dir),
	// 		w/2+5,
	// 		(h/2+25)+(yoff*dir))
	//
	// 	window.SwapBuffers()
	// 	glfw.PollEvents()
	// }
	var dir float32 = 1.0
	for !window.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT)

		w, h := float32(w), float32(h)

		gl.Rectf(
			w/2-15,
			h/2-25+yoff,
			w/2+15,
			h/2+25+yoff,
		)

		yoff += (0.25 * dir)
		dir = bounce(h/2-25+yoff, h/2+25+yoff, h)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
