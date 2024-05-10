package halflife

import (
	"fmt"

	"minmax.uk/game-of-life/pkg/boards"
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

func (m *MacroCell) BoardStrings() []string {
	if m.level == 0 {
		return []string{boards.CELL_TO_CHAR[m.value]}
	}
	up := addToRight(m.up_left.BoardStrings(), m.up_right.BoardStrings())
	down := addToRight(m.down_left.BoardStrings(), m.down_right.BoardStrings())
	return addToDown(up, down)
}

func (m *MacroCell) PrintDebug(prefix string, rec bool) {
	if m == nil {
		return
	}
	fmt.Printf("%s*%p %+v\n", prefix, m, m)
	if rec {
		m.up_left.PrintDebug(prefix+" ", m.up_left != m.up_right)
		m.up_right.PrintDebug(prefix+" ", m.up_right != m.down_left)
		m.down_left.PrintDebug(prefix+" ", m.down_left != m.down_right)
		m.down_right.PrintDebug(prefix+" ", true)
	}
}

var Cell_cache = hashset.New[*MacroCell]()

func (m *MacroCell) Normalize() *MacroCell {
	existing, ok := Cell_cache.Get(m)
	if ok {
		return existing
	}
	Cell_cache.Add(m)
	return m
}

func (m MacroCell) Set(row, col int, value bool) *MacroCell {
	w := 1 << m.level
	if !(0 <= row && row < w && 0 <= col && col < w) {
		panic("invalid pos")
	}
	if m.level == 0 {
		m.value = value
		return (&m).Normalize()
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
	return (&m).Normalize()
}
