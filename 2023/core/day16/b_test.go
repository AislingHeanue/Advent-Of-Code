package day16

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
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

	result := partB(input)

	require.Equal(t, 51, result)
}

func TestBAnimated(t *testing.T) {
	t.Skipf("ignore animation tests")
	// util.ForceNoWindow = true
	util.EbitenSetup()
	input := core.FromFile()

	_ = partB(input)
	go util.AwaitClosure()
	AnimateB(DrawDistances)

	// require.Equal(t, 7030, result)
}

func AnimateB(drawMatrix util.ThreeMatrix[int]) {
	for {
		for z := 0; z < drawMatrix.GetDepth(); z++ {
			// time.Sleep(50 * time.Millisecond)
			drawMatrix.Draw(z)
			// box.PrintEvenlySpaced(",")
			if !util.WindowBeingUsed {
				return
			}
		}
	}
}
