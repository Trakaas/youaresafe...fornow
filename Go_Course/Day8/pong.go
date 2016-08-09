package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"math/rand"
	"runtime"
	"strings"
	"time"

	"github.com/go-gl/gl/v3.3-compatibility/gl"
	"github.com/go-gl/glfw/v3.1/glfw"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
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

func reactEvents(ball *Ball, paddleLeft *Paddle, paddleRight *Paddle, w, h int) (bool, bool) {
	// collide with top or bottom
	if ball.y < 0 || ball.y+ball.height > float32(h) {
		ball.dy = -ball.dy
		return false, false
	}

	// paddleLeft collide with top
	if paddleLeft.y < 0 {
		paddleLeft.y = 0.0
	} else if paddleLeft.y+paddleLeft.height > float32(h) {
		paddleLeft.y = float32(h) - paddleLeft.height
	}

	// paddleRight collide with top or bottom
	if paddleRight.y < 0 {
		paddleRight.y = 0.0
	} else if paddleRight.y+paddleRight.height > float32(h) {
		paddleRight.y = float32(h) - paddleRight.height
	}

	// collide with paddleLeft long edge
	if ball.x == paddleLeft.x+paddleLeft.width &&
		(ball.y+ball.height >= paddleLeft.y &&
			ball.y <= paddleLeft.y+paddleLeft.height) {
		ball.x = paddleLeft.x + paddleLeft.width
		ball.dx = -ball.dx
	}

	// collide with paddleRight long edge
	if ball.x+ball.width == paddleRight.x &&
		(ball.y+ball.height >= paddleRight.y &&
			ball.y <= paddleRight.y+paddleRight.height) {

		ball.x = paddleRight.x - ball.width
		ball.dx = -ball.dx
	}

	// collide with board edges beyond paddles
	if ball.x+ball.width > float32(w) {
		return true, true
	}

	if ball.x < 0 {
		return true, false
	}
	// collide with paddles short edge (later)
	return false, false
}

func inputEvents(wKey, sKey, upKey, downKey bool, paddleLeft, paddleRight *Paddle) {
	// paddleLeft
	if wKey {
		paddleLeft.dy = -1.0
	} else if sKey {
		paddleLeft.dy = 1.0
	} else {
		paddleLeft.dy = 0.0
	}

	// paddleRight
	if upKey {
		paddleRight.dy = -1.0
	} else if downKey {
		paddleRight.dy = 1.0
	} else {
		paddleRight.dy = 0
	}

}

func drawEvents(ball *Ball, paddleLeft, paddleRight *Paddle) {

	gl.Clear(gl.COLOR_BUFFER_BIT)

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
	paddleLeft.y += paddleLeft.dy
	paddleRight.y += paddleRight.dy
}

func textTexture(text string, font *truetype.Font, size float64, width int, height int, x int, y int) (texture uint32) {
	dest := image.NewNRGBA(image.Rect(0, 0, width, height))

	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	lines := strings.Split(text, "\n")

	c := freetype.NewContext()
	c.SetDst(dest)
	c.SetClip(dest.Bounds())
	c.SetSrc(image.White)
	c.SetFont(font)
	c.SetFontSize(size)
	metrics := truetype.NewFace(font, &truetype.Options{Size: size}).Metrics()

	y += metrics.Ascent.Round()
	for _, line := range lines {
		_, err := c.DrawString(line, freetype.Pt(x, y))
		if err != nil {
			log.Fatalln(err)
		}
		y += metrics.Height.Round()
	}

	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(dest.Bounds().Dx()),
		int32(dest.Bounds().Dy()),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(dest.Pix),
	)

	return texture
}

