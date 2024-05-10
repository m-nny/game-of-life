package halflife

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addToRight(t *testing.T) {
	// 12+34
	// 56+78
	// ab+cd
	// ef+gh

	left := []string{
		"12",
		"56",
		"ab",
		"ef",
	}

	right := []string{
		"34",
		"78",
		"cd",
		"gh",
	}
	want := []string{
		"1234",
		"5678",
		"abcd",
		"efgh",
	}
	got := addToRight(left, right)
	require.Equal(t, want, got)
}

func Test_addToDown(t *testing.T) {
	// 1234
	// 5678
	// ++++
	// abcd
	// efgh
	up := []string{"1234", "5678"}
	down := []string{"abcd", "efgh"}
	want := []string{"1234", "5678", "abcd", "efgh"}
	got := addToDown(up, down)
	require.Equal(t, want, got)
}
