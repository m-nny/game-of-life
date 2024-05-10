package halflife

import (
	"fmt"
	"strings"

	"minmax.uk/game-of-life/pkg/engine/halflife/cell"
)

type Universe struct {
	size int
	Root *cell.MacroCell
}

func (u *Universe) BoardString() string {
	return strings.Join(u.Root.BoardStrings(), "\n")
}

func (u *Universe) DebugPrint() {
	fmt.Printf("universe: %+v\n", u)
	fmt.Printf("%s\n", u.BoardString())
	// fmt.Printf("cache: %+v\n", cell.cell_cache)
	u.Root.PrintDebug("", true)
	fmt.Println()
}

// BuildUniverse builds universe of size 2**level by 2**level
func BuildUniverse(level int) *Universe {
	root := cell.EmptyTree(level)
	return &Universe{
		size: 1 << level,
		Root: root,
	}
}

func (m *Universe) Set(row, col int, value bool) {
	m.Root = m.Root.Set(row, col, value)
}
