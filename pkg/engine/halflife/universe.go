package halflife

import "strings"

type Universe struct {
	size int
	root *MacroCell
}

func (u *Universe) BoardString() string {
	return strings.Join(u.root.BoardStrings(), "\n")
}

// BuildUniverse builds universe of size 2**level by 2**level
func BuildUniverse(level int) *Universe {
	root := emptyCell(level)
	return &Universe{
		size: 1 << level,
		root: root,
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
