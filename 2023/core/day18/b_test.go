package day18

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`)
	result := partB(input)

	util.AwaitClosure()

	require.Equal(t, 952408144115, result)
}

func TestBSmall(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`R 6 (#000030)
	R 6 (#000021)
	R 6 (#000032)
	R 6 (#000023)`)
	result := partB(input)

	util.AwaitClosure()

	require.Equal(t, 12, result)
}

func TestBMedium(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`R 6 (#000060)
D 5 (#000051)
L 2 (#000022)
D 2 (#000021)
R 2 (#000020)
D 2 (#000021)
L 5 (#000052)
U 2 (#000023)
L 1 (#000012)
U 2 (#000023)
R 2 (#000020)
U 3 (#000033)
L 2 (#000022)
U 2 (#000023)`)
	result := partB(input)

	util.AwaitClosure()

	require.Equal(t, 62, result)

}
