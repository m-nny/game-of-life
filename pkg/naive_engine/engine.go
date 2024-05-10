package naive_engine

import (
	"minmax.uk/game-of-life/pkg/engine"
)

var _ engine.Engine = (*NaiveEngine)(nil)

type NaiveEngine struct {
	Rows int64
	Cols int64

	cells     []bool
	prevCells []bool

	nBoard      []int
	nBoardReady bool
}

func (*NaiveEngine) Name() string {
	return "NaiveEngine"
}

var dxs = [][]int64{
	{-1, -1},
	{-1, 0},
	{-1, 1},

	{0, -1},
	// {0, 0},
	{0, 1},

	{1, -1},
	{1, 0},
	{1, 1},
}

var cell_map = map[bool][]bool{
	true:  {false, false, true, true, false, false, false, false, false},
	false: {false, false, false, true, false, false, false, false, false},
}

func (e *NaiveEngine) calcNboard() []int {
	if e.nBoardReady {
		return e.nBoard
	}
	for row := range e.Rows {
		for col := range e.Cols {
			n := 0
			for _, dx := range dxs {
				n_row, n_col := (row+dx[0]+e.Rows)%e.Rows, (col+dx[1]+e.Cols)%e.Cols
				n_i := n_col + n_row*e.Cols
				if e.cells[n_i] {
					n++
				}
			}
			i := col + row*e.Cols
			e.nBoard[i] = n
		}
	}
	e.nBoardReady = true
	return e.nBoard
}

func (e *NaiveEngine) Iterate() {
	e.calcNboard()
	e.prevCells, e.cells = e.cells, e.prevCells
	e.nBoardReady = false
	for i := range e.Rows * e.Cols {
		n := e.nBoard[i]
		e.cells[i] = cell_map[e.prevCells[i]][n]
	}
}
