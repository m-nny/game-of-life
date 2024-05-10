package cell

import (
	"math/bits"
	"strings"

	"minmax.uk/game-of-life/pkg/boards"
)

//	 15 14 13 12
//	 11 10  9  8
//		7  6  5  4
//		3  2  1  0
type bitset4 uint64

// __ __ __ __
// __ 10  9  8
// __  6  5  4
// __  2  1  0
func (b bitset4) Focus() bitset3 {
	return bitset3(
		b & (0x777),
	)
}

// __ __ __ __
// __ 10  9  8
// __  6  5  4
// __  2  1  0
func (b bitset4) DownRight() bitset3 {
	return b.Focus()
}

//	 __ __ __ __
//	 11 10  9 __
//		7  6  5 __
//		3  2  1 __
func (b bitset4) DownLeft() bitset3 {
	return (b >> 1).Focus()
}

// __ 14 13 12
// __ 10  9  8
// __  6  5  4
// __ __ __ __
func (b bitset4) UpRight() bitset3 {
	return (b >> 4).Focus()
}

//	 15 14 13 __
//	 11 10  9 __
//		7  6  5 __
//	  _ __ __ __
func (b bitset4) UpLeft() bitset3 {
	return (b >> 5).Focus()
}

func (b *bitset4) Set(idx int) {
	(*b) |= 1 << idx
}

func (b bitset4) Get(pos int) bool {
	return (b & (1 << pos)) > 0
}

func (b *bitset4) Push(value int) {
	*b = (*b)<<1 | bitset4(value)
}

func (b bitset4) String() string {
	var lines []string
	i := 15
	for range 4 {
		line := ""
		for range 4 {
			line += boards.CELL_TO_CHAR[b.Get(i)]
			i--
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

//	 10  9  8
//		6  5  4
//		2  1  0
type bitset3 uint64

func (b bitset3) Count() int {
	n := bits.OnesCount64(uint64(b))
	if b.HasMid() {
		n--
	}
	return n
}
func (b bitset3) HasMid() bool {
	return b.Get(5)
}

// createCell from downright 3x3 box
func (b bitset3) createCell() *MacroCell {
	n := b.Count()
	new_life := n == 3 || (b.HasMid() && n == 2)
	return createLeaf(new_life)
}

func (b bitset3) Get(pos int) bool {
	return (b & (1 << pos)) > 0
}

func (b bitset3) String() string {
	var lines []string
	i := 15
	for range 4 {
		line := ""
		for range 4 {
			line += boards.CELL_TO_CHAR[b.Get(i)]
			i--
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}
