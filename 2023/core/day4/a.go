package day4

import (
	"fmt"
	"regexp"

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

	res := core.InputMap(challenge, solveLineA)
	for _, num := range res {
		total += num
	}

	return total
}

var (
	re116 = regexp.MustCompile(`Card\s+(\d+):\s+` +
		`(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)\|\s+` +
		`(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)` +
		`(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)` +
		`(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(\d+)?`)
	re48 = regexp.MustCompile(`Card\s+(\d+):\s+` +
		`(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)\|\s+` +
		`(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(?:(\d+)\s+)(\d+)`)
)

func solveLineA(line string) int {
	var re *regexp.Regexp
	var elfDigits int
	switch len(line) {
	case 116:
		// regex format: 0: whole match, 2-11: elf digits, 12-36: winning digits
		re = re116
		elfDigits = 10
	case 48:
		// regex format: 0: whole match, 2-6: elf digits, 7-14: winning digits
		re = re48
		elfDigits = 5
	default:
		panic("wrong input line size")
	}
	regexRes := re.FindStringSubmatch(line)
	matches := 0
	for j := elfDigits + 2; j < len(regexRes); j++ {
		for i := 2; i < 2+elfDigits; i++ {
			if regexRes[i] == regexRes[j] {
				matches++
			}
		}
	}

	return value(matches)
}

func value(matches int) int {
	if matches == 0 {
		return 0
	}
	return util.Power(2, matches-1)
}
