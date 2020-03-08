// hello7_mouse - the sprite follows the cursor, or is chased if you click

package main

import (
	"fmt"
	"math"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth, windowHeight         = 640, 480
	speed                     float32 = 300 // speed in pixels/second
)

func main() {
	// attempt to initialize graphics
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		fmt.Printf("error initializing SDL: %s\n", err)
		return
	}
	// clean up resources before exiting
	defer sdl.Quit()

	win, err := sdl.CreateWindow(
		"Hello, CS50!",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		windowWidth,
		windowHeight,
		0,
	)
	if err != nil {
		fmt.Printf("error creating window: %s\n", err)
		return
	}
	// clean up resources before exiting
	defer win.Destroy()

	// create a renderer, which sets up the graphics hardware
	var renderFlags uint32 = sdl.RENDERER_ACCELERATED
	rend, err := sdl.CreateRenderer(win, -1, renderFlags)
	if err != nil {
		fmt.Printf("error creating renderer: %s\n", err)
		return
	}
	// clean up resources before exiting
	defer rend.Destroy()

	// load the image into memory using SDL_image library function
	surface, err := img.Load("../resources/hello.png")
	if err != nil {
		fmt.Println("error creating surface")
		return
	}

	// load the image data into the graphics hardware's memory
	tex, err := rend.CreateTextureFromSurface(surface)
	// free the surface after the texture has been created
	surface.Free()
	if err != nil {
		fmt.Printf("error creating texture: %s\n", err)
		return
	}
	// clean up resources before exiting
	defer tex.Destroy()

	// get the dimensions of texture
	_, _, w, h, err := tex.Query()
	if err != nil {
		fmt.Printf("error querying the texture: %s\n", err)
		return
	}

	// struct to hold the position and size of the sprite
	// get and scale the dimensions of texture
	dest := sdl.Rect{
		W: w / 4,
		H: h / 4,
	}

	// start sprite in center of screen
	xPos := float32((windowWidth - dest.W) / 2)
	yPos := float32((windowHeight - dest.H) / 2)
	var xVel, yVel float32 = 0, 0

	// set to true when window close button is pressed
	closeRequested := false

	// animation loop
	for !closeRequested {
		// process events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			if event.GetType() == sdl.QUIT {
				closeRequested = true
			}
		}

		// get cursor position relative to window
		mouseX, mouseY, buttons := sdl.GetMouseState()

		// determine velocity toward mouse
		targetX := mouseX - dest.W/2
		targetY := mouseY - dest.H/2
		deltaX := float32(targetX) - xPos
		deltaY := float32(targetY) - yPos
		distance := float32(math.Sqrt(float64(deltaX*deltaX + deltaY*deltaY)))

		// prevent jitter
		if distance < 5 {
			xVel, yVel = 0, 0
		} else {
			xVel = deltaX * speed / distance
			yVel = deltaY * speed / distance
		}

		// reverse velocity if mouse button 1 pressed
		if buttons&sdl.Button(sdl.BUTTON_LEFT) > 0 {
			xVel, yVel = -xVel, -yVel
		}

		// update positions
		xPos += xVel / 60
		yPos += yVel / 60

		// collision detection with bounds
		if xPos <= 0 {
			xPos = 0
		}
		if yPos <= 0 {
			yPos = 0
		}
		if int32(xPos) >= windowWidth-dest.W {
			xPos = float32(windowWidth - dest.W)
		}
		if int32(yPos) >= windowHeight-dest.H {
			yPos = float32(windowHeight - dest.H)
		}

		// set the positions in the struct
		dest.X, dest.Y = int32(xPos), int32(yPos)

		// clear the window
		rend.Clear()

		// draw the image to the window
		rend.Copy(tex, nil, &dest)
		rend.Present()

		// wait 1/60th of a second
		time.Sleep((1000 / 60) * time.Millisecond)
	}
}
