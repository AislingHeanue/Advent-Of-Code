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

type MarkInt struct {
	Symbol string
	Value  int
	Marked bool
}

func (m MarkInt) String() string {
	if m.Symbol == "" {
		return fmt.Sprint(m.Value)
	}

	return m.Symbol
}

func partA(challenge *core.Input) int {
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
			if v, marked := numberMat.MustGet(y, x), markedMat.MustGet(y, x); v != -1 && marked == 0 {
				value := 0
				length := 0
				for i := 0; v != -1; i++ {
					v, ok := numberMat.Get(y, x+i)
					if !ok || v == -1 {
						break
					}
					markedMat.MustSet(y, x+i, 1)
					value *= 10
					value += v
					length += 1
				}
				// check for symbols in surrounding area
				for y2 := y - 1; y2 <= y+1; y2++ {
					for x2 := x - 1; x2 < x+length+1; x2++ {
						symbol, ok := symbolMat.Get(y2, x2)
						if symbol != "." && ok {
							total += value
						}
					}
				}
			}
		}

	}
	return total
}
