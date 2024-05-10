package cell

import (
	"minmax.uk/game-of-life/pkg/utils"
)

// _  _  _  _ | _  _  _  _
// _  _  _  _ | _  _  _  _
// _  _  O  O | O  O  _  _
// _  _  O  O | O  O  _  _
// -----------+-----------
// _  _  O  O | O  O  _  _
// _  _  O  O | O  O  _  _
// _  _  _  _ | _  _  _  _
// _  _  _  _ | _  _  _  _
// center() returns m.level - 1 center Cell
func (m *MacroCell) center() *MacroCell {
	return createCell(
		m.up_left.down_right, m.up_right.down_left,
		m.down_left.up_right, m.down_right.up_left,
	)
}

// _  _  _  _ | _  _  _  _
// _  _  _  _ | _  _  _  _
// _  _  _  _ | _  _  _  _
// _  _  _  O | O  _  _  _
// ---------- + ----------
// _  _  _  O | O  _  _  _
// _  _  _  _ | _  _  _  _
// _  _  _  _ | _  _  _  _
// _  _  _  _ | _  _  _  _
// center() returns m.level - 2 center Cell
func (m *MacroCell) superCenter() *MacroCell {
	return createCell(
		m.up_left.down_right.down_right, m.up_right.down_left.down_left,
		m.down_left.up_right.up_right, m.down_right.up_left.up_left,
	)
}

// _  _  _  _ | _  _  _  _
// _  _  _  O | O  _  _  _
// _  _  _  O | O  _  _  _
// _  _  _  _ | _  _  _  _
// betweenHorizontalCell returns left.level - 1 center Cell
func betweenHorizontalCell(left *MacroCell, right *MacroCell) *MacroCell {
	return createCell(
		left.up_right.down_right, right.up_left.down_left,
		left.down_right.up_right, right.down_left.up_left,
	)
}

// _  _  _  _
// _  _  _  _
// _  _  _  _
// _  O  O  _
// ----------
// _  O  O  _
// _  _  _  _
// _  _  _  _
// _  _  _  _
// betweenHorizontalCell returns up.level center Cell
func betweenVerticalCell(up *MacroCell, down *MacroCell) *MacroCell {
	return createCell(
		up.down_left.down_right, up.down_right.down_left,
		down.up_left.up_right, down.up_right.up_left,
	)
}

func (m *MacroCell) Iterate() *MacroCell {
	if m.level == 2 {
		return m.slowIterate()
	}
	// fmt.Printf("Iterate: %+v\n", m)
	// m.PrintBoard()
	n00 := m.up_left.center()
	n01 := betweenHorizontalCell(m.up_left, m.up_right)
	n02 := m.up_right.center()
	n10 := betweenVerticalCell(m.up_left, m.down_left)
	n11 := m.superCenter()
	n12 := betweenVerticalCell(m.up_right, m.down_right)
	n20 := m.down_left.center()
	n21 := betweenHorizontalCell(m.down_left, m.down_right)
	n22 := m.down_right.center()

	return createCell(
		createCell(n00, n01, n10, n11).Iterate(), createCell(n01, n02, n11, n12).Iterate(),
		createCell(n10, n11, n20, n21).Iterate(), createCell(n11, n12, n21, n22).Iterate(),
	)
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
	res := createCell(
		b.UpLeft().createCell(), b.UpRight().createCell(),
		b.DownLeft().createCell(), b.DownRight().createCell(),
	)
	utils.Assert(res.level == m.level-1, "result should be m.level - 1")
	return res
}
