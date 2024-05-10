package cell

import (
	"minmax.uk/game-of-life/pkg/datastructs/hashset"
	// "minmax.uk/game-of-life/pkg/utils"
)

var _ hashset.Hashable = (*MacroCell)(nil)

type MacroCell struct {
	up_left    *MacroCell
	up_right   *MacroCell
	down_left  *MacroCell
	down_right *MacroCell

	next *MacroCell

	value bool
	level int
	hash  hashset.Hash
}

func createLeaf(value bool) *MacroCell {
	h := hashset.Hash(0)
	if value {
		h++
	}
	return (&MacroCell{value: value, hash: h}).normalize()
}

func createCell(up_left, up_right, down_left, down_right *MacroCell) *MacroCell {
	// utils.Assert(up_left != nil && up_right != nil && down_left != nil && down_right != nil, "subcells cannot not be nil: %v %v %v %v", up_left, up_right, down_left, down_right)
	// utils.Assert(up_left.level == up_right.level && up_left.level == down_left.level && up_left.level == down_right.level, "subcells should have same level: %v %v %v %v", up_left, up_right, down_left, down_right)
	h := 3*up_left.hash + 101*up_right.hash + 1001*down_left.hash + 1007*down_right.hash + 1000007*uint64(up_right.level+1)
	return (&MacroCell{up_left: up_left, up_right: up_right, down_left: down_left, down_right: down_right, level: up_left.level + 1, hash: h}).normalize()
}

func (m *MacroCell) Hash() hashset.Hash {
	return m.hash
}

func (m *MacroCell) Same(other *MacroCell) bool {
	if m == nil || other == nil {
		return m == other
	}
	return m.level == other.level &&
		m.value == other.value &&
		m.up_left == other.up_left &&
		m.up_right == other.up_right &&
		m.down_left == other.down_left &&
		m.down_right == other.down_right
}

func (m *MacroCell) Equals(other hashset.Hashable) bool {
	if other, ok := other.(*MacroCell); ok {
		return m.Same(other)
	}
	return false
}

var cell_cache = hashset.New[*MacroCell]()

func (m *MacroCell) normalize() *MacroCell {
	existing, ok := cell_cache.Get(m)
	if ok {
		return existing
	}
	cell_cache.Add(m)
	return m
}

func (m MacroCell) Set(row, col int, value bool) *MacroCell {
	w := 1 << m.level
	if !(0 <= row && row < w && 0 <= col && col < w) {
		panic("invalid pos")
	}
	if m.level == 0 {
		m.value = value
		return (&m).normalize()
	}
	mid := w >> 1
	if row < mid && col < mid {
		m.up_left = m.up_left.Set(row, col, value)
	} else if row < mid && mid <= col {
		m.up_right = m.up_right.Set(row, col-mid, value)
	} else if mid <= row && col < mid {
		m.down_left = m.down_left.Set(row-mid, col, value)
	} else if mid <= row && mid <= col {
		m.down_right = m.down_right.Set(row-mid, col-mid, value)
	} else {
		panic("should not be here")
	}
	return (&m).normalize()
}

func (m *MacroCell) Get(row, col int) bool {
	w := 1 << m.level
	if !(0 <= row && row < w && 0 <= col && col < w) {
		panic("invalid pos")
	}
	if m.level == 0 {
		return m.value
	}
	mid := w >> 1
	if row < mid && col < mid {
		return m.up_left.Get(row, col)
	} else if row < mid && mid <= col {
		return m.up_right.Get(row, col-mid)
	} else if mid <= row && col < mid {
		return m.down_left.Get(row-mid, col)
	} else if mid <= row && mid <= col {
		return m.down_right.Get(row-mid, col-mid)
	} else {
		panic("should not be here")
	}
}

func EmptyTree(level int) *MacroCell {
	if level == 0 {
		return createLeaf(false)
	}
	child := EmptyTree(level - 1)
	return createCell(child, child, child, child)
}

// _  _  _  _ | _  _  _  _
// _  _  _  _ | _  _  _  _
// _  _  O  O | O  O  _  _
// _  _  O  O | O  O  _  _
// -----------+-----------
// _  _  O  O | O  O  _  _
// _  _  O  O | O  O  _  _
// _  _  _  _ | _  _  _  _
// _  _  _  _ | _  _  _  _
func (m *MacroCell) Expand() *MacroCell {
	emptySpace := EmptyTree(m.level - 1)
	return createCell(
		createCell(
			emptySpace, emptySpace,
			emptySpace, m.up_left,
		),
		createCell(
			emptySpace, emptySpace,
			m.up_right, emptySpace,
		),
		createCell(
			emptySpace, m.down_left,
			emptySpace, emptySpace,
		),
		createCell(
			m.down_right, emptySpace,
			emptySpace, emptySpace,
		),
	)
}

func PrintStats() {
	cell_cache.PrintStats()
}
func ResetStats() {
	cell_cache.ResetStats()
}
