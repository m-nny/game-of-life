package bitset_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"minmax.uk/game-of-life/pkg/bitset"
)

func Test_Bitset(t *testing.T) {
	b := bitset.Bitset(0)
	{
		got := b.Repr()
		want := "...\n...\n..."
		require.Equalf(t, want, got, "want:\n%s\ngot:\n%s", want, got)
	}

	{
		b.SetForward(true, false, false)
		got := b.Repr()
		want := "..#\n...\n..."
		require.Equalf(t, want, got, "want:\n%s\ngot:\n%s", want, got)
	}

	{
		b.Shift()
		got := b.Repr()
		want := ".#.\n...\n..."
		require.Equalf(t, want, got, "want:\n%s\ngot:\n%s", want, got)
	}

	{
		b.SetForward(false, true, false)
		got := b.Repr()
		want := ".#.\n..#\n..."
		require.Equalf(t, want, got, "want:\n%s\ngot:\n%s", want, got)
	}

	{
		got := b.HasMid()
		want := false
		require.Equalf(t, want, got, "want:\n%v\ngot:\n%v", want, got)
	}

	{
		b.Shift()
		got := b.Repr()
		want := "#..\n.#.\n..."
		require.Equalf(t, want, got, "want:\n%s\ngot:\n%s", want, got)
	}

	{
		got := b.HasMid()
		want := true
		require.Equalf(t, want, got, "want:\n%v\ngot:\n%v", want, got)
	}

	{
		b.SetForward(false, false, true)
		got := b.Repr()
		want := "#..\n.#.\n..#"
		require.Equalf(t, want, got, "want:\n%s\ngot:\n%s", want, got)
	}
}

func Test_FromBoolSlice(t *testing.T) {
	testCases := []struct {
		name  string
		slice [bitset.BS_SIZE]bool
		want  string
	}{
		{
			name: "empty",
			slice: [bitset.BS_SIZE]bool{
				false, false, false,
				false, false, false,
				false, false, false,
			},
			want: "...\n...\n...",
		},
		{
			name: "diag",
			slice: [bitset.BS_SIZE]bool{
				true, false, false,
				false, true, false,
				false, false, true,
			},
			want: "#..\n.#.\n..#",
		},
		{
			name: "top",
			slice: [bitset.BS_SIZE]bool{
				true, true, true,
				false, false, false,
				false, false, false,
			},
			want: "###\n...\n...",
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			b := bitset.FromBoolSlice(test.slice)
			got := b.Repr()
			require.Equalf(t, test.want, got, "want:\n%s\ngot:\n%s", test.want, got)
		})
	}
}
