package engine

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	boards_testing "minmax.uk/game-of-life/pkg/boards/testing"
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
			want: []string{"1b", "1o", "1b", "2b", "1o", "3o"},
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

func Test_FromRLE(t *testing.T) {
	testCases := []struct {
		name string
		rle  string
		want string
	}{
		{
			name: "glider",
			rle:  "x = 3, y = 3\nbob$2bo$3o!",
			want: `
			.O.
			..O
			OOO
`,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			r := strings.NewReader(test.rle)
			b, err := FromRLE(r)
			require.NoError(t, err)

			boards_testing.EqualBoards(t, test.want, b.String())
		})
	}
}
