package day20

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`)
	result := partA(input)

	require.Equal(t, 32000000, result)
	util.AwaitClosure()
}

func TestA2(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`)

	result := partA(input)

	require.Equal(t, 11687500, result)
	util.AwaitClosure()
}
