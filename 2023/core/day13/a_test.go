package day13

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`)

	result := partA(input)

	require.Equal(t, 405, result)
	util.AwaitClosure()
}
