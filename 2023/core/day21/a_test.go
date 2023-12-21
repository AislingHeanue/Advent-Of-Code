package day21

import (
	"fmt"
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`)

	result := partA(input)

	require.Equal(t, 16, result)
	util.AwaitClosure()
}

func TestBWithInput(t *testing.T) {
	t.Parallel()
	input := core.FromFile()
	result := partB(input)
	fmt.Println(result)
}
