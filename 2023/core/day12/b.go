package day12

import (
	"fmt"
	"strings"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
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
	// fmt.Printf("%v %v\n", getWidths(parts[0]), strings.Split(parts[1], ","))
	runningTotal := 0 // I hope this is threadsafe
	line2 := strings.Split(parts[1]+","+parts[1]+","+parts[1]+","+parts[1]+","+parts[1], ",")
	dotOrHashB(parts[0]+parts[0]+parts[0]+parts[0]+parts[0], line2, 0, &runningTotal)
	// fmt.Println(line, runningTotal)line2+line2+line2
	return runningTotal
}

func dotOrHashB(line1 string, line2 []string, n int, runningTotal *int) {}
