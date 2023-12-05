package day5

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "5b",
		Short: "Day 5, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

var checkMap map[int]int

func partB(challenge *core.Input) int {
	lines := challenge.LineSlice()
	re := regexp.MustCompile(`\d+`)
	maps := make([][][]int, 7)
	i := 0
	mapsIndex := -1
	checkMap = make(map[int]int)
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

	for seedIndex := 0; seedIndex < len(seeds); seedIndex += 2 {
		seedsNum1, _ := strconv.Atoi(seeds[seedIndex])
		seedsNum2, _ := strconv.Atoi(seeds[seedIndex+1])
		ok := true

		for seed := seedsNum1; seed < seedsNum1+seedsNum2+1; seed++ {
			newSeed := seed
			for j := 0; j < 7; j++ {
				newSeed, ok = findValueB(maps[j], newSeed, j)
				if !ok {
					newSeed = -1
					break
				}
			}
			if (lowestLoc > newSeed || lowestLoc == -1) && newSeed != -1 {
				lowestLoc = newSeed
				fmt.Printf("New lowest %d found with seed %d. Seed ranges checked: %d/%d\n", lowestLoc, seed, seedIndex/2, len(seeds))
			}
		}
	}

	return lowestLoc
}

func findValueB(mat [][]int, num int, step int) (int, bool) {
	previousLowestJ, ok := checkMap[num]
	if ok && step >= previousLowestJ {
		return 0, false
	}
	for _, list := range mat {
		offset := num - list[1]
		if offset >= 0 && offset < list[2] {
			return offset + list[0], true
		}
	}

	return num, true
}