func drawQuad(texture uint32, top, right, bottom, left, topTex, rightTex, bottomTex, leftTex float32) {
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.Begin(gl.QUADS)
	// upper left
	gl.TexCoord2f(leftTex, topTex)
	gl.Vertex2f(left, top)
	// upper right
	gl.TexCoord2f(rightTex, topTex)
	gl.Vertex2f(right, top)
	// lower right
	gl.TexCoord2f(rightTex, bottomTex)
	gl.Vertex2f(right, bottom)
	// lower left
	gl.TexCoord2f(leftTex, bottomTex)
	gl.Vertex2f(left, bottom)
	gl.End()
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
	window, err := glfw.CreateWindow(1600, 900, "Pong", nil, nil)
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

	gl.Enable(gl.TEXTURE_2D)                           // set 2d texture mapping
	gl.Enable(gl.BLEND)                                // set alpha blending on
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA) // set alpha blending type
	gl.ClearColor(0, 0, 0, 1)                          // setting background alpha

	rand.Seed(time.Now().UnixNano())

	var wKey, sKey, upKey, downKey bool

	// Define Callbacks to event functions
	{

		window.SetKeyCallback(
			func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
				if action == glfw.Press {
					switch key {
					case glfw.KeyW:
						wKey = true
					case glfw.KeyS:
						sKey = true
					case glfw.KeyDown:
						downKey = true
					case glfw.KeyUp:
						upKey = true
					}
				} else if action == glfw.Release {
					switch key {
					case glfw.KeyW:
						wKey = false
					case glfw.KeyS:
						sKey = false
					case glfw.KeyDown:
						downKey = false
					case glfw.KeyUp:
						upKey = false
					}
				}
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
			},
		)

	}

	w, h := window.GetFramebufferSize()

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(w), float64(h), 0, 1, -1)
	gl.MatrixMode(gl.MODELVIEW)

	gl.Color4f(1.0, 1.0, 1.0, 1.0)

	// font
	var font *truetype.Font
	{
		data, err := ioutil.ReadFile("luxisr.ttf")
		if err != nil {
			log.Fatalln(err)
		}
		font, err = freetype.ParseFont(data)
		if err != nil {
			log.Fatalln(err)
		}
	}

	ball := Ball{
		width:  30,
		height: 30,
		dx:     -0.25,
		dy:     .25,
		x:      float32(w)/2 - 30/2,
		y:      float32(h/2) - 30/2,
	}

	if rand.Intn(2) == 0 {
		ball.dx = -ball.dx
	}
	if rand.Intn(2) == 0 {
		ball.dy = -ball.dy
	}
	scoreTex := textTexture("0                               0", font, 40, 600, 600, 0, 0)

	paddleLeft := Paddle{
		width:  15,
		height: 120,
		x:      50,
		y:      float32(h) / 2,
	}

	paddleRight := Paddle{
		width:  15,
		height: 120,
		x:      float32(w) - 50 - 15,
		y:      float32(h) / 2,
	}

	gameOver := false
	gameOverTex := textTexture("GAME OVER", font, 60, 600, 610, 0, 0)

	var playerScoreRight, playerScoreLeft int
	maxScore := 10

	reset := func() {
		scoreString := fmt.Sprintf("%d                               %d", playerScoreLeft, playerScoreRight)
		scoreTex = textTexture(scoreString, font, 40, 600, 600, 0, 0)

		ball.x = float32(w/2) - ball.width/2
		ball.y = float32(h/2) - ball.height/2

		if playerScoreLeft >= maxScore || playerScoreRight >= maxScore {
			gameOver = true
		}
	}

	var scoreState, sideState bool

	gl.Color4f(1.0, 1.0, 1.0, 1.0)
	for !window.ShouldClose() {

		if !gameOver {
			scoreState, sideState = reactEvents(&ball, &paddleLeft, &paddleRight, w, h)
		}

		inputEvents(wKey, sKey, upKey, downKey, &paddleLeft, &paddleRight)

		if scoreState == true && !gameOver {
			if sideState == true {
				playerScoreRight++
				fmt.Println("Player 2 Score :", playerScoreRight)
			} else {
				playerScoreLeft++
				fmt.Println("Player 1 Score :", playerScoreLeft)
			}
			reset()
		}

		if !gameOver {
			drawEvents(&ball, &paddleLeft, &paddleRight)
			gl.Enable(gl.TEXTURE_2D)
			drawQuad(scoreTex, 20, 1200, 620, 600,
				0, 1, 1, 0)
		} else {
			gl.Enable(gl.TEXTURE_2D)
			drawQuad(gameOverTex, 60, 1200, 660, 600,
				0, 1, 1, 0)
		}
		gl.Disable(gl.TEXTURE_2D)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
