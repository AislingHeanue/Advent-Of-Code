package day22

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
		Use:   "22a",
		Short: "Day 22, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

type Box struct {
	grounded  bool
	p1        util.Point3D
	p2        util.Point3D // p2.X > p1.X, p2.Y > p1.Y, p2.Z > p1.Z
	restingOn map[int]bool
	under     map[int]bool
}

func partA(challenge *core.Input) int {
	boxes, matrix := DoSetup(challenge)
	// POPULATE restingOn FOR EACH BOX (and keep in mind the special case for Z = 1)
	for i, box := range boxes {
		restingMap := box.restingOn
		for x := box.p1.X; x <= box.p2.X; x++ {
			for y := box.p1.Y; y <= box.p2.Y; y++ {
				if a := matrix.Get(box.p1.Z-1, y, x); a != -1 {
					restingMap[a] = true
				}
			}
		}
		boxes[i].restingOn = restingMap
	}
	// CHECK ALL RESTING MAPS OF SIZE ONE
	cannotBeRemoved := make(map[int]bool)
	for _, box := range boxes {
		if restingMap := box.restingOn; len(restingMap) == 1 {
			for key := range restingMap {
				cannotBeRemoved[key] = true
			}
		}
	}

	return len(boxes) - len(cannotBeRemoved)
}

func DoSetup(challenge *core.Input) ([]Box, util.ThreeMatrix[int]) {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	boxes := core.InputMap[Box](challenge, ToBox)
	// 0,0,0 defines one corner of the matrix, find the other corner
	maxPoint := util.Point3D{Z: 0, Y: 0, X: 0}
	for _, box := range boxes {
		if box.p2.X > maxPoint.X {
			maxPoint.X = box.p2.X
		}
		if box.p2.Y > maxPoint.Y {
			maxPoint.Y = box.p2.Y
		}
		if box.p2.Z > maxPoint.Z {
			maxPoint.Z = box.p2.Z
		}
	}
	matrix := util.NewThreeMatrix[int](maxPoint.Z+1, maxPoint.Y+1, maxPoint.X+1) //the height here is actually Z. Don't worry about it
	matrix.Fill(-1)
	// boxes are indexed by their position in this array (pretty convenient actually)
	// PUT THE BOXES IN THE MATRIX
	for id, box := range boxes {
		for x := box.p1.X; x <= box.p2.X; x++ {
			for y := box.p1.Y; y <= box.p2.Y; y++ {
				for z := box.p1.Z; z <= box.p2.Z; z++ {
					if matrix.Get(z, y, x) != -1 {
						panic(fmt.Sprintf("two boxes in the same place? %d,%d,%d", x, y, z))
					}
					matrix.Set(z, y, x, id)
				}
			}
		}
	}
	// MAKE ANY FLOATING BOXES FALL
	for z := 1; z < matrix.GetDepth(); z++ {
		for x := 0; x < matrix.GetWidth(); x++ {
			for y := 0; y < matrix.GetHeight(); y++ {
				value := matrix.Get(z, y, x)
				if value == -1 {
					continue
				}
				if !boxes[value].grounded {
					boxes[value].fall(&matrix, value)
				}
			}
		}
	}
	// POPULATE restingOn FOR EACH BOX (and keep in mind the special case for Z = 1)
	for i, box := range boxes {
		restingMap := box.restingOn
		for x := box.p1.X; x <= box.p2.X; x++ {
			for y := box.p1.Y; y <= box.p2.Y; y++ {
				if a := matrix.Get(box.p1.Z-1, y, x); a != -1 {
					restingMap[a] = true
					underMap := boxes[a].under
					underMap[i] = true
					boxes[a].under = underMap
				}
			}
		}
		boxes[i].restingOn = restingMap
	}

	return boxes, matrix
}

func ToBox(line string) Box {
	re := regexp.MustCompile(`(\d+),(\d+),(\d+)~(\d+),(\d+),(\d+)`)
	regexRes := re.FindStringSubmatch(line)
	points := [6]int{}
	for i := 1; i < 7; i++ {
		points[i-1], _ = strconv.Atoi(regexRes[i])
	}
	return Box{
		grounded: min(points[2], points[5]) == 1,
		p1: util.Point3D{
			X: min(points[0], points[3]),
			Y: min(points[1], points[4]),
			Z: min(points[2], points[5]),
		},
		p2: util.Point3D{
			X: max(points[0], points[3]),
			Y: max(points[1], points[4]),
			Z: max(points[2], points[5]),
		},
		restingOn: make(map[int]bool),
		under:     make(map[int]bool),
	}
}

func (b *Box) fall(matrix *util.ThreeMatrix[int], index int) {
	d := b.p1.Z - 1 // attempt to fall all the way to z = 1
	for i := 0; i < d; i++ {
		b.p1.Z -= 1
		b.p2.Z -= 1
		okToFall := true
		for x := b.p1.X; x <= b.p2.X; x++ {
			for y := b.p1.Y; y <= b.p2.Y; y++ {
				if matrix.Get(b.p1.Z, y, x) != -1 {
					okToFall = false
				}
			}
		}
		if !okToFall {
			b.p1.Z += 1
			b.p2.Z += 1
			b.grounded = true // not sure how confident about that i am yet
			return
		}
		for x := b.p1.X; x <= b.p2.X; x++ {
			for y := b.p1.Y; y <= b.p2.Y; y++ {
				matrix.Set(b.p1.Z, y, x, index)
				matrix.Set(b.p2.Z+1, y, x, -1)
			}
		}
	}
	b.grounded = true
}
