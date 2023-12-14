package day14

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "14a",
		Short: "Day 14, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

func partA(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	mat := challenge.TileMap()

	return getScore(mat)
}

func getScore(mat util.Matrix[rune]) int {
	h := mat.GetHeight()
	total := 0
	for x := 0; x < mat.GetWidth(); x++ {
		rocksFoundInColumn := 0
		for y := 0; y < mat.GetHeight(); y++ {
			if mat.MustGet(y, x) == '#' {
				rocksFoundInColumn = y + 1
			}
			if mat.MustGet(y, x) == 'O' {
				total += h - rocksFoundInColumn
				rocksFoundInColumn++
			}
		}
	}

	return total
}
