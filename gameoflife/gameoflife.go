package gameoflife

// Cell represents a cell in the Game of Life
type Cell struct {
	Xpos, Ypos int
	Alive      bool
	Grid       *Grid
	aliveNext  bool
}

func (c *Cell) getNeighbors() []*Cell {
	var neighbors []*Cell
	if c.Xpos > 0 { // Left
		neighbors = append(neighbors, c.Grid.GetCell(c.Xpos-1, c.Ypos))
	}
	if c.Xpos < c.Grid.Width-1 { // Right
		neighbors = append(neighbors, c.Grid.GetCell(c.Xpos+1, c.Ypos))
	}
	if c.Ypos > 0 { // Up
		neighbors = append(neighbors, c.Grid.GetCell(c.Xpos, c.Ypos-1))
	}
	if c.Ypos < c.Grid.Height-1 { // Down
		neighbors = append(neighbors, c.Grid.GetCell(c.Xpos, c.Ypos+1))
	}
	if c.Ypos > 0 && c.Xpos > 0 { // Up Left
		neighbors = append(neighbors, c.Grid.GetCell(c.Xpos-1, c.Ypos-1))
	}
	if c.Ypos > 0 && c.Xpos < c.Grid.Width-1 { // Up Right
		neighbors = append(neighbors, c.Grid.GetCell(c.Xpos+1, c.Ypos-1))
	}
	if c.Ypos < c.Grid.Height-1 && c.Xpos > 0 { // Down Left
		neighbors = append(neighbors, c.Grid.GetCell(c.Xpos-1, c.Ypos+1))
	}
	if c.Ypos < c.Grid.Height-1 && c.Xpos < c.Grid.Width-1 { // Down Right
		neighbors = append(neighbors, c.Grid.GetCell(c.Xpos+1, c.Ypos+1))
	}
	return neighbors
}

func (c *Cell) countLivingNeighbours() int {
	livingNeighbours := 0
	for _, cell := range c.getNeighbors() {
		if cell.Alive {
			livingNeighbours++
		}
	}
	return livingNeighbours
}

func (c *Cell) checkUpdate() {
	livingNeighbours := c.countLivingNeighbours()
	if c.Alive == true {
		c.aliveNext = livingNeighbours >= 2 && livingNeighbours <= 3
	} else {
		c.aliveNext = livingNeighbours == 3
	}
}

func (c *Cell) updateState() {
	c.Alive = c.aliveNext
}

// Grid is the grid on which Game of Life is played
type Grid struct {
	Width, Height int
	Cells         [][]*Cell
	AllCells      []*Cell
}

// GetCell returns the cell at position (x, y)
func (g *Grid) GetCell(x, y int) *Cell {
	return g.Cells[x][y]
}

func (g *Grid) newCell(x, y int) Cell {
	return Cell{Grid: g, Xpos: x, Ypos: y, Alive: false}
}

func (g *Grid) getNextState() {
	for _, cell := range g.AllCells {
		cell.checkUpdate()
	}
}

func (g *Grid) updateState() {
	for _, cell := range g.AllCells {
		cell.updateState()
	}
}

// Update moves the game to the next state
func (g *Grid) Update() {
	g.getNextState()
	g.updateState()
}

// NewGrid returns a Grid of size width x height
func NewGrid(width, height int) Grid {
	var cells [][]*Cell
	var allCells []*Cell
	grid := Grid{Width: width, Height: height, Cells: cells}
	for x := 0; x < width; x++ {
		var row []*Cell
		for y := 0; y < height; y++ {
			cell := grid.newCell(x, y)
			row = append(row, &cell)
			allCells = append(allCells, &cell)
		}
		cells = append(cells, row)
	}
	grid.Cells = cells
	grid.AllCells = allCells
	return grid
}
