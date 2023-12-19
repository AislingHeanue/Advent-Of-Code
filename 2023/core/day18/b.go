package day18

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "18b",
		Short: "Day 18, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

type HorizontalLine struct {
	y  int
	x1 int
	x2 int
}

type VerticalLine struct {
	x  int
	y1 int
	y2 int
}

func partB(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	re := regexp.MustCompile(`#([0-9a-f]{5})([0-3])`)
	lines := challenge.LineSlice()
	currentPoint := util.Point2D{Y: 0, X: 0}
	minX, minY, maxX, maxY := 0, 0, 0, 0
	horizontalLines := []HorizontalLine{}
	verticalLines := []VerticalLine{}
	xs := []int{}
	ys := []int{}
	for _, line := range lines {
		regexRes := re.FindStringSubmatch(line)
		distance64, _ := strconv.ParseInt(regexRes[1], 16, 64)
		distance := int(distance64)

		switch regexRes[2] {
		case "0":
			x := currentPoint.X + distance
			horizontalLines = append(horizontalLines, HorizontalLine{y: currentPoint.Y, x1: currentPoint.X, x2: x})
			xs = append(xs, x, currentPoint.X)
			currentPoint = util.Point2D{Y: currentPoint.Y, X: x}
			if x > maxX {
				maxX = x
			}
		case "1":
			y := currentPoint.Y + distance
			verticalLines = append(verticalLines, VerticalLine{x: currentPoint.X, y1: currentPoint.Y, y2: y})
			ys = append(ys, currentPoint.Y, y)
			currentPoint = util.Point2D{Y: y, X: currentPoint.X}
			if y > maxY {
				maxY = y
			}
		case "2":
			x := currentPoint.X - distance
			horizontalLines = append(horizontalLines, HorizontalLine{y: currentPoint.Y, x1: x, x2: currentPoint.X})
			xs = append(xs, x, currentPoint.X)
			currentPoint = util.Point2D{Y: currentPoint.Y, X: x}
			if x < minX {
				minX = x
			}
		case "3":
			y := currentPoint.Y - distance
			verticalLines = append(verticalLines, VerticalLine{x: currentPoint.X, y1: y, y2: currentPoint.Y})
			ys = append(ys, currentPoint.Y, y)
			currentPoint = util.Point2D{Y: y, X: currentPoint.X}
			if y < minY {
				minY = y
			}

		}
	}
	sort.Slice(ys, func(i, j int) bool {
		return ys[i] < ys[j]
	})
	sort.Slice(xs, func(i, j int) bool {
		return xs[i] < xs[j]
	})
	ysUnique := []int{}
	for _, y := range ys {
		if len(ysUnique) == 0 {
			ysUnique = append(ysUnique, y)
		} else if ysUnique[len(ysUnique)-1] != y {
			ysUnique = append(ysUnique, y)
		}
	}
	xsUnique := []int{}
	for _, x := range xs {
		if len(xsUnique) == 0 {
			xsUnique = append(xsUnique, x)
		} else if xsUnique[len(xsUnique)-1] != x {
			xsUnique = append(xsUnique, x)
		}
	}

	count := 0
	insideShape := false
	for i, y := range ysUnique {
		nextY := y + 1
		if i != len(ysUnique)-1 {
			nextY = ysUnique[i+1]
		}
		if nextY == y {
			panic("there's a dupe here")
		}

		for j, x := range xsUnique {
			nextX := x + 1
			if j != len(xsUnique)-1 {
				nextX = xsUnique[j+1]
			}
			if nextX == x {
				panic("there's a dupe here")
			}
			v := inVerticalLine(verticalLines, y, x)
			h := inHorizontalLine(horizontalLines, y, x)
			if v {
				insideShape = !insideShape
				// the left edge of this box is always inside, count it
				if !insideShape {
					count += (nextY - y)
				}
			}
			if h {
				//the top edge of this box is always inside, count it
				if !insideShape {
					count += (nextX - x)
					// case where the block is a corner, so the corner itself was double counted
					if v {
						count -= 1 // the corner case
					}
				}
			}
			if insideShape {
				count += (nextY - y) * (nextX - x)
			}
			if !insideShape && !v && !h {
				// special case of the vertical line check for a bottom right corner
				for _, line := range verticalLines {
					if line.x == x {
						if line.y1 <= y && y <= line.y2 {
							count += 1
						}
					}
				}
			}

		}
		insideShape = false
	}

	return count
}

func inVerticalLine(lines []VerticalLine, y, x int) bool { // TODO: literally corner cases
	for _, line := range lines {
		if line.x == x {
			if line.y2 < line.y1 {
				panic("invalid line")
			}
			if line.y1 <= y && y < line.y2 {
				return true
			}
		}
	}

	return false
}

func inHorizontalLine(lines []HorizontalLine, y, x int) bool { // TODO: literally corner cases
	for _, line := range lines {
		if line.x2 < line.x1 {
			panic("invalid line")
		}
		if line.y == y {
			if line.x1 <= x && x < line.x2 {
				return true
			}
		}
	}

	return false
}
