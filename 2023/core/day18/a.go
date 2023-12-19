package day18

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"regexp"
	"strconv"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "18a",
		Short: "Day 18, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

func partA(challenge *core.Input) int {
	util.EbitenSetup()
	colourMap := make(map[util.Point2D]color.RGBA)
	re := regexp.MustCompile(`([ULRD]) ([0-9]+) \(#([0-9a-f]{2})([0-9a-f]{2})([0-9a-f]{2})\)`)
	currentPoint := util.Point2D{Y: 0, X: 0}
	lines := challenge.LineSlice()
	for _, line := range lines {
		regexRes := re.FindStringSubmatch(line)
		colour := color.RGBA{R: hexToUint(regexRes[3]), G: hexToUint(regexRes[4]), B: hexToUint(regexRes[5]), A: 255}
		n, _ := strconv.Atoi(regexRes[2])
		for i := 1; i <= n; i++ {
			switch regexRes[1] {
			case "U":
				currentPoint = util.Point2D{Y: currentPoint.Y - 1, X: currentPoint.X}
				colourMap[currentPoint] = colour
			case "D":
				currentPoint = util.Point2D{Y: currentPoint.Y + 1, X: currentPoint.X}
				colourMap[currentPoint] = colour
			case "L":
				currentPoint = util.Point2D{Y: currentPoint.Y, X: currentPoint.X - 1}
				colourMap[currentPoint] = colour
			case "R":
				currentPoint = util.Point2D{Y: currentPoint.Y, X: currentPoint.X + 1}
				colourMap[currentPoint] = colour
			}
		}
	}
	minY := math.MaxInt
	maxY := 0
	minX := math.MaxInt
	maxX := 0
	for point := range colourMap {
		if point.Y < minY {
			minY = point.Y
		}
		if point.Y > maxY {
			maxY = point.Y
		}
		if point.X < minX {
			minX = point.X
		}
		if point.X > maxX {
			maxX = point.X
		}
	}

	tiles := util.NewUnorderedMatrix[color.RGBA](maxY-minY+3, maxX-minX+3)
	tiles.SetByRule(func(y, x int) color.RGBA {
		offsetPoint := util.Point2D{Y: y + minY - 1, X: x + minX - 1}
		return colourMap[offsetPoint]
	})
	// tiles.PrintEvenlySpaced(" ")
	img := image.NewRGBA(image.Rect(0, 0, tiles.GetWidth(), tiles.GetHeight()))
	for point, colour := range tiles.Iterator() {
		img.SetRGBA(point.X, point.Y, colour)
	}
	util.Image = img
	markedMatrix := util.UnorderedMapToOrdered[color.RGBA, int](tiles, func(y, x int, value color.RGBA) int {
		defaultColour := color.RGBA{}
		if value != defaultColour {
			return 1
		}
		return 0
	})
	currentPoint = util.Point2D{Y: 0, X: 0}
	queue := []util.Point2D{currentPoint}
	// markedMatrix.PrintEvenlySpaced("")
	// fmt.Println()
	for len(queue) > 0 {
		currentPoint = queue[0]
		queue = queue[1:]
		if num, ok := markedMatrix.Get(currentPoint.Y, currentPoint.X); ok && num != 1 && num != 2 {
			markedMatrix.MustSet(currentPoint.Y, currentPoint.X, 2)
			queue = append(queue,
				util.Point2D{Y: currentPoint.Y + 1, X: currentPoint.X},
				util.Point2D{Y: currentPoint.Y - 1, X: currentPoint.X},
				util.Point2D{Y: currentPoint.Y, X: currentPoint.X + 1},
				util.Point2D{Y: currentPoint.Y, X: currentPoint.X - 1},
			)
		}
	}
	// markedMatrix.PrintEvenlySpaced("")
	total := 0
	for point, value := range markedMatrix.Iterator() {
		if value != 2 {
			total += 1
		}
		if value == 0 {
			img.SetRGBA(point.X, point.Y, color.RGBA{255, 255, 255, 255})
		}
	}
	util.Image = img

	return total
}

func hexToUint(hex string) uint8 {
	val, _ := strconv.ParseUint(hex, 16, 16)
	return uint8(val)
}
