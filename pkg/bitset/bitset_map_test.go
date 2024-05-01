package bitset_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"minmax.uk/game-of-life/pkg/bitset"
)

func Test_Bitsetmap(t *testing.T) {
	testCases := []struct {
		name  string
		slice [bitset.BS_SIZE]bool
		want  bitset.CellState
	}{
		{
			name: "empty",
			slice: [bitset.BS_SIZE]bool{
				false, false, false,
				false, false, false,
				false, false, false,
			},
			want: bitset.CELL_DEAD,
		},
		{
			name: "diag",
			slice: [bitset.BS_SIZE]bool{
				true, false, false,
				false, true, false,
				false, false, true,
			},
			want: bitset.CELL_LIVE,
		},
		{
			name: "top_row",
			slice: [bitset.BS_SIZE]bool{
				true, true, true,
				false, false, false,
				false, false, false,
			},
			want: bitset.CELL_LIVE,
		},
		{
			name: "top_bottom",
			slice: [bitset.BS_SIZE]bool{
				false, true, false,
				false, false, false,
				false, true, false,
			},
			want: bitset.CELL_DEAD,
		},
		{
			name: "mm&fm",
			slice: [bitset.BS_SIZE]bool{
				false, false, false,
				false, true, true,
				false, false, false,
			},
			want: bitset.CELL_DEAD,
		},
		{
			name: "bu&mu",
			slice: [bitset.BS_SIZE]bool{
				true, true, false,
				false, false, false,
				false, false, false,
			},
			want: bitset.CELL_DEAD,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			b := bitset.FromBoolSlice(test.slice)
			got := b.NextValue()
			require.Equal(t, test.want, got, "board:\n%s\n", b.Repr())
		})
	}
}
