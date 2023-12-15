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
func ColourFunction(tPreScaled float64) color.Color {
	t := tPreScaled + 0.1
	if t <= 0.2 {
		return color.RGBA{0, uint8(1275 * t), 255, 255}
	} else if t <= 0.4 {
		return color.RGBA{0, 255, uint8(1275 * (0.4 - t)), 0}
	} else if t <= 0.6 {
		return color.RGBA{uint8(1275 * (t - 0.4)), 255, 0, 255}
	} else if t <= 0.8 {
		return color.RGBA{255, uint8(1275 * (0.8 - t)), 0, 255}
	} else if t <= 1 {
		return color.RGBA{255, 0, uint8(1275 * (t - 0.8)), 255}
	} else {
		return color.RGBA{uint8(1275 * (1.2 - t)), 0, 255, 255}
	}
}

type Point2D struct {
	Y int
	X int
}

type Point3d struct {
	Z int
	Y int
	X int
}

func Power[N ~float64 | ~int](a, b N) N {
	return N(math.Pow(float64(a), float64(b)))
}

func Gcd(x, y int) int {
	if y == 0 || x == 0 {
		return x
	}
	return Gcd(min(x, y), max(x, y)%min(x, y))
}

func Egcd(x, y int) (int, int, int) {
	if x == 0 {
		return y, 0, 1
	}
	g, a, b := Egcd(y%x, x)
	return g, b - y/x*a, a
}

func Abs[V ~float64 | ~int](x V) V {
	if x < 0 {
		return -x
	}

	return x
}
