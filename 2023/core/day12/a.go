package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "12a",
		Short: "Day 12, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

func partA(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	res := core.InputMap(challenge, countPossible)
	total := 0
	for _, num := range res {
		total += num
	}
	return total
}

func countPossible(line string) int {
	parts := strings.Split(line, " ")
	line1 := parts[0]
	line2 := strings.Split(parts[1], ",")
	line2Num := make([]int, len(line2))
	for i := range line2 {
		line2Num[i], _ = strconv.Atoi(line2[i])
	}
	possibleMatrix := util.NewMatrix[int](len(line2Num), len(line1))
	possibleMatrix.Fill(-1)
	b := getPossible(len(line1)-1, len(line2Num)-1, &possibleMatrix, line1, line2Num)
	// fmt.Println(b)
	return b
}

func getPossible(stringIndex int, numsIndex int, possibleMatrix *util.Matrix[int], line1 string, line2Num []int) int {
	total := 0
	// EDGE CASES
	if numsIndex == -1 {
		for i := 0; i <= stringIndex; i++ {
			if line1[i] == '#' {
				return 0 // extra # ruins everything
			}
		}
		return 1
	} else if line2Num[numsIndex]-1 > stringIndex {
		return 0
	}
	if stringIndex < 0 {
		return 0
	}

	// ALREADY CACHED (non-trivial) RESULTS
	if a := possibleMatrix.MustGet(numsIndex, stringIndex); a != -1 {
		return a
	}

	hashPossible := true
	for i := 0; i < line2Num[numsIndex]; i++ {
		if line1[stringIndex-i] == '.' {
			hashPossible = false
		}
	}
	if stringIndex-line2Num[numsIndex] >= 0 && line1[stringIndex-line2Num[numsIndex]] == '#' {
		hashPossible = false // want a . to separate blocks of #
	}

	// the actual recursion
	if hashPossible {
		total += getPossible(stringIndex-line2Num[numsIndex]-1, numsIndex-1, possibleMatrix, line1, line2Num)
	}
	if line1[stringIndex] != '#' {
		total += getPossible(stringIndex-1, numsIndex, possibleMatrix, line1, line2Num)
	}
	possibleMatrix.MustSet(numsIndex, stringIndex, total)

	return total
}
