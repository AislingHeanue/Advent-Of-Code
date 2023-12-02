package day2

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "2b",
		Short: "Day 2, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

func partB(challenge *core.Input) int {
	total := 0
	re := regexp.MustCompile(`(\d+) ([a-z]+)(?:, (\d+) ([a-z]+))?(?:, (\d+) ([a-z]+))?(?:;|\z)`)
	for _, line := range challenge.LineSlice() {
		counts := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		turns := re.FindAllStringSubmatch(line, -1)
		for _, turn := range turns {
			// format of `turn`: {(full match), digit 1, colour 1, digit 2, colour 2, digit 3, colour 3}
			// so consider pairs (1,2), (3,4) and (5,6)
			for i := 1; i+1 < 7; i += 2 {
				if turn[i] != "" {
					num, err := strconv.Atoi(turn[i])
					if err != nil {
						panic(err)
					}
					counts[turn[i+1]] = max(counts[turn[i+1]], num)
				}
			}
		}
		total += counts["red"] * counts["green"] * counts["blue"]
	}

	return total
}
