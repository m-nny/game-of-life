package cell

import (
	"fmt"
	"strings"

	"minmax.uk/game-of-life/pkg/boards"
)

func (m *MacroCell) BoardStrings() []string {
	if m == nil {
		return []string{"<nil>"}

	}
	if m.level == 0 {
		return []string{boards.CELL_TO_CHAR[m.value]}
	}
	up := addToRight(m.up_left.BoardStrings(), m.up_right.BoardStrings())
	down := addToRight(m.down_left.BoardStrings(), m.down_right.BoardStrings())
	return addToDown(up, down)
}

func (m *MacroCell) PrintBoard() {
	fmt.Println(strings.Join(m.BoardStrings(), "\n"))
}

func (m *MacroCell) PrintTree(prefix string, rec bool) {
	if m == nil {
		return
	}
	fmt.Printf("%s*%p %+v\n", prefix, m, m)
	if rec {
		m.up_left.PrintTree(prefix+" ", m.up_left != m.up_right)
		m.up_right.PrintTree(prefix+" ", m.up_right != m.down_left)
		m.down_left.PrintTree(prefix+" ", m.down_left != m.down_right)
		m.down_right.PrintTree(prefix+" ", true)
	}
}
