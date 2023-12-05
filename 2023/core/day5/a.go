package day5

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "5a",
		Short: "Day 5, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

func partA(challenge *core.Input) int {
	lines := challenge.LineSlice()
	re := regexp.MustCompile(`\d+`)
	maps := make([][][]int, 7)
	i := 0
	mapsIndex := -1
	seeds := re.FindAllString(lines[i], -1)
	i++
	for i < len(lines) {
		strings := re.FindAllString(lines[i], 3)
		if len(strings) != 3 {
			i += 2
			mapsIndex++
			maps[mapsIndex] = [][]int{}
		} else {
			numbers := make([]int, 3)
			for i, str := range strings {
				numbers[i], _ = strconv.Atoi(str)
			}
			maps[mapsIndex] = append(maps[mapsIndex], numbers)
			i++
		}
	}
	lowestLoc := -1

	for _, seed := range seeds {
		number, _ := strconv.Atoi(seed)
		for j := 0; j < 7; j++ {
			number = findValue(maps[j], number)
		}
		if lowestLoc > number || lowestLoc == -1 {
			lowestLoc = number
		}
	}

	return lowestLoc
}

func findValue(mat [][]int, num int) int {
	for _, list := range mat {
		offset := num - list[1]
		if offset >= 0 && offset < list[2] {
			return offset + list[0]
		}
	}

	return num
}
