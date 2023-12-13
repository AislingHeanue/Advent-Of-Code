package day13

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "13b",
		Short: "Day 13, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

func partB(challenge *core.Input) int {
	lines := challenge.LineSlice()
	i := 0
	total := 0
	for i < len(lines) {
		image := [][]rune{}
		for i < len(lines) && lines[i] != "" {
			image = append(image, []rune(lines[i]))
			i++
		}
		total += lookForMirrorB(image)
		i++
	}

	return total
}

func lookForMirrorB(matrix [][]rune) int {
	height := len(matrix)
	width := len(matrix[0])
	x, found := lookForVerticalB(matrix, height, width)
	if found {
		// need number of the row (1-indexed) to the left of the vertical line
		return x
	}
	y, found := lookForHorizontalB(matrix, height, width)
	if found {
		// need number of the column (1-indexed) above the vertical line
		return 100 * y
	}

	panic("no reflection found!")
}

// checks pairs in the following order:
// 321123..
// 654456..
// 987789..
func lookForVerticalB(matrix [][]rune, height int, width int) (int, bool) {
	for xLeft := 0; xLeft < width-1; xLeft++ {
		mirrorValid := true
		smudges := 0
		for y := 0; y < height; y++ {
			for offset := 0; offset < min(xLeft+1, width-xLeft-1); offset++ {
				if matrix[y][xLeft-offset] != matrix[y][xLeft+1+offset] {
					smudges++
					if smudges > 1 {
						mirrorValid = false
						break
					}
				}
			}
			if !mirrorValid {
				break
			}
		}
		if mirrorValid && smudges == 1 {
			return xLeft + 1, true
		}
	}

	return 0, false
}

func lookForHorizontalB(matrix [][]rune, height int, width int) (int, bool) {
	for yTop := 0; yTop < height-1; yTop++ {
		mirrorValid := true
		smudges := 0
		for x := 0; x < width; x++ {
			for offset := 0; offset < min(yTop+1, height-yTop-1); offset++ {
				if matrix[yTop-offset][x] != matrix[yTop+1+offset][x] {
					smudges++
					if smudges > 1 {
						mirrorValid = false
						break
					}
				}
			}
			if !mirrorValid {
				break
			}
		}
		if mirrorValid && smudges == 1 {
			return yTop + 1, true
		}
	}

	return 0, false
}
