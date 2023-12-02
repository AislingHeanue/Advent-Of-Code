package day2

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "2a",
		Short: "Day 2, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

func partA(challenge *core.Input) int {
	maxMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	total := 0
	re := regexp.MustCompile(`(\d+) ([a-z]+)(?:, (\d+) ([a-z]+))?(?:, (\d+) ([a-z]+))?(?:;|\z)`)
	idRe := regexp.MustCompile(`Game (\d+):`)
	for _, line := range challenge.LineSlice() {
		counts := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		possible := true
		id, err := strconv.Atoi(idRe.FindStringSubmatch(line)[1])
		if err != nil {
			panic(err)
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
		for colour, value := range maxMap {
			if counts[colour] > value {
				possible = false
			}
		}
		if possible {
			total += id
		}
	}

	return total
}
