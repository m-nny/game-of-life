package engine

import (
	"fmt"
	"strings"
)

type Engine struct {
	Rows int64
	Cols int64

	cells     []bool
	prevCells []bool

	nBoard      []int
	nBoardReady bool
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

var cell_map = [][]bool{
	// empty case
	{false, false, false, true, false, false, false, false, false},
	// full case
	{false, false, true, true, false, false, false, false, false},
}

func (e *Engine) calcNboard() []int {
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

func (e *Engine) Iterate() {
	e.calcNboard()
	e.prevCells, e.cells = e.cells, e.prevCells
	e.nBoardReady = false
	for i := range e.Rows * e.Cols {
		n := e.nBoard[i]
		if !e.prevCells[i] {
			e.cells[i] = cell_map[0][n]
		} else {
			e.cells[i] = cell_map[1][n]
		}
	}
}

func debugString(item any) string {
	switch item := item.(type) {
	case bool:
		if item {
			return "#"
		} else {
			return "."
		}
	}
	return string(fmt.Sprintf("%v", item)[0])
}

func debugSlice[S ~[]E, E comparable](arr S, lineLen int64) string {
	var builder strings.Builder
	for i, item := range arr {
		builder.WriteString(debugString(item))
		if i%int(lineLen) == int(lineLen)-1 && i+1 != len(arr) {
			builder.WriteRune('\n')
		}
	}
	return builder.String()
}
