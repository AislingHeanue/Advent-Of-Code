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

type Coordinate struct {
	y int
	x int
}

type Element struct {
	number               int
	symbol               string
	markedBySymbolCoords *map[Coordinate]bool
}

func partB(challenge *core.Input) int {
	mat := challenge.TileMap()
	matrix := util.MapToUnordered[rune, Element](mat, func(y, x int, v rune) Element {
		if unicode.IsDigit(v) {
			newMap := make(map[Coordinate]bool)
			return Element{
				number:               int(v - '0'),
				symbol:               ".",
				markedBySymbolCoords: &newMap,
			}
		} else {
			return Element{
				number:               -1,
				symbol:               string(v),
				markedBySymbolCoords: nil,
			}
		}
	})
	total := 0

	for y := 0; y < matrix.GetHeight(); y++ {
		for x := 0; x < mat.GetWidth(); x++ {
			if elem := matrix.MustGet(y, x); elem.symbol != "." {
				values := []int{}
				// check for numbers in surrounding area
				for y2 := y - 1; y2 <= y+1; y2++ {
					for x2 := x - 1; x2 <= x+1; x2++ {
						newElem, ok := matrix.Get(y2, x2)
						if ok && newElem.number != -1 && !(*newElem.markedBySymbolCoords)[Coordinate{y: y, x: x}] {
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
								(*elem3.markedBySymbolCoords)[Coordinate{y: y, x: x}] = true
								matrix.MustSet(y2, x3, elem3)
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
