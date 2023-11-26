package util

import (
	"cmp"
	"fmt"
	"image/color"
	"math"
)

func Between[V cmp.Ordered](minimum, value, maximum V) V {
	if maximum < minimum {
		fmt.Printf("ERROR: max %v is less than min %v", maximum, minimum)
	}
	return max(min(value, maximum), minimum)
}

// not sure how to use a predefined colour spectrum, so I'll make my own
// works best when provided a value of t between zero and one
// for the full spectrum, use 0 and 2 instead
func ColourFunction(t float64) color.Color {
	return color.RGBA{
		R: uint8(math.Min(255, (128 * (1 + math.Sin(float64(t)*math.Pi))))),
		G: uint8(math.Min(255, (128 * (1 + math.Sin((2*math.Pi/3)+float64(t)*math.Pi))))),
		B: uint8(math.Min(255, (128 * (1 + math.Sin((4*math.Pi/3)+float64(t)*math.Pi))))),
		A: 255,
	}
}
