package engine

import (
	"strings"

	"minmax.uk/game-of-life/pkg/boards"
)

func (state *Engine) String() string {
	if state == nil {
		return "<empty>"
	}
	var builder strings.Builder
	for row := range state.Rows {
		for col := range state.Cols {
			i := col + row*state.Cols
			cell_str := boards.DefaultStyles.Cell(state.cells[i])
			builder.WriteString(cell_str)
		}
		if row+1 < state.Rows {
			builder.WriteRune('\n')
		}
	}
	return builder.String()
}

// func (state *Engine) ColoredString() string {
// 	if state == nil {
// 		return "<empty>"
// 	}
// 	n_board := state.calcNboard()
// 	var builder strings.Builder
// 	for row := range state.Rows {
// 		for col := range state.Cols {
// 			i := col + row*state.Cols
// 			cell_str := _defaultStyle.CellColor(state.cells[i], n_board[i])
// 			builder.WriteString(cell_str)
// 		}
// 		if row+1 < state.Rows {
// 			builder.WriteRune('\n')
// 		}
// 	}
// 	return builder.String()
// }
