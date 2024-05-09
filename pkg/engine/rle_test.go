package engine

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rleSplitter(t *testing.T) {
	testCases := []struct {
		name string
		rle  string
		want []string
	}{
		{
			name: "glider",
			rle:  "bob$2bo$3o!",
			want: []string{"b", "o", "b", "2b", "o", "3o"},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			r := strings.NewReader(test.rle)
			s := bufio.NewScanner(r)
			s.Split(rleSplitter)

			var got []string
			for s.Scan() {
				got = append(got, s.Text())
			}
			require.NoError(t, s.Err())

			require.Equal(t, test.want, got)
		})
	}
}

// func TestFromRLE(t *testing.T) {
// 	testCases := []struct {
// 		name string
// 		rle  string
// 		want string
// 	}{
// 		{
// 			name: "glider",
// 			rle:  "x = 3, y = 3\nbob$2bo$3o!",
// 			want: `
// 			.O.
// 			..O
// 			OOO
// `,
// 		},
// 	}
// 	for _, test := range testCases {
// 		t.Run(test.name, func(t *testing.T) {
// 			r := strings.NewReader(test.rle)
// 			b, err := FromRLE(r)
// 			require.NoError(t, err)
//
// 			EqualBoards(t, test.want, b.String())
// 		})
// 	}
// }
