package day6

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Parallel()
	input := core.FromLiteral(`Time:      7  15   30
Distance:  9  40  200`)

	result := partB(input)

	require.Equal(t, 71503, result)
}
