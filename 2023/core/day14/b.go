package day14

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "14b",
		Short: "Day 14, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

func partB(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	mat := challenge.TileMap()
	boulderMemory := make(map[string]bool)
	for i := 0; i < 10000000; i++ {
		cycle(&mat, &boulderMemory)
	}
	strMat := util.Map[rune, string](mat, func(y, x int, v rune) string {
		return string(v)
	})
	strMat.Print("")
	return getScoreB(mat)
}

func cycle(mat *util.Matrix[rune], memory *map[string]bool) {
	boulderString := getBoulders(*mat)
	if _, ok := (*memory)[boulderString]; ok {
		fmt.Println("I found a cycle")
	}
	(*memory)[boulderString] = true

	for x := 0; x < mat.GetWidth(); x++ {
		nextBoulderY := 0
		for y := 0; y < mat.GetHeight(); y++ {
			if mat.MustGet(y, x) == '#' {
				nextBoulderY = y + 1
			}
			if mat.MustGet(y, x) == 'O' {
				if nextBoulderY != y {
					mat.MustSet(y, x, '.')
					mat.MustSet(nextBoulderY, x, 'O')
				}
				nextBoulderY++
			}
		}
	}
	for y := 0; y < mat.GetHeight(); y++ {
		nextBoulderX := 0
		for x := 0; x < mat.GetWidth(); x++ {
			if mat.MustGet(y, x) == '#' {
				nextBoulderX = x + 1
			}
			if mat.MustGet(y, x) == 'O' {
				if nextBoulderX != x {
					mat.MustSet(y, x, '.')
					mat.MustSet(y, nextBoulderX, 'O')
				}
				nextBoulderX++
			}
		}
	}
	for x := 0; x < mat.GetWidth(); x++ {
		nextBoulderY := mat.GetHeight() - 1
		for y := mat.GetHeight() - 1; y >= 0; y-- {
			if mat.MustGet(y, x) == '#' {
				nextBoulderY = y - 1
			}
			if mat.MustGet(y, x) == 'O' {
				if nextBoulderY != y {
					mat.MustSet(y, x, '.')
					mat.MustSet(nextBoulderY, x, 'O')
				}
				nextBoulderY--
			}
		}
	}
	for y := 0; y < mat.GetHeight(); y++ {
		nextBoulderX := mat.GetWidth() - 1
		for x := mat.GetWidth() - 1; x >= 0; x-- {
			if mat.MustGet(y, x) == '#' {
				nextBoulderX = x - 1
			}
			if mat.MustGet(y, x) == 'O' {
				if nextBoulderX != x {
					mat.MustSet(y, x, '.')
					mat.MustSet(y, nextBoulderX, 'O')
				}
				nextBoulderX--
			}
		}
	}
}

func getBoulders(mat util.Matrix[rune]) string {
	boulders := []util.Point2D{}
	for y := 0; y < mat.GetHeight(); y++ {
		for x := 0; x < mat.GetWidth(); x++ {
			if mat.MustGet(y, x) == 'O' {
				boulders = append(boulders, util.Point2D{Y: y, X: x})
			}
		}
	}
	return fmt.Sprintf("%v", boulders)
}

func getScoreB(mat util.Matrix[rune]) int {
	total := 0
	for y := 0; y < mat.GetHeight(); y++ {
		for x := 0; x < mat.GetWidth(); x++ {
			if mat.MustGet(y, x) == 'O' {
				total += mat.GetHeight() - y
			}
		}
	}

	return total
}
