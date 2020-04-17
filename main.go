package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

const (
	screenWidth         = 800
	screenHeight        = 600
	fps                 = 8
	delay        uint32 = 1000 / fps
)

type gridSquare struct {
	xPos, yPos int
	alive      bool
	rect       sdl.Rect
}

func newGridSquare(x, y, size int) *gridSquare {
	alive := rand.Float64() > 0.8
	rect := sdl.Rect{X: int32(x), Y: int32(y), H: int32(size), W: int32(size)}
	return &gridSquare{alive: alive, rect: rect, xPos: x / size, yPos: y / size}
}

type grid struct {
	width, height, size int
	squares             [][]*gridSquare
}

func (g *grid) update() {
	for _, row := range g.squares {
		for _, square := range row {
			aliveNeighbours := 0
			for _, cell := range g.getNeighbors(square) {
				if cell.alive {
					aliveNeighbours++
				}
			}
			if square.alive == true {
				if aliveNeighbours < 2 || aliveNeighbours > 3 {
					square.alive = false
				}
			} else {
				if aliveNeighbours == 3 {
					square.alive = true
				}
			}
		}
	}
}

func (g *grid) getCell(x, y int) *gridSquare {
	return g.squares[x][y]
}

func (g *grid) getNeighbors(square *gridSquare) []*gridSquare {
	var neighbors []*gridSquare
	if square.xPos > 0 { // Left
		neighbors = append(neighbors, g.squares[square.xPos-1][square.yPos])
	}
	if square.xPos < g.width-1 { // Right
		neighbors = append(neighbors, g.squares[square.xPos+1][square.yPos])
	}
	if square.yPos > 0 { // Up
		neighbors = append(neighbors, g.squares[square.xPos][square.yPos-1])
	}
	if square.yPos < g.height-1 { // Down
		neighbors = append(neighbors, g.squares[square.xPos][square.yPos+1])
	}
	if square.yPos > 0 && square.xPos > 0 { // Up Left
		neighbors = append(neighbors, g.squares[square.xPos-1][square.yPos-1])
	}
	if square.yPos > 0 && square.xPos < g.width-1 { // Up Right
		neighbors = append(neighbors, g.squares[square.xPos+1][square.yPos-1])
	}
	if square.yPos < g.height-1 && square.xPos > 0 { // Down Left
		neighbors = append(neighbors, g.squares[square.xPos-1][square.yPos+1])
	}
	if square.yPos < g.height-1 && square.xPos < g.width-1 { // Down Right
		neighbors = append(neighbors, g.squares[square.xPos+1][square.yPos+1])
	}
	return neighbors
}

func (g *grid) draw(screenSurface *sdl.Surface) {
	for _, row := range g.squares {
		for _, square := range row {
			if square.alive {
				screenSurface.FillRect(&square.rect, sdl.MapRGB(screenSurface.Format, 0, 180, 0))
			} else {
				screenSurface.FillRect(&square.rect, sdl.MapRGB(screenSurface.Format, 255, 255, 255))
			}
		}
	}
}

func newGrid(size int) grid {
	var squares [][]*gridSquare
	width := screenWidth / size
	height := screenHeight / size
	for x := 0; x < screenWidth; x += size {
		var col []*gridSquare
		for y := 0; y < screenHeight; y += size {
			col = append(col, newGridSquare(x, y, size))
		}
		squares = append(squares, col)
	}
	return grid{width: width, height: height, squares: squares, size: size}
}

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
	grid := newGrid(10)
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
		grid.draw(screenSurface)
		grid.update()
		window.UpdateSurface()
		frameTime = sdl.GetTicks() - frameStart
		if frameTime < delay {
			sdl.Delay(delay - frameTime)
		}
	}
}
