package cell

import (
	"minmax.uk/game-of-life/pkg/utils"
)

func (m *MacroCell) Iterate() *MacroCell {
	if m.level == 2 {
		return m.slowIterate()
	}
	utils.Assert(false, "not implemented")
	return m
}

func (m *MacroCell) slowIterate() *MacroCell {
	utils.Assert(m.level == 2, "should slowIterate only for m.level == 2")
	w := 1 << m.level
	b := bitset4(0)
	for row := range w {
		for col := range w {
			if m.Get(row, col) {
				b.Push(1)
			} else {
				b.Push(0)
			}
		}
	}
	return createCell(
		b.UpLeft().createCell(),
		b.UpRight().createCell(),
		b.DownLeft().createCell(),
		b.DownRight().createCell(),
	)
}
