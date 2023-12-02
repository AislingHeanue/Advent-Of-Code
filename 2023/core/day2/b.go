package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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
	re, err := regexp.Compile(` (\d+) (red|green|blue)`)
	if err != nil {
		panic(err)
	}
	for _, line := range challenge.LineSlice() {
		counts := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		parts := strings.Split(line, ":")
		if err != nil {
			panic(err)
		}
		turns := strings.Split(parts[1], ";")
		for _, turn := range turns {
			items := re.FindAllStringSubmatch(turn, -1)
			for _, item := range items {
				num, err := strconv.Atoi(item[1])
				if err != nil {
					panic(err)
				}
				counts[item[2]] = max(counts[item[2]], num)
			}
		}
		total += counts["red"] * counts["green"] * counts["blue"]
	}

	return total
}
