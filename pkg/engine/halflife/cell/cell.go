package cell

import (
	"minmax.uk/game-of-life/pkg/datastructs/hashset"
)

var _ hashset.Hashable = (*MacroCell)(nil)

type MacroCell struct {
	up_left    *MacroCell
	up_right   *MacroCell
	down_left  *MacroCell
	down_right *MacroCell

	value bool
	level int
}

func (m *MacroCell) Hash() hashset.Hash {
	if m == nil {
		panic("m is nil")
	}
	// TODO: this is bad hash
	return uint64(m.level)
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
	cur := &MacroCell{level: level}
	if level > 0 {
		child := EmptyTree(level - 1)
		cur.up_left = child
		cur.up_right = child
		cur.down_left = child
		cur.down_right = child
	}
	return cur.normalize()
}
