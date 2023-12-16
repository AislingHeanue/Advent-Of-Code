package day15

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
	input := core.FromLiteral(`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`)

	result := partA(input)

	require.Equal(t, 1320, result)
	util.AwaitClosure()
}

func TestDecode(t *testing.T) {
	for _, letter := range "abcdefghijklmnopqrstuvwxyz-=123456789" {
		fmt.Println(int(letter))
	}
}

func TestHash(t *testing.T) {
	require.Equal(t, HASH("rn"), 0)
}
