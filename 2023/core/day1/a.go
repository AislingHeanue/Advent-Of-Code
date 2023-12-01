package day1

import (
	"fmt"
	"unicode"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "1a",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

func partA(challenge *core.Input) int {
	total := 0
	for _, line := range challenge.LineSlice() {
		firstNumber := -1
		lastNumber := -1
		for _, letter := range line {
			if unicode.IsDigit(letter) {
				if firstNumber == -1 {
					firstNumber = int(letter - '0')
				}
				lastNumber = int(letter - '0')
			}
		}
		total += 10*firstNumber + lastNumber
	}

	return total
}
