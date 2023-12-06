package day6

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "6b",
		Short: "Day 6, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

func partB(challenge *core.Input) int {
	lines := challenge.LineSlice()
	re := regexp.MustCompile(`\d+`)
	times := re.FindAllString(lines[0], -1)
	distances := re.FindAllString(lines[1], -1)
	total := 1
	timeString := ""
	distanceString := ""
	for i := range times {
		timeString += times[i]
		distanceString += distances[i]
	}

	a := -1
	b, _ := strconv.Atoi(timeString)
	minusC, _ := strconv.Atoi(distanceString)
	roots := quadraticRoots(a, b, -minusC)
	possible := int(roots[0]) - int(roots[1])
	if roots[0]-float64(int(roots[0])) == 0 && roots[1]-float64(int(roots[1])) == 0 {
		possible--
	}
	total *= possible

	return total

}
