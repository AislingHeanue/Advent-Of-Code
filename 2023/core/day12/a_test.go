package day12

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`)

	result := partA(input)

	require.Equal(t, 21, result)
	util.AwaitClosure()
}

func TestASmall(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`?#?#?#?#?#?#?#? 1,3,1,6`)
	result := partA(input)
	require.Equal(t, 1, result)
}

func TestASmaller(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`???? 1`)
	result := partA(input)
	require.Equal(t, 4, result)
}

func TestAEdge(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`..???#??????#?????? 3,6`)
	result := partA(input)
	require.Equal(t, 15, result)
}

func TestAEdge2(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`..???#??????#?????? 3`)
	result := partA(input)
	require.Equal(t, 0, result)
}
