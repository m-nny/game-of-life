package naive_engine

import (
	"strings"

	"minmax.uk/game-of-life/pkg/boards"
)

func (state *NaiveEngine) String() string {
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
