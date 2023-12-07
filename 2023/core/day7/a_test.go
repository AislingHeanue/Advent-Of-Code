package day7

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	t.Parallel()
	input := core.FromLiteral(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`)

	result := partA(input)

	require.Equal(t, 6440, result)
}
