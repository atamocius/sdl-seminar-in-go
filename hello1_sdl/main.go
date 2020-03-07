// hello1_sdl - Initializes SDL

package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// attempt to initialize graphics system
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		fmt.Printf("error initializing SDL: %s\n", err)
		return
	}
	// clean up resources before exiting
	defer sdl.Quit()

	fmt.Println("initialization successful!")
}
