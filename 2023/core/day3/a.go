package day3

import (
	"fmt"
	"unicode"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "3a",
		Short: "Day 3, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

type ElementA struct {
	symbol string
	number int
	marked bool
}

func partA(challenge *core.Input) int {
	mat := challenge.TileMap()
	matrix := util.MapToUnordered[rune, ElementA](mat, func(y, x int, v rune) ElementA {
		if unicode.IsDigit(v) {
			return ElementA{
				number: int(v - '0'),
				symbol: ".",
				marked: false,
			}
		} else {
			return ElementA{
				number: -1,
				symbol: string(v),
				marked: false,
			}
		}
	})
	total := 0

	for y := 0; y < matrix.GetHeight(); y++ {
		for x := 0; x < matrix.GetWidth(); x++ {
			if elem := matrix.MustGet(y, x); elem.symbol != "." {
				// check for numbers in surrounding area
				for y2 := y - 1; y2 <= y+1; y2++ {
					for x2 := x - 1; x2 <= x+1; x2++ {
						newElem, ok := matrix.Get(y2, x2)
						if ok && newElem.number != -1 && !newElem.marked {
							left := x2
							right := x2
							value := 0
							// grow number to the left
							for left > 0 && matrix.MustGet(y2, left-1).number != -1 {
								left--
							}
							// grow number to the right
							for right < matrix.GetWidth()-1 && matrix.MustGet(y2, right+1).number != -1 {
								right++
							}
							// scan the number and mark it as seen
							for x3 := left; x3 <= right; x3++ {
								elem3 := matrix.MustGet(y2, x3)
								value *= 10
								value += elem3.number
								elem3.marked = true
								matrix.MustSet(y2, x3, elem3)
							}
							total += value
						}
					}
				}
			}
		}
	}
	return total
}
