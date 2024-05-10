package halflife

import (
	"fmt"
	"strings"

	"minmax.uk/game-of-life/pkg/engine/halflife/cell"
)

type Universe struct {
	size int
	root *cell.MacroCell
}

func (u *Universe) BoardString() string {
	return strings.Join(u.root.BoardStrings(), "\n")
}

func (u *Universe) DebugPrint() {
	fmt.Printf("universe: %+v\n", u)
	fmt.Printf("%s\n", u.BoardString())
	// fmt.Printf("cache: %+v\n", cell.cell_cache)
	u.root.PrintDebug("", true)
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

func (m *Universe) Iterate() {
	m.root = m.root.Iterate()
}
