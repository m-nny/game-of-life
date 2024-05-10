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
	// if board.Rows != board.Cols {
	// 	return nil, fmt.Errorf("board should be square")
	// }
	level := 0
	for (1 << level) < max(board.Rows, board.Cols) {
		level++
	}
	fmt.Printf("rows: %d level %d w %d\n", board.Rows, level, 1<<level)
	u := BuildUniverse(level)
	u.initRows = int(board.Rows)
	u.initCols = int(board.Cols)
	str := board.Normalized()
	i := 0
	for row := range board.Rows {
		for col := range board.Cols {
			if boards.CHAR_TO_CELL[rune(str[i])] {
				u.Set(int(row), int(col), true)
			}
			i++
		}
	}
	return u, nil
}
