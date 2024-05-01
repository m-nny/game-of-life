package bitset

import (
	"iter"
	"strings"
)

// the bitmap is arranged with
// the forward cells in bits 6-8,
// the middle cells in bits 3-5,
// and the rear cells in bits 0-2
//
// board:
//
//	ru mu fu
//	rm mm fm
//	rb mb fb
//
// int from least significat bit to most
//
//	...  8  7  6  5  4  3  2  1  0
//	... fb fm fu mb mm mu rb rm ru

type Bitset int32

const (
	BS_SIZE         = 9
	ALL_BS_SET_SIZE = 1 << BS_SIZE

	BS_RU_BIT = 1 << 0
	BS_RM_BIT = 1 << 1
	BS_RB_BIT = 1 << 2

	BS_MU_BIT = 1 << 3
	BS_MM_BIT = 1 << 4
	BS_MB_BIT = 1 << 5

	BS_FU_BIT = 1 << 6
	BS_FM_BIT = 1 << 7
	BS_FB_BIT = 1 << 8
)

func (b *Bitset) HasMid() bool {
	return ((*b) & BS_MM_BIT) > 0
}

func (b *Bitset) HasBit(i int) bool {
	return ((*b) & (1 << i)) > 0
}

func (b *Bitset) SetForward(up, mid, bot bool) {
	if up {
		*b |= BS_FU_BIT
	}
	if mid {
		*b |= BS_FM_BIT
	}
	if bot {
		*b |= BS_FB_BIT
	}
}

func (b *Bitset) Shift() {
	(*b) >>= 3
}

var reprOrder = [BS_SIZE]int{
	BS_RU_BIT, BS_MU_BIT, BS_FU_BIT,
	BS_RM_BIT, BS_MM_BIT, BS_FM_BIT,
	BS_RB_BIT, BS_MB_BIT, BS_FB_BIT,
}

func (b *Bitset) Repr() string {
	var builder strings.Builder
	for i, bit := range reprOrder {
		if (int(*b) & bit) > 0 {
			builder.WriteRune('#')
		} else {
			builder.WriteRune('.')
		}
		if i%3 == 2 && i+1 != len(reprOrder) {
			builder.WriteRune('\n')
		}
	}
	return builder.String()
}

func AllBitsets() iter.Seq[Bitset] {
	return func(yield func(Bitset) bool) {
		for b := Bitset(0); b < ALL_BS_SET_SIZE; b++ {
			if !yield(b) {
				return
			}
		}
	}
}

func FromBoolSlice(s [BS_SIZE]bool) Bitset {
	b := Bitset(0)
	for i, item := range reprOrder {
		if s[i] {
			b |= Bitset(item)
		}
	}
	return b
}
