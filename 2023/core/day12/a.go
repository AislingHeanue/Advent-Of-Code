package day12

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "12a",
		Short: "Day 12, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

func partA(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	res := core.InputMap(challenge, countPossible)
	total := 0
	for _, num := range res {
		total += num
	}
	return total
}

func countPossible(line string) int {
	parts := strings.Split(line, " ")
	// fmt.Printf("%v %v\n", getWidths(parts[0]), strings.Split(parts[1], ","))
	runningTotal := 0 // I hope this is threadsafe
	dotOrHash(parts[0], strings.Split(parts[1], ","), 0, &runningTotal)
	// fmt.Println(line, runningTotal)
	return runningTotal
}

func dotOrHash(line1 string, line2 []string, n int, runningTotal *int) {
	// TODO add caching probably

	if n == len(line1) {
		widths := getWidths(line1)
		if len(line2) != len(widths) {
			return
		}
		// fmt.Println(widths, line2)
		for i := range line2 {
			if line2[i] != widths[i] {
				// fmt.Println("no")
				return
			}
		}
		// fmt.Printf("yes, %v\n", line1)
		*runningTotal++
		return
	}
	if line1[n] != '?' {
		dotOrHash(line1, line2, n+1, runningTotal)
		return
	}
	lineBytes := []byte(line1)
	lineBytes[n] = byte('.')
	dotOrHash(string(lineBytes), line2, n+1, runningTotal)
	lineBytes[n] = byte('#')
	dotOrHash(string(lineBytes), line2, n+1, runningTotal)
}

var widthRe = regexp.MustCompile(`#+`)

func getWidths(line string) []string {
	widthRes := widthRe.FindAllString(line, -1)
	res := []string{}
	for i := 0; i < len(widthRes); i++ {
		res = append(res, fmt.Sprint(len(widthRes[i])))
	}

	return res
}
