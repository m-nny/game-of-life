package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"minmax.uk/game-of-life/pkg/utils"
)

func Test_Assert(t *testing.T) {
	require.Panics(t, func() {
		utils.Assert(false, "Fails")
	})
	require.NotPanics(t, func() {
		utils.Assert(true, "Fails")
	})
}
