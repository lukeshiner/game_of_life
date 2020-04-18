package main

import (
	"fmt"
	"github.com/lukeshiner/displaygol"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth         = 800
	screenHeight        = 600
	fps                 = 8
	delay        uint32 = 1000 / fps
)

func main() {
	var frameTime, frameStart uint32
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Game of Life",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window;", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
	}
	defer renderer.Destroy()
	grid := displaygol.NewScreenGrid(78, 58, 20, 10, 10, 789, 589)
	for {
		frameStart = sdl.GetTicks()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()
		screenSurface, err := window.GetSurface()
		if err != nil {
			fmt.Println("Failed to create surface:", err)
		}
		grid.Draw(screenSurface)
		grid.Update()
		window.UpdateSurface()
		frameTime = sdl.GetTicks() - frameStart
		if frameTime < delay {
			sdl.Delay(delay - frameTime)
		}
	}
}
