package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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
		possible := true
		parts := strings.Split(line, ":")
		id, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
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
