// hello3_image - Initializes SDL, loads an image, And displays it in a window

package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// attempt to initialize graphics and timer system
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
		640,
		480,
		0,
	)
	if err != nil {
		fmt.Printf("error creating window: %s\n", err)
		return
	}
	// clean up resources before exiting
	defer win.Destroy()

	// create a renderer, which sets up the graphics hardware
	var renderFlags uint32 = sdl.RENDERER_ACCELERATED | sdl.RENDERER_PRESENTVSYNC
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

	// clear the window
	rend.Clear()

	// draw the image to the window
	rend.Copy(tex, nil, nil)
	rend.Present()

	// wait a few seconds
	time.Sleep(5000 * time.Millisecond)
}
