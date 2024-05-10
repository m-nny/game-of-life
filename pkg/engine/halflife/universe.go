package halflife

import (
	"fmt"
	"strings"
)

type Universe struct {
	size int
	Root *MacroCell
}

func (u *Universe) BoardString() string {
	return strings.Join(u.Root.BoardStrings(), "\n")
}

func (u *Universe) DebugPrint() {
	fmt.Printf("universe: %+v\n", u)
	fmt.Printf("%s\n", u.BoardString())
	fmt.Printf("cache: %+v\n", Cell_cache)
	u.Root.PrintDebug("", true)
	fmt.Println()
}

// BuildUniverse builds universe of size 2**level by 2**level
func BuildUniverse(level int) *Universe {
	root := emptyCell(level)
	return &Universe{
		size: 1 << level,
		Root: root,
	}
}

func emptyCell(level int) *MacroCell {
	cur := &MacroCell{level: level}
	if level > 0 {
		child := emptyCell(level - 1)
		cur.up_left = child
		cur.up_right = child
		cur.down_left = child
		cur.down_right = child
	}
	return cur.Normalize()
}

func (m *Universe) Set(row, col int, value bool) {
	m.Root = m.Root.Set(row, col, value)
}
