package engine

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"minmax.uk/game-of-life/pkg/boards"
)

type Styles struct {
	Underpopulated lipgloss.Style
	Fine           lipgloss.Style
	Overpopulated  lipgloss.Style
	Baby           lipgloss.Style
}

func DefaultStyles() (s Styles) {
	s.Underpopulated = lipgloss.NewStyle().
		Foreground(lipgloss.Color("202"))
	s.Fine = lipgloss.NewStyle()
	s.Overpopulated = lipgloss.NewStyle().
		Foreground(lipgloss.Color("196"))
	s.Baby = lipgloss.NewStyle().
		Foreground(lipgloss.Color("40"))
	return s
}

func (styles *Styles) Cell(hasLife bool) string {
	return boards.CELL_TO_CHAR[hasLife]
}

func (styles *Styles) CellColor(hasLife bool, n_board int) string {
	r := boards.CELL_TO_CHAR[hasLife]
	style := styles.Fine
	if hasLife {
		if n_board < 2 {
			style = styles.Underpopulated
		} else if 2 <= n_board && n_board <= 3 {
			style = styles.Baby
		} else if 4 <= n_board {
			style = styles.Overpopulated
		}
	} else {
		if n_board == 3 {
			style = styles.Baby
		}
	}
	return style.Render(r)
}

var _defaultStyle = DefaultStyles()

func (state *Engine) String() string {
	if state == nil {
		return "<empty>"
	}
	var builder strings.Builder
	for row := range state.Rows {
		for col := range state.Cols {
			i := col + row*state.Cols
			cell_str := _defaultStyle.Cell(state.cells[i])
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
