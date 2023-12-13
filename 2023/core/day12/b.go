package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "12b",
		Short: "Day 12, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

func partB(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	res := core.InputMap(challenge, countPossibleB)
	total := 0
	for _, num := range res {
		total += num
	}
	return total
}

func countPossibleB(line string) int {
	parts := strings.Split(line, " ")
	line1 := parts[0] + "?" + parts[0] + "?" + parts[0] + "?" + parts[0] + "?" + parts[0]
	line2 := strings.Split(parts[1]+","+parts[1]+","+parts[1]+","+parts[1]+","+parts[1], ",")
	line2Num := make([]int, len(line2))
	for i := range line2 {
		line2Num[i], _ = strconv.Atoi(line2[i])
	}
	possibleMatrix := util.NewMatrix[int](len(line2Num), len(line1))
	possibleMatrix.Fill(-1)
	b := getPossible(len(line1)-1, len(line2Num)-1, &possibleMatrix, line1, line2Num)
	return b
}
