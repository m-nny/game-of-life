package bitset

type CellState bool

var (
	CELL_DEAD CellState = false
	CELL_LIVE CellState = true
)

type bitsetmap [ALL_BS_SET_SIZE]CellState

var _bitsetmap bitsetmap

func BuildBitsetmap() bitsetmap {
	// bitsetmap := make([ALL_BS_SET_SIZE]cellState, ALL_BS_SET_SIZE)
	for b := range AllBitsets() {
		n := 0
		for i := 0; i < BS_SIZE; i++ {
			if b.HasBit(i) {
				n++
			}
		}
		if b.HasMid() {
			if 2 <= n && n <= 3 {
				_bitsetmap[b] = CELL_LIVE
			} else {
				_bitsetmap[b] = CELL_DEAD
			}
		} else {
			if n == 2 {
				_bitsetmap[b] = CELL_LIVE
			} else {
				_bitsetmap[b] = CELL_DEAD
			}
		}
	}
	return _bitsetmap
}

func (b *Bitset) NextValue() CellState {
	return _bitsetmap[*b]
}
