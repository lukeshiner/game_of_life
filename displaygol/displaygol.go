package displaygol

import (
	"github.com/lukeshiner/gameoflife"
	"github.com/lukeshiner/grid"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

var aliveColour = [3]uint8{0, 255, 0}
var deadColour = [3]uint8{255, 255, 255}

// Cell represents a cell in the display grid
type Cell struct {
	screenCell *grid.Cell
	golCell    *gameoflife.Cell
}

// Draw updates the cell on the screen
func (c *Cell) Draw(surface *sdl.Surface) {
	if c.golCell.Alive {
		c.screenCell.SetColour(aliveColour[0], aliveColour[1], aliveColour[2])
	} else {
		c.screenCell.SetColour(deadColour[0], deadColour[1], deadColour[2])
	}
	c.screenCell.Draw(surface)
}

// GameOfLifeGrid displays a grid on an SDL surface
type GameOfLifeGrid struct {
	displayGrid *grid.ScreenGrid
	game        *gameoflife.Grid
	cells       []*Cell
	rows        [][]*Cell
}

// NewScreenGrid creates a new Game of Life on an SDL surface
func NewScreenGrid(
	golWidth, golHeight, golOverflow, topLeftX, topLeftY,
	bottomRightX, bottomRightY int) GameOfLifeGrid {
	var cells []*Cell
	var rows [][]*Cell
	cellPixels := (bottomRightX - topLeftX) / golWidth
	displayGrid := grid.NewScreenGrid(
		cellPixels, golWidth, golHeight, topLeftX,
		topLeftY, bottomRightX, bottomRightY,
	)
	game := gameoflife.NewGrid(golWidth+(golOverflow*2), golHeight+(golOverflow*2))
	for x := 0; x < golWidth; x++ {
		var row []*Cell
		for y := 0; y < golHeight; y++ {
			golCell := game.GetCell(x+golOverflow, y+golOverflow)
			displayCell := displayGrid.GetCell(x, y)
			cell := Cell{screenCell: displayCell, golCell: golCell}
			row = append(row, &cell)
			cells = append(cells, &cell)
		}
		rows = append(rows, row)
	}
	for _, cell := range cells {
		if rand.Float32() > 0.7 {
			cell.golCell.Alive = true
		}
	}
	return GameOfLifeGrid{displayGrid: &displayGrid, game: &game, cells: cells, rows: rows}
}

// Update updates the state of the game
func (g *GameOfLifeGrid) Update() {
	g.game.Update()
}

// Draw the current state of the game to a surface
func (g *GameOfLifeGrid) Draw(surface *sdl.Surface) {
	for _, cell := range g.cells {
		cell.Draw(surface)
	}
}
