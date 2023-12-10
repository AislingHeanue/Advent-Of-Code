package day10

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	t.Parallel()
	input := core.FromLiteral(`.....
.S-7.
.|.|.
.L-J.
.....`)

	result := partA(input)

	require.Equal(t, 4, result)
}

func TestA2(t *testing.T) {
	t.Parallel()
	input := core.FromLiteral(`7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`)

	result := partA(input)

	require.Equal(t, 8, result)
}
