package day10

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "10b",
		Short: "Day 10, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

// how to check if an existing contour path has encircled a point
// check these places for transitions and direction, and count the total clockwises - anticlockwises / 4 for the cycle count.
// It will either be 1,0 or -1, since the contour is not self intersecting
// ....--.......
// ....--.......
// ....--||||||
// |||||x||||||
// |||||--.....
// .....--.....
// .....--.....

// that was my old solution, which was basically 4 instances of the nonzero rule
// this time, lets apply the even-odd rule instead. This can be done by scanning
// all of the points above a given point, and if an odd number of them match conditions
// 1. they are in the contour
// 2. they have a path to the left
// then we are inside the contour

func partB(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	util.EbitenSetup()
	letters := challenge.TileMap()
	tiles := util.MapToUnordered(letters, toTile)
	sLocation := util.Point2D{}
	for point, value := range tiles.Iterator() {
		if value.isStart {
			sLocation = point
			pointsToCheck := map[string]util.Point2D{
				"up":    {Y: point.Y - 1, X: point.X},
				"left":  {Y: point.Y, X: point.X - 1},
				"right": {Y: point.Y, X: point.X + 1},
				"down":  {Y: point.Y + 1, X: point.X},
			}
			for direction, point2 := range pointsToCheck {
				value2, _ := tiles.Get(point2.Y, point2.X)
				switch direction {
				case "up":
					value.up = value2.down
				case "down":
					value.down = value2.up
				case "left":
					value.left = value2.right
				case "right":
					value.right = value2.left
				}
			}
			tiles.MustSet(point.Y, point.X, value)
			break
		}
	}

	distances := util.UnorderedMapToOrdered[Tile, int](tiles, func(y, x int, v Tile) int {
		return 0
	})
	_ = march(sLocation, tiles, &distances)
	total := 0
	for point, value := range distances.Iterator() {
		if value != 0 {
			continue
		}
		leftCount := 0
		for y := 0; y < point.Y; y++ {
			if tiles.MustGet(y, point.X).left && distances.MustGet(y, point.X) != 0 {
				leftCount++
			}
		}
		total += leftCount % 2
	}
	return total
}
