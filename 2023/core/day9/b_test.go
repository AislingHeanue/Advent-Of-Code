package day9

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Parallel()
	input := core.FromLiteral(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`)

	result := partB(input)

	require.Equal(t, 2, result)
}
