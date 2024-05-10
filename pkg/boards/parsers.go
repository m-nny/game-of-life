package boards

import (
	"strings"
)

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
}()

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
