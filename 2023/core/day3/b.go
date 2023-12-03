package day3

import (
	"fmt"
	"unicode"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "3b",
		Short: "Day 3, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

func partB(challenge *core.Input) int {
	mat := challenge.TileMap()
	markedMat := util.Map[rune, int](mat, func(y, x int, v rune) int {
		return 0
	})
	numberMat := util.Map[rune, int](mat, func(y, x int, v rune) int {
		if unicode.IsDigit(v) {
			return int(v - '0')
		}
		return -1
	})
	symbolMat := util.Map[rune, string](mat, func(y, x int, v rune) string {
		if unicode.IsDigit(v) {
			return "."
		}
		return string(v)
	})
	total := 0

	for y := 0; y < numberMat.GetHeight(); y++ {
		for x := 0; x < numberMat.GetWidth(); x++ {
			if symbol, marked := symbolMat.MustGet(y, x), markedMat.MustGet(y, x); symbol != "." && marked == 0 {
				markedMat := util.Map[rune, int](mat, func(y, x int, v rune) int {
					return 0
				})
				values := []int{}
				// check for numbers in surrounding area
				for y2 := y - 1; y2 <= y+1; y2++ {
					for x2 := x - 1; x2 <= x+1; x2++ {
						number, ok := numberMat.Get(y2, x2)
						if marked, alsoOk := markedMat.Get(y2, x2); number != -1 && ok && marked == 0 && alsoOk { // this line is terrible
							left := x2
							right := x2
							value := 0
							// grow number to the left
							for left > 0 && numberMat.MustGet(y2, left-1) != -1 {
								left--
							}
							// grow number to the right
							for right < numberMat.GetWidth()-1 && numberMat.MustGet(y2, right+1) != -1 {
								right++
							}
							// scan the number and mark it as seen
							for x3 := left; x3 <= right; x3++ {
								value *= 10
								value += numberMat.MustGet(y2, x3)
								markedMat.MustSet(y2, x3, 1)
							}
							values = append(values, value)
						}
					}
				}
				if len(values) == 2 {
					total += values[0] * values[1]
				}
			}
		}

	}
	return total
}
