package day6

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "6a",
		Short: "Day 6, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}
func partA(challenge *core.Input) int {
	return solve(challenge, false)
}

func solve(challenge *core.Input, b bool) int {
	lines := challenge.LineSlice()
	re := regexp.MustCompile(`\d+`)
	times := re.FindAllString(lines[0], -1)
	distances := re.FindAllString(lines[1], -1)
	total := 1
	if b {
		timeString := ""
		distanceString := ""
		for i := range times {
			timeString += times[i]
			distanceString += distances[i]
		}
		times = []string{timeString}
		distances = []string{distanceString}
	}
	for i := range times {
		a := -1
		b, _ := strconv.Atoi(times[i])
		minusC, _ := strconv.Atoi(distances[i])
		roots := quadraticRoots(a, b, -minusC)
		possible := int(roots[0]) - int(roots[1])
		if roots[0]-float64(int(roots[0])) == 0 && roots[1]-float64(int(roots[1])) == 0 {
			possible--
		}
		total *= possible
	}

	return total
}

func quadraticRoots(a, b, c int) []float64 {
	return []float64{
		(float64(-b) - math.Sqrt(float64(b*b-4*a*c))) / float64(2*a),
		(float64(-b) + math.Sqrt(float64(b*b-4*a*c))) / float64(2*a),
	}
}
