package day7

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Parallel()
	input := core.FromLiteral(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`)

	result := partB(input)

	require.Equal(t, 5905, result)
}

func TestBSmall(t *testing.T) {
	input := core.FromLiteral(`3JJJJ 1
JJJJJ 10
TTTTT 100
KKKKK 1000
22222 10000`)

	result := partB(input)

	require.Equal(t, 25413, result)
}
