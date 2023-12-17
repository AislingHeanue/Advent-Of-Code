package day9

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "9a",
		Short: "Day 9, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

func partA(challenge *core.Input) int {
	res := core.InputMap(challenge, getNextValue)
	total := 0
	for _, num := range res {
		total += num
	}

	return total
}

func getNextValue(line string) int {
	numbers := makeNumbers(line)
	return getPolynomial(numbers)(len(numbers))
}

func makeNumbers(line string) []int {
	numberStrings := strings.Split(line, " ")
	numbers := make([]int, len(numberStrings))
	for i := range numberStrings {
		numbers[i], _ = strconv.Atoi(numberStrings[i])
	}
	return numbers
}

func getPolynomial(ys []int) func(x int) int {
	// https://en.wikipedia.org/wiki/Lagrange_polynomial
	return func(x int) int {
		total := 0.
		for xj, yj := range ys {
			lj := 1.
			for xm := range ys {
				if xm == xj {
					continue
				}
				lj *= float64(x-xm) / float64(xj-xm)
			}
			total += float64(yj) * lj
		}
		// if total-math.Round(total) != 0 {
		// 	fmt.Printf("non-int total: %v is now %d\n", total, int(math.Round(total)))
		// }
		return int(math.Round(total))
	}
}
