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
		moveB(currentTile, direction, 0, tiles, &distances)
		processDistances(distances, tiles, z, &maxTotal)
		z++
	}
	for y := 0; y < tiles.GetHeight(); y++ {
		distances := util.MapToUnordered[rune, DistanceEntry](tiles, func(y, x int, value rune) DistanceEntry {
			return DistanceEntry{0, false, false, false, false}
		})
		direction := Left
		currentTile := util.Point2D{Y: y, X: tiles.GetWidth()} //start outside the matrix
		moveB(currentTile, direction, 0, tiles, &distances)
		processDistances(distances, tiles, z, &maxTotal)
		z++
	}
	for x := tiles.GetWidth() - 1; x >= 0; x-- {
		distances := util.MapToUnordered[rune, DistanceEntry](tiles, func(y, x int, value rune) DistanceEntry {
			return DistanceEntry{0, false, false, false, false}
		})
		direction := Up
		currentTile := util.Point2D{Y: tiles.GetHeight(), X: x} //start outside the matrix
		moveB(currentTile, direction, 0, tiles, &distances)
		processDistances(distances, tiles, z, &maxTotal)
		z++
	}
	for y := tiles.GetHeight() - 1; y >= 0; y-- {
		distances := util.MapToUnordered[rune, DistanceEntry](tiles, func(y, x int, value rune) DistanceEntry {
			return DistanceEntry{0, false, false, false, false}
		})
		direction := Down
		currentTile := util.Point2D{Y: y, X: -1} //start outside the matrix
		moveB(currentTile, direction, 0, tiles, &distances)
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

func moveB(currentTile util.Point2D, direction Direction, count int, tiles util.Matrix[rune], distances *util.UnorderedMatrix[DistanceEntry]) {
	d, ok := distances.Get(currentTile.Y, currentTile.X)

	if count < d.distance || d.distance == 0 {
		d.distance = count //want distances to be overwritten each time (purely because it looks nice)
	}
	if count > MaxCount {
		MaxCount = count
	}

	// move to the next tile
	switch direction {
	case Up:
		if d.hasUp {
			return //we've already been here in this direction, must be a loop
		}
		d.hasUp = true
		if ok {
			distances.MustSet(currentTile.Y, currentTile.X, d)
		}
		currentTile.Y -= 1
	case Down:
		if d.hasDown {
			return
		}
		d.hasDown = true
		if ok {
			distances.MustSet(currentTile.Y, currentTile.X, d)
		}
		currentTile.Y += 1
	case Left:
		if d.hasLeft {
			return
		}
		d.hasLeft = true
		if ok {
			distances.MustSet(currentTile.Y, currentTile.X, d)
		}
		currentTile.X -= 1
	case Right:
		if d.hasRight {
			return
		}
		d.hasRight = true
		if ok {
			distances.MustSet(currentTile.Y, currentTile.X, d)
		}
		currentTile.X += 1
	}

	// check the next tile
	// evoke move based on next directions
	letter, ok := tiles.Get(currentTile.Y, currentTile.X)
	if !ok {
		return //have fallen outside the matrix
	}
	// pureDistances := util.UnorderedMapToOrdered[DistanceEntry, int](*distances, func(y, x int, value DistanceEntry) int {
	// 	return value.distance
	// })
	// pureDistances.PrintEvenlySpaced(",")
	switch letter {
	case '.':
		move(currentTile, direction, count+1, tiles, distances)
	case '-':
		if direction != Right {
			move(currentTile, Left, count+1, tiles, distances)
		}
		if direction != Left {
			move(currentTile, Right, count+1, tiles, distances)
		}
	case '|':
		if direction != Down {
			move(currentTile, Up, count+1, tiles, distances)
		}
		if direction != Up {
			move(currentTile, Down, count+1, tiles, distances)
		}
	case '\\':
		switch direction {
		case Up:
			move(currentTile, Left, count+1, tiles, distances)
		case Down:
			move(currentTile, Right, count+1, tiles, distances)
		case Left:
			move(currentTile, Up, count+1, tiles, distances)
		case Right:
			move(currentTile, Down, count+1, tiles, distances)
		}
	case '/':
		switch direction {
		case Up:
			move(currentTile, Right, count+1, tiles, distances)
		case Down:
			move(currentTile, Left, count+1, tiles, distances)
		case Left:
			move(currentTile, Down, count+1, tiles, distances)
		case Right:
			move(currentTile, Up, count+1, tiles, distances)
		}
	default:
		panic("bad character")
	}
}
