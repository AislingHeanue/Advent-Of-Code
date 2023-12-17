package day16

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "16b",
		Short: "Day 16, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

var DrawDistances util.ThreeMatrix[int]

func partB(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	tiles := challenge.TileMap()
	DrawDistances = util.NewThreeMatrix[int](2*(tiles.GetHeight()+tiles.GetWidth()), tiles.GetHeight(), tiles.GetWidth())
	z := 0
	maxTotal := 0
	for x := 0; x < tiles.GetWidth(); x++ {
		distances := util.MapToUnordered[rune, DistanceEntry](tiles, func(y, x int, value rune) DistanceEntry {
			return DistanceEntry{0, false, false, false, false}
		})
		direction := Down
		currentTile := util.Point2D{Y: -1, X: x} //start outside the matrix
		move(currentTile, direction, 0, tiles, &distances)
		processDistances(distances, tiles, z, &maxTotal)
		z++
	}
	for y := 0; y < tiles.GetHeight(); y++ {
		distances := util.MapToUnordered[rune, DistanceEntry](tiles, func(y, x int, value rune) DistanceEntry {
			return DistanceEntry{0, false, false, false, false}
		})
		direction := Left
		currentTile := util.Point2D{Y: y, X: tiles.GetWidth()} //start outside the matrix
		move(currentTile, direction, 0, tiles, &distances)
		processDistances(distances, tiles, z, &maxTotal)
		z++
	}
	for x := tiles.GetWidth() - 1; x >= 0; x-- {
		distances := util.MapToUnordered[rune, DistanceEntry](tiles, func(y, x int, value rune) DistanceEntry {
			return DistanceEntry{0, false, false, false, false}
		})
		direction := Up
		currentTile := util.Point2D{Y: tiles.GetHeight(), X: x} //start outside the matrix
		move(currentTile, direction, 0, tiles, &distances)
		processDistances(distances, tiles, z, &maxTotal)
		z++
	}
	for y := tiles.GetHeight() - 1; y >= 0; y-- {
		distances := util.MapToUnordered[rune, DistanceEntry](tiles, func(y, x int, value rune) DistanceEntry {
			return DistanceEntry{0, false, false, false, false}
		})
		direction := Down
		currentTile := util.Point2D{Y: y, X: -1} //start outside the matrix
		move(currentTile, direction, 0, tiles, &distances)
		processDistances(distances, tiles, z, &maxTotal)
		z++
	}

	return maxTotal
}

func processDistances(distances util.UnorderedMatrix[DistanceEntry], tiles util.Matrix[rune], z int, maxTotal *int) {
	total := 0
	for _, value := range distances.Iterator() {
		if value.distance != 0 {
			total += 1
		}
	}
	if total > *maxTotal {
		*maxTotal = total
	}
	bigNumber := 100000
	drawMatrix := util.UnorderedMapToOrdered[DistanceEntry, int](distances, func(y, x int, value DistanceEntry) int {
		switch tiles.MustGet(y, x) {
		case '.':
			return value.distance
		case '\\':
			return bigNumber
		case '/':
			return bigNumber + 1
		case '-':
			return bigNumber + 2
		case '|':
			return bigNumber + 3
		default:
			panic("illegal character")
		}
	})
	for x1 := 0; x1 < distances.GetWidth(); x1++ {
		for y1 := 0; y1 < distances.GetHeight(); y1++ {
			DrawDistances.Set(z, y1, x1, drawMatrix.MustGet(y1, x1))
		}
	}

}
