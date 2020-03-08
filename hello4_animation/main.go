// hello4_animation - Animates hello.png scrolling up the window

package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth  = 640
	windowHeight = 480
	scrollSpeed  = 300 // speed in pixels/second
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
	dest := sdl.Rect{
		W: w,
		H: h,
	}

	// position the sprite at the bottom center of the window
	// origin is the top left corner, positive y is down
	dest.X = (windowWidth - dest.W) / 2

	// require float resolution for y position
	var yPos float32 = windowHeight

	// animation loop
	for dest.Y >= -dest.H {
		// clear the window
		rend.Clear()

		// set the y position in the struct
		dest.Y = int32(yPos)

		// draw the image to the window
		rend.Copy(tex, nil, &dest)
		rend.Present()

		// update sprite position
		yPos -= scrollSpeed / 60

		// wait 1/60th of a second
		time.Sleep((1000 / 60) * time.Millisecond)
	}
}
