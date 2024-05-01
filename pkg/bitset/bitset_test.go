package bitset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Bitset(t *testing.T) {
	b := bitset(0)
	{
		got := b.Repr()
		want := "...\n...\n..."
		fmt.Printf("b: %03o %03d\n", b, b)
		require.Equalf(t, want, got, "want:\n%s\ngot:\n%s", want, got)
	}

	{
		b.SetForward(true, false, false)
		got := b.Repr()
		want := "..#\n...\n..."
		fmt.Printf("b: %03o %03d\n", b, b)
		require.Equalf(t, want, got, "want:\n%s\ngot:\n%s", want, got)
	}

	{
		b.Shift()
		got := b.Repr()
		want := ".#.\n...\n..."
		fmt.Printf("b: %03o %03d\n", b, b)
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
		fmt.Printf("b: %03o %03d\n", b, b)
		require.Equalf(t, want, got, "want:\n%v\ngot:\n%v", want, got)
	}

	{
		b.Shift()
		got := b.Repr()
		want := "#..\n.#.\n..."
		fmt.Printf("b: %03o %03d\n", b, b)
		require.Equalf(t, want, got, "want:\n%s\ngot:\n%s", want, got)
	}

	{
		got := b.HasMid()
		want := true
		fmt.Printf("b: %03o %03d\n", b, b)
		require.Equalf(t, want, got, "want:\n%v\ngot:\n%v", want, got)
	}

	{
		b.SetForward(false, false, true)
		got := b.Repr()
		want := "#..\n.#.\n..#"
		require.Equalf(t, want, got, "want:\n%s\ngot:\n%s", want, got)
	}
}
