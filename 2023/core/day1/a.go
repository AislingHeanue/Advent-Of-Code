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
	return solve(challenge, false)
}

func solve(challenge *core.Input, b bool) int {
	total := 0
	numberList := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, line := range challenge.LineSlice() {
		firstNumber := -1
		lastNumber := -1
		for i, letter := range line {
			if unicode.IsDigit(letter) {
				if firstNumber == -1 {
					firstNumber = int(letter - '0')
				}
				lastNumber = int(letter - '0')
			} else if b {
				for j, num := range numberList {
					if line[i:min(len(line), i+len(num))] == num {
						if firstNumber == -1 {
							firstNumber = j
						}
						lastNumber = j
					}
				}
			}
		}
		total += 10*firstNumber + lastNumber
	}

	return total
}
