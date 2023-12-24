package day24

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "24a",
		Short: "Day 24, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

type Stone struct {
	Position util.Point3D
	Velocity util.Point3D
}

func partA(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	stones := core.InputMap(challenge, toStone)
	minDimension := 7.
	maxDimension := 27.
	if stones[0].Position.X > 25 {
		minDimension = 200000000000000.
		maxDimension = 400000000000000.
	}
	total := 0
	for i, stone1 := range stones {
		for j, stone2 := range stones {
			if i < j {
				y, x, ok := getIntersection(stone1, stone2)
				if !ok { // parallel lines
					continue
				}
				if x < maxDimension &&
					x > minDimension &&
					y < maxDimension &&
					y > minDimension {
					total += 1
				}
			}
		}
	}
	return total
}

func toStone(line string) Stone {
	re := regexp.MustCompile(`(\d+),\s+(\d+),\s+(\d+)\s+@\s+(-?\d+),\s+(-?\d+),\s+(-?\d+)`)
	res := re.FindStringSubmatch(line)
	return Stone{
		Position: util.Point3D{X: convert(res[1]), Y: convert(res[2]), Z: convert(res[3])},
		Velocity: util.Point3D{X: convert(res[4]), Y: convert(res[5]), Z: convert(res[6])},
	}
}

func convert(line string) int {
	res, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	return res
}

func getIntersection(s1, s2 Stone) (y float64, x float64, exists bool) {
	// parallel
	if s1.Velocity.X*s2.Velocity.Y-s1.Velocity.Y*s2.Velocity.X == 0 {
		return 0., 0., false
	}
	// s1.Position.X + s1.Velocity.X * t = s2.Position.X + s2.Velocity.X * s
	// s1.Position.Y + s1.Velocity.Y * t = s2.Position.Y + s2.Velocity.Y * s
	//
	// s1.Vx*t - s2.Vx*s = (s2.Px - s1.Px)
	// s1.Vy*t - s2.Vy*s = (s2.Py - s1.Py)
	// s1.Vy*t - (s1.Vy/s1.Vx)s2.Vx*s = (s1.Vy/s1.Vx)(s2Px - s1.Px)
	// ((s1.Vy/s1.Vx)s2.Vx - s2.Vy)*s = (s2.Py - s1.Py) - (s1.Vy/s1.Vx)(s2Px - s1.Px)
	// s = ((s2.Py - s1.Py) - (s1.Vy/s1.Vx)(s2Px - s1.Px))/((s1.Vy/s1.Vx)s2.Vx - s2.Vy)
	s := ((float64(s2.Position.Y) - float64(s1.Position.Y)) - (float64(s1.Velocity.Y)/float64(s1.Velocity.X))*(float64(s2.Position.X)-float64(s1.Position.X))) /
		((float64(s1.Velocity.Y)/float64(s1.Velocity.X))*float64(s2.Velocity.X) - float64(s2.Velocity.Y))
	t := (float64(s2.Position.X) - float64(s1.Position.X) + float64(s2.Velocity.X)*s) / (float64(s1.Velocity.X))
	if t < 0 || s < 0 {
		return 0., 0., false
	}
	return float64(s2.Position.Y) + s*float64(s2.Velocity.Y), float64(s2.Position.X) + s*float64(s2.Velocity.X), true
}
