package day4

import (
	"fmt"
	"regexp"

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
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	// defer util.AwaitClosure()
	var re *regexp.Regexp
	var elfDigits int
	lines := challenge.LineSlice()
	scoreCache = make(map[int]int)
	switch len(lines[0]) {
	case 116:
		// regex format: 0: whole match, 2-11: elf digits, 12-36: winning digits
		re = regexp.MustCompile(`Card\s+(\d+):\s+` +
			`(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)\|\s+` +
			`(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)` +
			`(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)` +
			`(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(\d+)?`)
		elfDigits = 10
	case 48:
		// regex format: 0: whole match, 2-6: elf digits, 7-14: winning digits
		re = regexp.MustCompile(`Card\s+(\d+):\s+` +
			`(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)\|\s+` +
			`(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(\d+)`)
		elfDigits = 5
	default:
		panic("wrong input line size")
	}
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

	regexRes := re.FindStringSubmatch(lines[index])
	matches := 0
	for j := elfDigits + 2; j < len(regexRes); j++ {
		for i := 2; i < 2+elfDigits; i++ {
			if regexRes[i] == regexRes[j] {
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
