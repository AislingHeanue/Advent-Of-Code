package day4

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "4b",
		Short: "Day 4, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

func partB(challenge *core.Input) int {
	var elfDigits int
	lines := challenge.LineSlice()
	scoreCache = make(map[int]int)
	re := regexp.MustCompile(`(\d+)`)
	total := 0
	for i := 0; i < len(lines); i++ {
		total += 1 + scoreLine(lines, i, re, elfDigits)
	}

	return total
}

var scoreCache map[int]int

func scoreLine(lines []string, index int, re *regexp.Regexp, elfDigits int) int {

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
	for k := index + 1; k <= matches+index; k++ {
		total += 1 + scoreLine(lines, k, re, elfDigits)
	}
	scoreCache[index] = total
	return total
}
