package day23

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "23a",
		Short: "Day 23, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

func partA(challenge *core.Input) int {
	return solve(challenge, false)
}

var leavesFound int = 0

func solve(challenge *core.Input, b bool) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	tiles := challenge.TileMap()
	startTile := util.Point2D{}
	endTile := util.Point2D{}
	for x := 0; x < tiles.GetWidth(); x++ {
		if tiles.MustGet(0, x) == '.' {
			startTile = util.Point2D{Y: 0, X: x}
		}
		if tiles.MustGet(tiles.GetHeight()-1, x) == '.' {
			endTile = util.Point2D{Y: tiles.GetHeight() - 1, X: x}
		}
	}
	maxValue := 0
	findNextTile(startTile, endTile, make(map[util.Point2D]bool), tiles, &maxValue, b)

	return maxValue
}

func findNextTile(currentTile util.Point2D, endTile util.Point2D, checkedTiles map[util.Point2D]bool, tiles util.Matrix[rune], maxValue *int, b bool) {
	checkedTiles[currentTile] = true
	// I'm not sure if this map will be copied by ref or by value so just in case I'll handle both cases
	// and come back to it later when I'm more brave
	defer func() { checkedTiles[currentTile] = false }()
	if currentTile == endTile {
		leavesFound++
		if steps := numTrue(checkedTiles); steps > *maxValue {
			if b {
				fmt.Printf("%d found after %d leaves checked\n", numTrue(checkedTiles), leavesFound)
			}
			*maxValue = steps
			return
		}
	}
	tileValue := tiles.MustGet(currentTile.Y, currentTile.X)
	// move down
	if tileValue == '.' || tileValue == 'v' || b {
		nextTile := util.Point2D{Y: currentTile.Y + 1, X: currentTile.X}
		if nextValue, ok := tiles.Get(nextTile.Y, nextTile.X); ok && nextValue != '#' && !checkedTiles[nextTile] {
			findNextTile(nextTile, endTile, checkedTiles, tiles, maxValue, b)
		}
	}
	// move up
	if tileValue == '.' || tileValue == '^' || b {
		nextTile := util.Point2D{Y: currentTile.Y - 1, X: currentTile.X}
		if nextValue, ok := tiles.Get(nextTile.Y, nextTile.X); ok && nextValue != '#' && !checkedTiles[nextTile] {
			findNextTile(nextTile, endTile, checkedTiles, tiles, maxValue, b)
		}
	}
	// move left
	if tileValue == '.' || tileValue == '<' || b {
		nextTile := util.Point2D{Y: currentTile.Y, X: currentTile.X - 1}
		if nextValue, ok := tiles.Get(nextTile.Y, nextTile.X); ok && nextValue != '#' && !checkedTiles[nextTile] {
			findNextTile(nextTile, endTile, checkedTiles, tiles, maxValue, b)
		}
	}
	// move right
	if tileValue == '.' || tileValue == '>' || b {
		nextTile := util.Point2D{Y: currentTile.Y, X: currentTile.X + 1}
		if nextValue, ok := tiles.Get(nextTile.Y, nextTile.X); ok && nextValue != '#' && !checkedTiles[nextTile] {
			findNextTile(nextTile, endTile, checkedTiles, tiles, maxValue, b)
		}
	}

}

func numTrue[V comparable](a map[V]bool) int {
	total := 0
	for _, value := range a {
		if value {
			total += 1
		}
	}

	return total - 1 // don't count the start tile
}
