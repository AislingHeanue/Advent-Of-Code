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
// check there places for transitions and direction, and count the total clockwises - anticlockwises / 4 for the cycle count.
// It will either be 1,0 or -1, since the contour is not self intersecting
// ....--.......
// ....--.......
// ....--||||||
// |||||x||||||
// |||||--.....
// .....--.....
// .....--.....

func partB(challenge *core.Input) int {
	// use ebiten to draw the main loop after it's been found :)
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	util.EbitenSetup()
	defer util.AwaitClosure()
	letters := challenge.TileMap()
	tiles := util.MapToUnordered(letters, toTile)
	sLocation := util.Point2D{}
	for point, value := range tiles.Iterator() {
		if !value.isStart {
			continue
		} else {
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
	d := march(sLocation, tiles, &distances)
	magicalOffset := 2*d - 3
	// fmt.Println(magicalOffset)
	// distances.PrintEvenlySpaced(" ")
	cycleCounts := util.Map[int, int](distances, func(y, x, value int) int {
		if value != 0 {
			return 0
		}
		return 2 // 2 is the stand-in for "as of yet unknown". The final values will be either all -1 and 0 or all 1 and 0
	})
	// cycleCounts.PrintEvenlySpaced(" ")
	total := 0
	for point, value := range cycleCounts.Iterator() {
		clockwiseTurns := 0
		if value != 2 {
			continue
		}
		// UP
		if point.X != 0 && point.Y != 0 { //conditions needed to be able to check the top cycles
			for y := 0; y < point.Y; y++ {
				offset := distances.MustGet(y, point.X) - distances.MustGet(y, point.X-1)
				if offset == 1 || offset == -1 {
					clockwiseTurns += offset
				}
				neitherZero := distances.MustGet(y, point.X) > 0 && distances.MustGet(y, point.X-1) > 0
				if neitherZero && (offset == magicalOffset || offset == -magicalOffset) {
					clockwiseTurns -= offset / magicalOffset
				}
			}
		}
		// DOWN
		if point.X != cycleCounts.GetWidth()-1 && point.Y != cycleCounts.GetHeight()-1 {
			for y := cycleCounts.GetHeight() - 1; y > point.Y; y-- {
				offset := distances.MustGet(y, point.X) - distances.MustGet(y, point.X+1)
				if offset == 1 || offset == -1 {
					clockwiseTurns += offset
				}
				neitherZero := distances.MustGet(y, point.X) > 0 && distances.MustGet(y, point.X+1) > 0
				if neitherZero && (offset == magicalOffset || offset == -magicalOffset) {
					clockwiseTurns -= offset / magicalOffset
				}
			}
		}
		// LEFT
		if point.X != 0 && point.Y != cycleCounts.GetHeight()-1 {
			for x := 0; x < point.X; x++ {
				offset := distances.MustGet(point.Y, x) - distances.MustGet(point.Y+1, x)
				if offset == 1 || offset == -1 {
					clockwiseTurns += offset
				}
				neitherZero := distances.MustGet(point.Y, x) > 0 && distances.MustGet(point.Y+1, x) > 0
				if neitherZero && (offset == magicalOffset || offset == -magicalOffset) {
					clockwiseTurns -= offset / magicalOffset
				}
			}
		}
		// RIGHT
		if point.X != cycleCounts.GetWidth()-1 && point.Y != 0 {
			for x := cycleCounts.GetWidth() - 1; x > point.X; x-- {
				offset := distances.MustGet(point.Y, x) - distances.MustGet(point.Y-1, x)
				if offset == 1 || offset == -1 {
					clockwiseTurns += offset
				}
				neitherZero := distances.MustGet(point.Y, x) > 0 && distances.MustGet(point.Y-1, x) > 0
				if neitherZero && (offset == magicalOffset || offset == -magicalOffset) {
					clockwiseTurns -= offset / magicalOffset
				}
			}
		}
		cycleCounts.MustSet(point.Y, point.X, clockwiseTurns)
		if clockwiseTurns != 0 {
			total += 1
		}
	}
	// cycleCounts.PrintEvenlySpaced(" ")
	return total
}
