package engine

import (
	"fmt"
	"strings"

	"minmax.uk/game-of-life/pkg/boards"
)

func Empty(rows, cols int64) *Engine {
	cells := make([]bool, cols*rows)
	prevCells := make([]bool, cols*rows)
	return &Engine{
		Cols: cols,
		Rows: rows,

		cells:     cells,
		prevCells: prevCells,
	}
}

func FromBoardSpec(board boards.BoardSpec) (*Engine, error) {
	rows, cols, str := board.Rows, board.Cols, board.Normalized()

	if len(str) < int(cols*rows) {
		return nil, fmt.Errorf("not enough chars")
	}
	if len(str) > int(cols*rows) {
		return nil, fmt.Errorf("too much chars")
	}
	g := Empty(rows, cols)
	i := int64(0)
	for _, rune := range str {
		if strings.ContainsRune(boards.EMPTY_CELLS, rune) {
			g.cells[i] = false
		} else if strings.ContainsRune(boards.FULL_CELLS, rune) {
			g.cells[i] = true
		} else {
			return nil, fmt.Errorf("illegal rune: {%+v} {%d}", rune, rune)
		}
		i++
	}
	if i != cols*rows {
		return nil, fmt.Errorf("not enough chars")
	}
	return g, nil
}
