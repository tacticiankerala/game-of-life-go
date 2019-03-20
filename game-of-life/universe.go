package gameoflife

const (
	CellStateDead    = 0
	CellStateNewCell = 1
	CellStateAlive   = 2
)

type CellState int

func (state CellState) isAlive() bool {
	if state == CellStateNewCell || state == CellStateAlive {
		return true
	}
	return false
}

type Universe struct {
	Width, Height int
	Cells         [][]CellState
}

func NewUniverse(width, height int) *Universe {
	cells := make([][]CellState, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]CellState, width)
		for j := 0; j < width; j++ {
			cells[i][j] = CellStateDead
		}
	}
	return &Universe{Width: width, Height: height, Cells: cells}
}

func NewUniverseFromLiveCoordinates(width, height int, liveCells [][2]int) *Universe {
	cells := make([][]CellState, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]CellState, width)
		for j := 0; j < width; j++ {
			if contains(liveCells, [2]int{i, j}) {
				cells[i][j] = CellStateAlive
			} else {
				cells[i][j] = CellStateDead
			}
		}
	}
	return &Universe{Width: width, Height: height, Cells: cells}
}

type Neighbours [3][3]CellState

func NewNeighbours() *Neighbours {
	width, height := 3, 3
	neighbours := Neighbours{}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			neighbours[i][j] = CellStateDead
		}
	}
	return &neighbours
}

func (neighbours *Neighbours) numberOfAliveCells() int {
	count := 0
	for _, row := range neighbours {
		for _, col := range row {
			if col.isAlive() {
				count++
			}
		}
	}
	return count
}

func neighbourCoordinates(x, y int) [3][3][2]int {
	return [3][3][2]int{
		[3][2]int{
			[2]int{x - 1, y - 1}, [2]int{x - 1, y}, [2]int{x - 1, y + 1},
		},
		[3][2]int{
			[2]int{x, y - 1}, [2]int{x, y}, [2]int{x, y + 1},
		},
		[3][2]int{
			[2]int{x + 1, y - 1}, [2]int{x + 1, y}, [2]int{x + 1, y + 1},
		},
	}
}

func (universe *Universe) neighbours(x, y int) *Neighbours {
	neighbours := NewNeighbours()
	for i, row := range neighbourCoordinates(x, y) {
		for j, col := range row {
			neighbourX, neighbourY := col[0], col[1]
			if (neighbourX < 0) || (neighbourY < 0) || (neighbourX >= universe.Width) || (neighbourY >= universe.Height) {
				continue
			} else {
				neighbours[i][j] = universe.Cells[neighbourX][neighbourY]
			}
		}
	}
	return neighbours
}

func (universe *Universe) RefreshUniverse() {
	nextUniverse := NewUniverse(universe.Width, universe.Height)
	copy(nextUniverse.Cells, universe.Cells)
	for x, row := range universe.Cells {
		for y, col := range row {
			numberOfAliveCells := universe.neighbours(x, y).numberOfAliveCells()
			numberOfAliveCellsExceptCurrent := numberOfAliveCells
			if col.isAlive() {
				numberOfAliveCellsExceptCurrent--
			}

			nextUniverse.Cells[x][y] = nextCellStatus(col, numberOfAliveCellsExceptCurrent)
		}
	}
	copy(universe.Cells, nextUniverse.Cells)
}

func nextCellStatus(currentCellStatus CellState, numberOfAliveCellsExceptCurrent int) CellState {
	var nextStatus CellState = CellStateDead
	if !currentCellStatus.isAlive() {
		if numberOfAliveCellsExceptCurrent == 3 {
			nextStatus = CellStateNewCell
		}
	} else {
		if numberOfAliveCellsExceptCurrent < 2 || numberOfAliveCellsExceptCurrent > 3 {
			nextStatus = CellStateDead
		} else if numberOfAliveCellsExceptCurrent == 2 || numberOfAliveCellsExceptCurrent == 3 {
			nextStatus = CellStateAlive
		}
	}
	return nextStatus
}

func contains(coordinates [][2]int, coordinate [2]int) bool {
	for _, c := range coordinates {
		if c == coordinate {
			return true
		}
	}
	return false
}
