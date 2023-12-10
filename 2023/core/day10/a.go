package day10

import (
	"fmt"
	"math"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "10a",
		Short: "Day 10, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

type Tile struct {
	up      bool
	down    bool
	left    bool
	right   bool
	isStart bool
}

func partA(challenge *core.Input) int {
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
	distances.MustSet(sLocation.Y, sLocation.X, -1)

	return march(sLocation, tiles, &distances) - 1
}

func toTile(y, x int, letter rune) Tile {
	switch letter {
	case '-':
		return Tile{left: true, right: true}
	case '7':
		return Tile{left: true, down: true}
	case '|':
		return Tile{up: true, down: true}
	case 'J':
		return Tile{up: true, left: true}
	case 'L':
		return Tile{up: true, right: true}
	case 'F':
		return Tile{down: true, right: true}
	case 'S':
		return Tile{isStart: true}
	default:
		return Tile{}
	}
}

func march(sLocation util.Point2D, tiles util.UnorderedMatrix[Tile], distances *util.Matrix[int]) int {
	distance := 1
	loc := sLocation
	previousDirection := ""
	sTile := tiles.MustGet(sLocation.Y, sLocation.X)

	if sTile.up {
		previousDirection = "up"
	}
	if sTile.down {
		previousDirection = "down"
	}
	if sTile.left {
		previousDirection = "left"
	}
	if sTile.right {
		previousDirection = "right"
	}
	for {
		distance++
		tile := tiles.MustGet(loc.Y, loc.X)
		direction := previousDirection
		//march to the next location, making sure not to double back ever
		switch direction {
		case "up":
			if tile.down {
				previousDirection = "up" //we just came from the up tile by going down
				loc = util.Point2D{Y: loc.Y + 1, X: loc.X}
			} else if tile.left {
				previousDirection = "right"
				loc = util.Point2D{Y: loc.Y, X: loc.X - 1}
			} else if tile.right {
				previousDirection = "left"
				loc = util.Point2D{Y: loc.Y, X: loc.X + 1}
			}
		case "down":
			if tile.up {
				previousDirection = "down"
				loc = util.Point2D{Y: loc.Y - 1, X: loc.X}
			} else if tile.left {
				previousDirection = "right"
				loc = util.Point2D{Y: loc.Y, X: loc.X - 1}
			} else if tile.right {
				previousDirection = "left"
				loc = util.Point2D{Y: loc.Y, X: loc.X + 1}
			}
		case "left":
			if tile.down {
				previousDirection = "up"
				loc = util.Point2D{Y: loc.Y + 1, X: loc.X}
			} else if tile.up {
				previousDirection = "down"
				loc = util.Point2D{Y: loc.Y - 1, X: loc.X}
			} else if tile.right {
				previousDirection = "left"
				loc = util.Point2D{Y: loc.Y, X: loc.X + 1}
			}
		case "right":
			if tile.down {
				previousDirection = "up"
				loc = util.Point2D{Y: loc.Y + 1, X: loc.X}
			} else if tile.left {
				previousDirection = "right"
				loc = util.Point2D{Y: loc.Y, X: loc.X - 1}
			} else if tile.up {
				previousDirection = "down"
				loc = util.Point2D{Y: loc.Y - 1, X: loc.X}
			}
		}

		if distances.MustGet(loc.Y, loc.X) != 0 {
			distances.Draw()
			return int(math.Round(float64(distance) / 2.))
		}
		distances.MustSet(loc.Y, loc.X, distance)
	}

}
