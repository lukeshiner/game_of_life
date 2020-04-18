package gameoflife

import (
	"testing"
)

func TestCell(t *testing.T) {
	x := 5
	y := 8
	alive := true
	cell := Cell{Xpos: x, Ypos: y, Alive: alive}
	if cell.Xpos != x {
		t.Errorf("Cell.Xpos should be %d, was %d", cell.Xpos, x)
	}
	if cell.Ypos != y {
		t.Errorf("Cell.Ypos should be %d, was %d", cell.Ypos, y)
	}
	if cell.Alive != alive {
		t.Errorf("Cell.Alive should be %t, was %t", cell.Alive, alive)
	}
}

func TestNewGrid(t *testing.T) {
	width := 6
	height := 4
	grid := NewGrid(width, height)
	if grid.Width != width {
		t.Errorf("Expected grid.Width to be %d, got %d", grid.Width, width)
	}
	if grid.Height != height {
		t.Errorf("Expected grid.Height to be %d, got %d", grid.Height, height)
	}
	if len(grid.Cells) != width {
		t.Errorf("Expected grid to have %d rows, has %d", width, len(grid.Cells))
	}
	for _, column := range grid.Cells {
		if len(column) != height {
			t.Errorf("Expected grid to have %d columns, has %d", height, len(column))
		}
	}
	for x, row := range grid.Cells {
		for y, cell := range row {
			if cell.Xpos != x {
				t.Errorf("Expected cell.Xpos to be %d, was %d", x, cell.Xpos)
			}
			if cell.Ypos != y {
				t.Errorf("Expected cell.Ypos to be %d, was %d", y, cell.Ypos)
			}
		}
	}
}

func TestGridGetCell(t *testing.T) {
	grid := NewGrid(4, 6)
	cell := grid.GetCell(2, 4)
	if cell.Xpos != 2 || cell.Ypos != 4 {
		t.Errorf("Expected cell, (2,4), got cell (%d,%d)", cell.Xpos, cell.Ypos)
	}
}

func TestGetNeighbors(t *testing.T) {
	var tests = []struct {
		cellX, cellY int
		expected     [][]int
	}{
		{
			cellX: 4,
			cellY: 4,
			expected: [][]int{
				{3, 4}, {5, 4}, {4, 3}, {4, 5}, {3, 3}, {5, 3}, {3, 5}, {5, 5},
			},
		},
		{
			cellX: 0,
			cellY: 4,
			expected: [][]int{
				{1, 4}, {0, 3}, {0, 5}, {1, 3}, {1, 5},
			},
		},
		{
			cellX: 7,
			cellY: 4,
			expected: [][]int{
				{6, 4}, {7, 3}, {7, 5}, {6, 3}, {6, 5},
			},
		},
		{
			cellX: 4,
			cellY: 0,
			expected: [][]int{
				{3, 0}, {5, 0}, {4, 1}, {3, 1}, {5, 1},
			},
		},
		{
			cellX: 4,
			cellY: 7,
			expected: [][]int{
				{3, 7}, {5, 7}, {4, 6}, {3, 6}, {5, 6},
			},
		},
		{
			cellX: 0,
			cellY: 0,
			expected: [][]int{
				{1, 0}, {0, 1}, {1, 1},
			},
		},
		{
			cellX: 7,
			cellY: 0,
			expected: [][]int{
				{6, 0}, {7, 1}, {6, 1},
			},
		},
		{
			cellX: 0,
			cellY: 7,
			expected: [][]int{
				{1, 7}, {0, 6}, {1, 6},
			},
		},
		{
			cellX: 7,
			cellY: 7,
			expected: [][]int{
				{6, 7}, {7, 6}, {6, 6},
			},
		},
	}
	grid := NewGrid(8, 8)
	for _, test := range tests {
		cell := grid.GetCell(test.cellX, test.cellY)
		neighbours := cell.getNeighbors()
		if len(neighbours) != len(test.expected) {
			t.Errorf(
				"Expected cell (%d,%d) to have %d neighbors, has %d",
				test.cellX, test.cellY, len(test.expected), len(neighbours),
			)
		} else {
			for i, neighbour := range neighbours {
				position := test.expected[i]
				if neighbour.Xpos != position[0] || neighbour.Ypos != position[1] {
					t.Errorf(
						"Expected cell (%d,%d), got cell (%d,%d)",
						position[0], position[1], neighbour.Xpos, neighbour.Ypos,
					)
				}
			}
		}
	}
}

func cellInSlice(cell *Cell, coords [][2]int) bool {
	for _, coord := range coords {
		if cell.Xpos == coord[0] && cell.Ypos == coord[1] {
			return true
		}
	}
	return false
}

func TestUpdateCellState(t *testing.T) {
	grid := NewGrid(8, 8)
	cell := grid.GetCell(4, 4)
	cell.Alive = false
	cell.aliveNext = true
	cell.updateState()
	if cell.Alive != true {
		t.Error("Cell did not update state")
	}
}

func TestUpdate(t *testing.T) {
	var tests = []struct {
		startAlive [][2]int
		endAlive   [][2]int
	}{
		{
			startAlive: [][2]int{{4, 4}},
			endAlive:   [][2]int{},
		},
		{
			startAlive: [][2]int{{3, 3}, {5, 4}, {3, 5}},
			endAlive:   [][2]int{{4, 4}},
		},
		{
			startAlive: [][2]int{{4, 4}, {4, 5}},
			endAlive:   [][2]int{},
		},
		{
			startAlive: [][2]int{{4, 4}, {3, 3}, {5, 3}, {3, 5}, {5, 5}},
			endAlive:   [][2]int{{4, 3}, {4, 5}, {3, 4}, {5, 4}},
		},
	}
	for _, test := range tests {
		grid := NewGrid(7, 7)
		for _, cellCords := range test.startAlive {
			grid.GetCell(cellCords[0], cellCords[1]).Alive = true
		}
		grid.Update()
		for _, cell := range grid.AllCells {
			cellInEndAlive := cellInSlice(cell, test.endAlive)
			if cell.Alive == true && cellInEndAlive == false {
				t.Errorf(
					"Cell (%d,%d) was alive when it should be dead",
					cell.Xpos, cell.Ypos,
				)
			} else if cell.Alive == false && cellInEndAlive == true {
				t.Errorf(
					"Cell (%d,%d) was dead when it should be alive",
					cell.Xpos, cell.Ypos,
				)
			}
		}
	}
}
