package grid

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Cell represents a cell in the ScreenGrid
type Cell struct {
	X, Y       int
	ScreenGrid ScreenGrid
	colour     [3]uint8
	rect       *sdl.Rect
}

// Draw updates the colour of the cell on the screen
func (c *Cell) Draw(surface *sdl.Surface) {
	surface.FillRect(c.rect, sdl.MapRGB(
		surface.Format, c.colour[0], c.colour[1], c.colour[2],
	),
	)
}

// SetColour sets the colour of the cell
func (c *Cell) SetColour(r, g, b uint8) {
	c.colour[0] = r
	c.colour[1] = g
	c.colour[2] = b
}

// ScreenGrid is a tool for displaying a ScreenGrid in SDL
type ScreenGrid struct {
	CellSize, Width, Height              int
	TopLeftPixelX, TopLeftPixelY         int
	BottomRightPixelX, BottomRightPixelY int
	Cells                                [][]*Cell
	AllCells                             []*Cell
}

// GetCell returns the cell at coordinates (x,y)
func (s *ScreenGrid) GetCell(x, y int) *Cell {
	return s.Cells[x][y]
}

// NewScreenGrid returns a new ScreenGrid
func NewScreenGrid(
	size, width, height, topLeftX, topLeftY,
	bottomRightX, bottomRightY int) ScreenGrid {
	var cells [][]*Cell
	var allCells []*Cell

	for x := 0; x < width; x++ {
		pixelX := topLeftX + (size * x)
		var row []*Cell
		for y := 0; y < height; y++ {
			pixelY := topLeftY + (size * y)
			rect := sdl.Rect{
				X: int32(pixelX), Y: int32(pixelY), H: int32(size), W: int32(size),
			}
			cell := Cell{X: x, Y: y, rect: &rect}
			row = append(row, &cell)
			allCells = append(allCells, &cell)
		}
		cells = append(cells, row)
	}
	screenGrid := ScreenGrid{
		CellSize: size, Width: width, Height: height, TopLeftPixelX: topLeftX,
		TopLeftPixelY: topLeftY, BottomRightPixelX: bottomRightX,
		BottomRightPixelY: bottomRightY, Cells: cells, AllCells: allCells,
	}
	return screenGrid
}
