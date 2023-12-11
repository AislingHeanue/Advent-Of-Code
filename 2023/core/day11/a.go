package day11

import (
	"fmt"
	"math"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "11a",
		Short: "Day 11, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

func partA(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	tiles := challenge.TileMap()
	rowNotEmpty := make(map[int]int)
	columnNotEmpty := make(map[int]int)
	galaxies := []util.Point2D{}
	for point, value := range tiles.Iterator() {
		if value == '#' {
			rowNotEmpty[point.Y] = 1
			columnNotEmpty[point.X] = 1
			galaxies = append(galaxies, point)
		}
	}
	total := 0
	for i, g1 := range galaxies {
		for j, g2 := range galaxies {
			if i <= j {
				continue
			}
			total += taxicabMetric(g1, g2, rowNotEmpty, columnNotEmpty)
			// fmt.Printf("p1: (%d,%d), p2: (%d,%d), metric: %d\n", g1.Y, g1.X, g2.Y, g2.X, taxicabMetric(g1, g2, rowNotEmpty, columnNotEmpty))
		}
	}

	return total
}

func taxicabMetric(p1, p2 util.Point2D, rowNotEmpty, columnNotEmpty map[int]int) int {
	total := int(math.Abs(float64(p1.X-p2.X))) + int(math.Abs(float64(p1.Y-p2.Y)))
	for x := min(p1.X, p2.X); x < max(p1.X, p2.X); x++ {
		if _, ok := columnNotEmpty[x]; !ok {
			total += 1
		}
	}
	for y := min(p1.Y, p2.Y); y < max(p1.Y, p2.Y); y++ {
		if _, ok := rowNotEmpty[y]; !ok {
			total += 1
		}
	}
	return total
}
