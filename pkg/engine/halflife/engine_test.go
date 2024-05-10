package halflife

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
			universe, err := FromBoardSpec(spec)
			require.NoError(t, err)
			boards_testing.EqualBoards(t, spec.Str, universe.String())
			for range spec.Freq {
				prevBoard := universe.String()
				universe.Iterate()
				boards_testing.NotEqualBoards(t, prevBoard, universe.String())
			}
			// For non changing oscilators
			if spec.Freq == 0 {
				universe.Iterate()
			}
			boards_testing.EqualBoards(t, spec.Str, universe.String())
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
