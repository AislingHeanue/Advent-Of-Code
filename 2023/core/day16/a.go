package day16

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "16a",
		Short: "Day 16, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

var Test bool = false
var DrawMatrix util.Matrix[int] // contains information about mirrors, as well as the distance of each step from the starting point
var MaxCount int = 0

type Direction int

const (
	Up Direction = iota
	Right
	Left
	Down
)

type DistanceEntry struct {
	distance int
	hasUp    bool
	hasDown  bool
	hasLeft  bool
	hasRight bool
}

func partA(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	util.EbitenSetup()
	tiles := challenge.TileMap()
	distances := util.MapToUnordered[rune, DistanceEntry](tiles, func(y, x int, value rune) DistanceEntry {
		return DistanceEntry{0, false, false, false, false}
	})
	direction := Right
	currentTile := util.Point2D{Y: 0, X: -1} //start outside the matrix
	move(currentTile, direction, 0, tiles, &distances)
	total := 0
	for _, value := range distances.Iterator() {
		if value.distance != 0 {
			total += 1
		}
	}
	bigNumber := 100000
	DrawMatrix = util.UnorderedMapToOrdered[DistanceEntry, int](distances, func(y, x int, value DistanceEntry) int {
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
	DrawMatrix.Draw()

	return total
}

func move(currentTile util.Point2D, direction Direction, count int, tiles util.Matrix[rune], distances *util.UnorderedMatrix[DistanceEntry]) {
	d, _ := distances.Get(currentTile.Y, currentTile.X)

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
		_ = distances.Set(currentTile.Y, currentTile.X, d)
		currentTile.Y -= 1
	case Down:
		if d.hasDown {
			return
		}
		d.hasDown = true

		_ = distances.Set(currentTile.Y, currentTile.X, d)
		currentTile.Y += 1
	case Left:
		if d.hasLeft {
			return
		}
		d.hasLeft = true
		_ = distances.Set(currentTile.Y, currentTile.X, d)
		currentTile.X -= 1
	case Right:
		if d.hasRight {
			return
		}
		d.hasRight = true
		_ = distances.Set(currentTile.Y, currentTile.X, d)

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
