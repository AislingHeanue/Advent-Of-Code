package day3

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	t.Parallel()
	input := core.FromLiteral(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..	
`)
	result := partA(input)

	require.Equal(t, 4361, result)
}
