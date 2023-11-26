package util

import (
	"cmp"
	"fmt"
	"image/color"
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
	if t <= 0.2 {
		return color.RGBA{0, uint8(1275 * t), 255, 255}
	} else if t <= 0.4 {
		return color.RGBA{0, 255, uint8(1275 * (0.4 - t)), 0}
	} else if t <= 0.6 {
		return color.RGBA{uint8(1275 * (t - 0.4)), 255, 0, 255}
	} else if t <= 0.8 {
		return color.RGBA{255, uint8(1275 * (0.8 - t)), 0, 255}
	} else {
		return color.RGBA{255, 0, uint8(1275 * (t - 0.8)), 255}
	}
}
