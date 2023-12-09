package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "9b",
		Short: "Day 9, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

func partB(challenge *core.Input) int {
	res := core.InputMap(challenge, getPreviousValue)
	total := 0
	for _, num := range res {
		total += num
	}

	return total
}

func getPreviousValue(line string) int {
	numberStrings := strings.Split(line, " ")
	numbers := make([]int, len(numberStrings))
	for i := range numberStrings {
		numbers[i], _ = strconv.Atoi(numberStrings[i])
	}
	return getPolynomial(numbers)(-1)
}
