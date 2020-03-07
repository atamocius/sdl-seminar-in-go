// hello2_window - Initializes SDL and creates a window

package main

import (
	"fmt"
	"time"

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

	// wait a few seconds
	time.Sleep(5000 * time.Millisecond)
}
