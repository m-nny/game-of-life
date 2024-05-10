package bitset

import "iter"

type CellState bool

const ALL_BS_SET_SIZE = 1 << BS_SIZE

var (
	CELL_DEAD CellState = false
	CELL_LIVE CellState = true
)

type bitsetmap [ALL_BS_SET_SIZE]CellState

func AllBitsets() iter.Seq[Bitset] {
	return func(yield func(Bitset) bool) {
		for b := Bitset(0); b < ALL_BS_SET_SIZE; b++ {
			if !yield(b) {
				return
			}
		}
	}
}

var _bitsetmap = func() bitsetmap {
	var bitsetmap bitsetmap
	for b := range AllBitsets() {
		n := 0
		for i := 0; i < BS_SIZE; i++ {
			if b.HasBit(i) {
				n++
			}
		}
		if b.HasMid() {
			// because mid itself is counted
			n--
			if 2 <= n && n <= 3 {
				bitsetmap[b] = CELL_LIVE
			} else {
				bitsetmap[b] = CELL_DEAD
			}
		} else {
			if n == 3 {
				bitsetmap[b] = CELL_LIVE
			} else {
				bitsetmap[b] = CELL_DEAD
			}
		}
	}
	return bitsetmap
}()

func (b *Bitset) NextValue() CellState {
	return _bitsetmap[*b]
}
