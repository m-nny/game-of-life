package halflife

import (
	"fmt"
	"strings"

	"minmax.uk/game-of-life/pkg/boards"
	"minmax.uk/game-of-life/pkg/engine/halflife/cell"
)

type Universe struct {
	initRows int
	initCols int

	size int
	root *cell.MacroCell
}

func (u *Universe) BoardString() string {
	lines := u.root.BoardStrings()
	lines = lines[:u.initRows]
	for i := range lines {
		lines[i] = lines[i][:u.initCols]
	}
	return strings.Join(lines, "\n")
}

func (u *Universe) DebugPrint(tree bool) {
	fmt.Printf("universe: %+v\n", u)
	u.root.PrintBoard()
	// fmt.Printf("cache: %+v\n", cell.cell_cache)
	if tree {
		u.root.PrintTree("", true)
	}
	fmt.Println()
}

// BuildUniverse builds universe of size 2**level by 2**level
func BuildUniverse(level int) *Universe {
	root := cell.EmptyTree(level)
	return &Universe{
		size: 1 << level,
		root: root,
	}
}

func (m *Universe) Set(row, col int, value bool) {
	m.root = m.root.Set(row, col, value)
}

func FromBoardSpec(board boards.BoardSpec) (*Universe, error) {
	level := 0
	for (1 << level) < max(board.Rows, board.Cols) {
		level++
	}
	fmt.Printf("rows: %d level %d w %d\n", board.Rows, level, 1<<level)
	w := 1 << level
	var cells [][]bool
	str := board.Normalized()
	i := 0
	for range board.Rows {
		var line []bool
		for range board.Cols {
			line = append(line, boards.CHAR_TO_CELL[rune(str[i])])
			i++
		}
		for len(line) < w {
			line = append(line, false)
		}
		cells = append(cells, line)
	}
	for len(cells) < w {
		line := make([]bool, w)
		cells = append(cells, line)
	}
	root := cell.FromString(0, w-1, 0, w-1, 0, cells)
	u := &Universe{root: root, size: w, initRows: int(board.Rows), initCols: int(board.Cols)}
	u.initRows = int(board.Rows)
	u.initCols = int(board.Cols)
	return u, nil
}
