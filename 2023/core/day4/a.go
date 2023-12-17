package day4

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "4a",
		Short: "Day 4, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

func partA(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	// defer util.AwaitClosure()
	total := 0
	re := regexp.MustCompile(`(\d+)`)

	lines := challenge.LineSlice()
	for i := range lines {
		total += scoreLine(lines, i, re, false)
	}

	return total
}

func value(matches int) int {
	if matches == 0 {
		return 0
	}
	return util.Power(2, matches-1)
}

var scoreCache map[int]int

func scoreLine(lines []string, index int, re *regexp.Regexp, b bool) int {
	total, ok := scoreCache[index]
	if ok {
		return total
	}
	lineParts := strings.Split(strings.Split(lines[index], ":")[1], "|")
	leftNums := re.FindAllString(lineParts[0], -1)
	rightNums := re.FindAllString(lineParts[1], -1)
	matches := 0
	for j := 0; j < len(leftNums); j++ {
		for i := 0; i < len(rightNums); i++ {
			if rightNums[i] == leftNums[j] {
				matches++
			}
		}
	}
	if b {
		for k := index + 1; k <= matches+index; k++ {
			total += 1 + scoreLine(lines, k, re, true)
		}
		scoreCache[index] = total
	} else {
		total += value(matches)
	}

	return total
}
