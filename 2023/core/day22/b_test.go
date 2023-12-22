package day22

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`)

	result := partB(input)

	require.Equal(t, 7, result)
	util.AwaitClosure()
}
