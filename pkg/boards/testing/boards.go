package boards_testing

import (
	"testing"

	"github.com/stretchr/testify/require"
	"minmax.uk/game-of-life/pkg/boards"
)

func EqualBoards(t *testing.T, want, got string) {
	t.Helper()
	n_want := boards.NormalizeBoard(want)
	n_got := boards.NormalizeBoard(got)
	require.Equal(t, n_want, n_got, "want:\n%s\n====\ngot:\n%s\n", want, got)
}

func NotEqualBoards(t testing.TB, want, got string) {
	t.Helper()
	n_want := boards.NormalizeBoard(want)
	n_got := boards.NormalizeBoard(got)
	require.NotEqual(t, n_want, n_got, "don't want:\n%s\n====\ngot:\n%s\n", want, got)
}
