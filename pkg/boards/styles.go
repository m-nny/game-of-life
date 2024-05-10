package boards

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Underpopulated lipgloss.Style
	Fine           lipgloss.Style
	Overpopulated  lipgloss.Style
	Baby           lipgloss.Style
}

func getDefaultStyles() (s Styles) {
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
	return CELL_TO_CHAR[hasLife]
}

func (styles *Styles) CellColor(hasLife bool, n_board int) string {
	r := CELL_TO_CHAR[hasLife]
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

var DefaultStyles = getDefaultStyles()
