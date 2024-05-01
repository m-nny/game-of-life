package bitset

type CellState bool

var (
	CELL_DEAD CellState = false
	CELL_LIVE CellState = true
)

type bitsetmap [ALL_BS_SET_SIZE]CellState

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
