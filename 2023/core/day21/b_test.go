package day21

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`#..
.S.
#.#`)
	input = core.FromFile()

	result := partB(input)

	require.Equal(t, 16, result)
	util.AwaitClosure()
}
