package hashset_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"minmax.uk/game-of-life/pkg/datastructs/hashset"
)

type IntItem struct {
	val uint64
}

func (i IntItem) Equals(other hashset.Hashable) bool {
	if other, ok := other.(IntItem); ok {
		return i.val == other.val
	}
	return false
}

func (i IntItem) Hash() hashset.Hash {
	return i.val
}

func Test_Simple(t *testing.T) {
	h := hashset.New[IntItem]()

	{
		one := IntItem{1}
		require.False(t, h.Contains(one))
		h.Add(one)
		require.True(t, h.Contains(one))
	}

	{
		other_one := IntItem{1}
		require.True(t, h.Contains(other_one))
		h.Add(other_one)
		require.True(t, h.Contains(other_one))
	}

	{
		two := IntItem{2}
		require.False(t, h.Contains(two))
		h.Add(two)
		require.True(t, h.Contains(two))
	}
}

type BadHashItem struct {
	val uint64
}

func (i BadHashItem) Equals(other hashset.Hashable) bool {
	if other, ok := other.(BadHashItem); ok {
		return i.val == other.val
	}
	return false
}

func (i BadHashItem) Hash() hashset.Hash {
	return 0
}

func Test_BadHash(t *testing.T) {
	h := hashset.New[BadHashItem]()

	{
		one := BadHashItem{1}
		require.False(t, h.Contains(one))
		h.Add(one)
		require.True(t, h.Contains(one))
	}

	{
		other_one := BadHashItem{1}
		require.True(t, h.Contains(other_one))
		h.Add(other_one)
		require.True(t, h.Contains(other_one))
	}

	{
		two := BadHashItem{2}
		require.False(t, h.Contains(two))
		h.Add(two)
		require.True(t, h.Contains(two))
	}
}
