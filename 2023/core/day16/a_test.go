package day16

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	// Test = true
	input := core.FromLiteral(`.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`)

	result := partA(input)
	// go util.AwaitClosure()
	// Animate(DrawMatrix)

	require.Equal(t, 46, result)
}

func TestAAnimated(t *testing.T) {
	t.Skipf("ignore animation tests")
	// util.ForceNoWindow = true
	Test = true
	input := core.FromFile()

	_ = partA(input)
	go util.AwaitClosure()
	Animate(DrawMatrix)

	// require.Equal(t, 7030, result)
}

func Animate(drawMatrix util.Matrix[int]) {
	box := util.NewThreeMatrix[int](MaxCount/10, drawMatrix.GetHeight(), drawMatrix.GetWidth())
	box.SetByRule(func(z, y, x int) int {
		v := drawMatrix.MustGet(y, x)
		if v < z*10 {
			return v
		} else if v == z*10 && z != 0 {
			return 200000
		} else {
			if v > 100000 {
				return v
			} else {
				return 0
			}
		}
	})

	for {
		for z := 0; z < box.GetDepth(); z++ {
			// time.Sleep(50 * time.Millisecond)
			box.Draw(z)
			// box.PrintEvenlySpaced(",")
			if !util.WindowBeingUsed {
				return
			}
		}
	}
}
