package bitset_engine

import (
	"testing"

	"github.com/stretchr/testify/require"
	"minmax.uk/game-of-life/pkg/boards"
	boards_testing "minmax.uk/game-of-life/pkg/boards/testing"
)

func TestOsclilators(t *testing.T) {
	t.Helper()
	oscilators := []boards.BoardSpec{
		boards.Blinker,
		boards.Block,
		boards.Beehive,
		boards.Toad,
		boards.Pulsar,
	}
	for _, spec := range oscilators {
		t.Run(spec.Name, func(t *testing.T) {
			state, err := FromBoardSpec(spec)
			require.NoError(t, err)
			boards_testing.EqualBoards(t, spec.Str, state.String())
			for range spec.Freq {
				prevBoard := state.String()
				state.Iterate()
				boards_testing.NotEqualBoards(t, prevBoard, state.String())
			}
			// For non changing oscilators
			if spec.Freq == 0 {
				state.Iterate()
			}
			boards_testing.EqualBoards(t, spec.Str, state.String())
		})
	}
}

func Test_Next_4Rules(t *testing.T) {
	oscilators := []struct {
		name       string
		board      boards.BoardSpec
		want_state string
	}{
		{
			name: "n <= 1",
			board: boards.BoardSpec{
				Rows: 5,
				Cols: 5,
				Str: `
				.....
				.OO..
				.....
				.....
				.....
			`,
			},
			want_state: `
				.....
				.....
				.....
				.....
				.....
			`,
		},
		{
			name: "2 <= n <= 3",
			board: boards.BoardSpec{
				Rows: 5,
				Cols: 5,
				Str: `
				.....
				.OO..
				.O...
				.....
				.....
			`,
			},
			want_state: `
				.....
				.OO..
				.OO..
				.....
				.....
			`,
		},
		{
			name: "4 <= n",
			board: boards.BoardSpec{
				Rows: 9,
				Cols: 9,
				Str: `
				.........
				.........
				.........
				...OOO...
				...OOO...
				...OOO...
				.........
				.........
				.........
			`,
			},
			want_state: `
				.........
				.........
				....O....
				...O.O...
				..O...O..
				...O.O...
				....O....
				.........
				.........
			`,
		},
	}
	for _, test := range oscilators {
		t.Run(test.name, func(t *testing.T) {
			require := require.New(t)
			initState, err := FromBoardSpec(test.board)
			require.NoError(err)
			boards_testing.EqualBoards(t, test.board.Str, initState.String())

			initState.Iterate()
			boards_testing.EqualBoards(t, test.want_state, initState.String())
		})
	}
}

// func Benchmark_Iterate(b *testing.B) {
// 	seed := int64(42)
// 	bench_cases := []struct {
// 		Cols int64
// 		Rows int64
// 	}{
// 		{100, 100},
// 		{1000, 100},
// 		{1000, 1000},
// 		{10000, 1000},
// 	}
// 	for _, test := range bench_cases {
// 		benchBoard := boards.Random(test.Cols, test.Rows, seed)
// 		state, err := FromBoardSpec(benchBoard)
// 		require.NoError(b, err)
// 		test_name := fmt.Sprintf("%dx%d", test.Rows, test.Cols)
// 		b.ResetTimer()
// 		b.Run(test_name, func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				state.Iterate()
// 			}
// 		})
// 	}
// }
