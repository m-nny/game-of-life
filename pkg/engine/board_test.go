package engine

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func EqualBoards(t *testing.T, want, got string) {
	t.Helper()
	n_want := NormalizeBoard(want)
	n_got := NormalizeBoard(got)
	require.Equal(t, n_want, n_got, "want:\n%s\n====\ngot:\n%s\n", want, got)
}

func NotEqualBoards(t testing.TB, want, got string) {
	t.Helper()
	n_want := NormalizeBoard(want)
	n_got := NormalizeBoard(got)
	require.NotEqual(t, n_want, n_got, "don't want:\n%s\n====\ngot:\n%s\n", want, got)
}

func TestOsclilators(t *testing.T) {
	t.Helper()
	oscilators := []BoardSpec{
		Blinker,
		Block,
		Beehive,
		Toad,
		Pulsar,
	}
	for _, spec := range oscilators {
		t.Run(spec.Name, func(t *testing.T) {
			state, err := FromBoardSpec(spec)
			require.NoError(t, err)
			EqualBoards(t, spec.Str, state.String())
			for range spec.Freq {
				prevBoard := state.String()
				state.Iterate()
				NotEqualBoards(t, prevBoard, state.String())
			}
			// For non changing oscilators
			if spec.Freq == 0 {
				state.Iterate()
			}
			EqualBoards(t, spec.Str, state.String())
		})
	}
}

func Test_Next_4Rules(t *testing.T) {
	oscilators := []struct {
		name       string
		rows       int64
		cols       int64
		init_state string
		want_state string
	}{
		{
			name: "n <= 1",
			rows: 5,
			cols: 5,
			init_state: `
				.....
				.OO..
				.....
				.....
				.....
			`,
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
			rows: 5,
			cols: 5,
			init_state: `
				.....
				.OO..
				.O...
				.....
				.....
			`,
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
			rows: 9,
			cols: 9,
			init_state: `
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
			initState, err := FromString(test.cols, test.rows, test.init_state)
			require.NoError(err)
			EqualBoards(t, test.init_state, initState.String())

			initState.Iterate()
			EqualBoards(t, test.want_state, initState.String())
		})
	}
}
