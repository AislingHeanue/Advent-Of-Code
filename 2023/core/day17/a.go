package day17

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "17a",
		Short: "Day 17, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type PointWithDistance struct {
	point    util.Point3D
	distance int
}

// idea: let's do dijkstra (no!)
// idea: let's do BFS but we can move 1,2,3 steps each time, and we must turn (also no!)
// idea: let's do dijkstra but you can hop up to 3 nodes at once, and the state space is 3D because last-direction is important

func partA(challenge *core.Input) int {
	tiles := challenge.TileMap()
	source := util.Point3D{Z: 1, Y: 0, X: 0}
	source2 := util.Point3D{Z: 3, Y: 0, X: 0}
	distances := util.NewThreeMatrix[int](4, tiles.GetHeight(), tiles.GetWidth())
	distances.Fill(100000) // "infinite"
	distances.Set(source.Z, source.Y, source.X, 0)
	distances.Set(source2.Z, source2.Y, source2.X, 0)
	queue := []PointWithDistance{}
	queue = append(queue, PointWithDistance{source, 0}, PointWithDistance{source2, 0})
	for len(queue) > 0 {
		currentTile := queue[0]
		queue = queue[1:]
		for z := 0; z < 4; z++ {
		NewDirection:
			for i := 1; i < 4; i++ {
				// NewTile:
				switch Direction(z) {
				case Up:
					if Direction(currentTile.point.Z) == Up || Direction(currentTile.point.Z) == Down {
						break NewDirection
					}
					neighbour := util.Point3D{Z: z, Y: currentTile.point.Y - i, X: currentTile.point.X}
					newTemp := distances.Get(currentTile.point.Z, currentTile.point.Y, currentTile.point.X)
					for j := 0; j < i; j++ {
						temp, ok := tiles.Get(neighbour.Y+j, neighbour.X)
						if !ok {
							break NewDirection // if we've fallen out of the matrix, incrementing `i` won't get up back into it.
						}
						newTemp += int(temp - '0')
					}
					if newTemp < distances.Get(neighbour.Z, neighbour.Y, neighbour.X) {
						distances.Set(neighbour.Z, neighbour.Y, neighbour.X, newTemp)
						queue = append(queue, PointWithDistance{neighbour, i})
					}
				case Down:
					if Direction(currentTile.point.Z) == Up || Direction(currentTile.point.Z) == Down {
						break NewDirection
					}
					neighbour := util.Point3D{Z: z, Y: currentTile.point.Y + i, X: currentTile.point.X}
					newTemp := distances.Get(currentTile.point.Z, currentTile.point.Y, currentTile.point.X)
					for j := 0; j < i; j++ {
						temp, ok := tiles.Get(neighbour.Y-j, neighbour.X)
						if !ok {
							break NewDirection // if we've fallen out of the matrix, incrementing `i` won't get up back into it.
						}
						newTemp += int(temp - '0')
					}
					if newTemp < distances.Get(neighbour.Z, neighbour.Y, neighbour.X) {
						distances.Set(neighbour.Z, neighbour.Y, neighbour.X, newTemp)
						queue = append(queue, PointWithDistance{neighbour, i})
					}
				case Left:
					if Direction(currentTile.point.Z) == Left || Direction(currentTile.point.Z) == Right {
						break NewDirection
					}
					neighbour := util.Point3D{Z: z, Y: currentTile.point.Y, X: currentTile.point.X - i}
					newTemp := distances.Get(currentTile.point.Z, currentTile.point.Y, currentTile.point.X)
					for j := 0; j < i; j++ {
						temp, ok := tiles.Get(neighbour.Y, neighbour.X+j)
						if !ok {
							break NewDirection // if we've fallen out of the matrix, incrementing `i` won't get up back into it.
						}
						newTemp += int(temp - '0')
					}
					if newTemp < distances.Get(neighbour.Z, neighbour.Y, neighbour.X) {
						distances.Set(neighbour.Z, neighbour.Y, neighbour.X, newTemp)
						queue = append(queue, PointWithDistance{neighbour, i})
					}
				case Right:
					if Direction(currentTile.point.Z) == Left || Direction(currentTile.point.Z) == Right {
						break NewDirection
					}
					neighbour := util.Point3D{Z: z, Y: currentTile.point.Y, X: currentTile.point.X + i}
					newTemp := distances.Get(currentTile.point.Z, currentTile.point.Y, currentTile.point.X)
					for j := 0; j < i; j++ {
						temp, ok := tiles.Get(neighbour.Y, neighbour.X-j)
						if !ok {
							break NewDirection // if we've fallen out of the matrix, incrementing `i` won't get up back into it.
						}
						newTemp += int(temp - '0')
					}
					if newTemp < distances.Get(neighbour.Z, neighbour.Y, neighbour.X) {
						distances.Set(neighbour.Z, neighbour.Y, neighbour.X, newTemp)
						queue = append(queue, PointWithDistance{neighbour, i})
					}
				}
			}
		}
	}
	return min(
		distances.Get(int(Down), tiles.GetHeight()-1, tiles.GetWidth()-1),
		distances.Get(int(Right), tiles.GetHeight()-1, tiles.GetWidth()-1),
	)

}
