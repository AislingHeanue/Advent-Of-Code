package day21

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "21a",
		Short: "Day 21, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

func partA(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	tiles := challenge.TileMap()
	steps := 0
	switch tiles.GetWidth() {
	case 11:
		steps = 6
	case 131:
		steps = 64
	default:
		fmt.Println(tiles.GetWidth())
	}
	startTile := util.Point2D{}
	for point, value := range tiles.Iterator() {
		if value == 'S' {
			startTile = point
			break
		}
	}
	distances := getDistances(startTile, tiles)
	total := 0
	for _, value := range distances.Iterator() {
		if value <= steps && (value-steps)%2 == 0 && value != -1 {
			total += 1
		}
	}

	return total
}

func getDistances(startTile util.Point2D, tiles util.Matrix[rune]) util.Matrix[int] {
	distances := util.Map[rune, int](tiles, func(y, x int, value rune) int { return -1 })
	distances.MustSet(startTile.Y, startTile.X, 0)
	queue := []util.Point2D{startTile}
	for len(queue) > 0 {
		currentTile := queue[0]
		currentDistance := distances.MustGet(currentTile.Y, currentTile.X)
		queue = queue[1:]
		newPoints := []util.Point2D{
			{Y: currentTile.Y + 1, X: currentTile.X},
			{Y: currentTile.Y - 1, X: currentTile.X},
			{Y: currentTile.Y, X: currentTile.X + 1},
			{Y: currentTile.Y, X: currentTile.X - 1},
		}
		for _, point := range newPoints {
			newTile, _ := tiles.Get(point.Y, point.X)
			newDistance, ok := distances.Get(point.Y, point.X)
			if ok && newTile != '#' && (newDistance > currentDistance+1 || newDistance == -1) {
				distances.MustSet(point.Y, point.X, currentDistance+1)
				queue = append(queue, point)
			}
		}
	}

	return distances
}
