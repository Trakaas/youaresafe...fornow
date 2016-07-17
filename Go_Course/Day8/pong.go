package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v3.3-compatibility/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

// Ball item
type Ball struct {
	x      float32
	y      float32
	width  float32
	height float32
	dx     float32
	dy     float32
}

// Paddle item
type Paddle struct {
	x      float32
	y      float32
	width  float32
	height float32
	dy     float32
}

func reactEvents(ball *Ball, paddleLeft *Paddle, paddleRight *Paddle, w, h int) bool {
	// collide with top or bottom
	if ball.y < 0 || ball.y+ball.height > float32(h) {
		ball.dy = -ball.dy
		return false
	}

	// collide with board edges beyond paddles
	if ball.x < 0 || ball.x+ball.width > float32(w) {
		return true
	}

	// collide with left paddleLeft long edge
	if ball.x-ball.width == paddleLeft.x+paddleLeft.width &&
		(ball.y-ball.height < paddleLeft.y+paddleLeft.height ||
			ball.y+ball.height > paddleLeft.y-paddleLeft.height) {

		ball.dy = -ball.dy
		ball.dx = -ball.dx
		return false
	}

	// collide with left paddleRight long edge
	if ball.x+ball.width == paddleRight.x-paddleRight.width &&
		(ball.y-ball.height < paddleRight.y+paddleRight.height ||
			ball.y+ball.height > paddleRight.y-paddleRight.height) {

		ball.dy = -ball.dy
		ball.dx = -ball.dx
		return false
	}
	// collide with paddles short edge (later)
	return true
}

func movePaddle(ball *Paddle, h int) {
	// collide with top or bottom
	if ball.y < 0 || ball.y+ball.height > float32(h) {
		ball.dy = -ball.dy
	}

	// collide with left or right
	// if ball.x < 0 || ball.x+ball.width > float32(w) {
	// 	ball.dx = -ball.dx
	// }
}

func deflect(ball *Ball, w, h int) {
	// collide with top or bottom
	if ball.y < 0 || ball.y+ball.height > float32(h) {
		ball.dy = -ball.dy
	}

	// collide with left or right
	if ball.x < 0 || ball.x+ball.width > float32(w) {
		ball.dx = -ball.dx
	}
}

func init() {
	runtime.LockOSThread()
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
	{
		window.SetKeyCallback(
			func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
				fmt.Printf("keypress: key %v, scancode %v, action %v, mods %v\n", key, scancode, action, mods)
			},
		)

		var cursorX float64
		var cursorY float64

		window.SetMouseButtonCallback(
			func(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
				fmt.Printf("click: button %v, x %.2f, y %.2f, action %v, mods %v\n", button, cursorX, cursorY, action, mods)
			},
		)

		window.SetCursorPosCallback(
			func(window *glfw.Window, xpos float64, ypos float64) {
				cursorX = xpos
				cursorY = ypos
				//fmt.Printf("position: xpox %v, ypos %v\n", xpos, ypos)
			},
		)
	}

	// window.SetCursorPosCallback(
	// 	func(window *glfw.Window, xpos, ypos float64) {
	// 		fmt.Printf("position: %.2f %.2f\n", xpos, ypos)
	// 	},
	// )

	w, h := window.GetFramebufferSize()
	// w, h := float32(w1), float32(h1)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(w), float64(h), 0, 1, -1)
	gl.MatrixMode(gl.MODELVIEW)

	gl.Color4f(1.0, 1.0, 1.0, 1.0)
	// var yoff float32 = 1.0

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
	// var dir float32 = 1.0

	ball := Ball{
		width:  15,
		height: 120,
		dx:     0.5,
		dy:     -1.3,
	}

	paddleLeft := Paddle{
		width:  15,
		height: 120,
		x:      10,
		y:      float32(h) / 2,
		// dy:     -1.0,
	}

	paddleRight := Paddle{
		width:  30,
		height: 30,
		x:      0.5,
		y:      float32(h) / 2,
		// dy:     -1.0,
	}

	ball.x = float32(w)/2 - ball.width/2
	ball.y = float32(h)/2 - ball.height/2

	for !window.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT)

		// gl.Rectf(
		// 	w/2-15,
		// 	h/2-25+yoff,
		// 	w/2+15,
		// 	h/2+25+yoff,
		// )
		//
		// yoff += (0.25 * dir)
		// dir = bounce(h/2-25+yoff, h/2+25+yoff, h)
		gl.Rectf(
			ball.x,
			ball.y,
			ball.x+ball.width,
			ball.y+ball.height,
		)

		gl.Rectf(
			paddleLeft.x,
			paddleLeft.y,
			paddleLeft.x+paddleLeft.width,
			paddleLeft.y+paddleLeft.height,
		)

		gl.Rectf(
			paddleRight.x,
			paddleRight.y,
			paddleRight.x+paddleRight.width,
			paddleRight.y+paddleRight.height,
		)

		ball.x += ball.dx
		ball.y += ball.dy

		deflect(&ball, w, h)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
