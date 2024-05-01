package engine

import (
	"fmt"
	"strings"
)

func FromBoardSpec(board BoardSpec) (*Engine, error) {
	return FromString(board.Rows, board.Cols, board.Str)
}

func FromString(rows, cols int64, str string) (*Engine, error) {
	str = NormalizeBoard(str)
	if len(str) < int(cols*rows) {
		return nil, fmt.Errorf("not enough chars")
	}
	if len(str) > int(cols*rows) {
		return nil, fmt.Errorf("too much chars")
	}
	g := EmptyGame(rows, cols)
	i := int64(0)
	for _, rune := range str {
		if strings.ContainsRune(EMPTY_CELLS, rune) {
			g.cells[i] = false
		} else if strings.ContainsRune(FULL_CELLS, rune) {
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

func MustFromString(rows, cols int64, str string) *Engine {
	g, err := FromString(rows, cols, str)
	if err != nil {
		panic(fmt.Sprintf("FromString() errored: %v", err))
	}
	return g
}

func EmptyGame(rows, cols int64) *Engine {
	cells := make([]bool, cols*rows)
	prevCells := make([]bool, cols*rows)
	return &Engine{
		Cols: cols,
		Rows: rows,

		cells:     cells,
		prevCells: prevCells,
	}
}

var EMPTY_CELLS = "."
var FULL_CELLS = "O#"
var SKIP_CELLS = "\n\t "

var CELL_TO_CHAR = map[bool]string{
	false: ".",
	true:  "O",
}
var CHAR_TO_CELL = func() map[rune]bool {
	m := make(map[rune]bool)
	for _, cell := range EMPTY_CELLS {
		m[cell] = false
	}
	for _, cell := range FULL_CELLS {
		m[cell] = true
	}
	return m
}

var boardUnifier = func() *strings.Replacer {
	var pairs []string
	for _, cell := range SKIP_CELLS {
		pairs = append(pairs, string(cell), "")
	}
	for _, cell := range FULL_CELLS {
		pairs = append(pairs, string(cell), CELL_TO_CHAR[true])
	}
	for _, cell := range EMPTY_CELLS {
		pairs = append(pairs, string(cell), CELL_TO_CHAR[false])
	}
	return strings.NewReplacer(pairs...)
}()

var NormalizeBoard = boardUnifier.Replace
