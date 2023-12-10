package day10

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`.....
.S-7.
.|.|.
.L-J.
.....`)

	result := partA(input)

	require.Equal(t, 4, result)
	util.AwaitClosure()
}

func TestA2(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`)

	result := partA(input)

	require.Equal(t, 8, result)
	util.AwaitClosure()
}

func TestDrawPath(t *testing.T) {
	t.Skip("Don't run animations in overall tests")
	Test = true
	input := core.FromFile()
	d := partA(input)
	distance3 := util.NewThreeMatrix[int](d/50, Distances.GetHeight(), Distances.GetWidth())
	distance3.SetByRule(func(z, y, x int) int {
		v := Distances.MustGet(y, x)
		if v > 110*z {
			return 0
		}
		return v
	})
	go util.AwaitClosure()
	for {
		for z := 0; z < distance3.GetDepth(); z++ {
			distance3.Draw(z)
			// time.Sleep(10 * time.Millisecond)
			if !util.WindowBeingUsed {
				return
			}
		}
	}
}
