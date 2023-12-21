package day4

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "4b",
		Short: "Day 4, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

func partB(challenge *core.Input) int {
	lines := challenge.LineSlice()
	scoreCache := make(map[int]int)
	total := 0
	for i := 0; i < len(lines); i++ {
		total += 1 + scoreLine(lines, i, true, &scoreCache)
	}

	return total
}
