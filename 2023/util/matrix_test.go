package util

import (
	"math"
	"testing"
	"time"
)

func TestMatrix(t *testing.T) {
	// EbitenSetup()
	// defer AwaitClosure()

	m := NewMatrix[int](40, 40)
	m.SetByRule(func(y, x int) int {
		return x * x * y
	})
	m.PrintEvenlySpaced(", ")
	m.Draw()
}

func TestThreeMatrix_Sphere(t *testing.T) {
	// EbitenSetup()
	// go AwaitClosure()
	r := 55
	m := NewThreeMatrix[int](2*r+1, 2*r+1, 2*r+1)
	m.SetByRule(func(z, y, x int) int {
		return min((x-r)*(x-r)+(y-r)*(y-r)+(z-r)*(z-r), r*r)
		// return min(math.Abs(float64(x-r))+math.Abs(float64(y-r))+math.Abs(float64(z-r)), float64(r))
	})
	// m.PrintEvenlySpaced(", ")
	for {
		for z := 0; z < 2*r; z++ {
			if !WindowBeingUsed {
				return
			}
			m.Draw(z)
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func TestThreeMatrix_Cube(t *testing.T) {
	// EbitenSetup()
	// go AwaitClosure()
	r := 40
	m := NewThreeMatrix[float64](2*r+1, 2*r+1, 2*r+1)
	m.SetByRule(func(z, y, x int) float64 {
		// return min((x-r)*(x-r)+(y-r)*(y-r)+(z-r)*(z-r), r*r)
		return min(math.Abs(float64(x-r))+math.Abs(float64(y-r))+math.Abs(float64(z-r)), float64(r))
	})
	// m.PrintEvenlySpaced(", ")
	for {
		for z := 0; z < 2*r; z++ {
			if !WindowBeingUsed {
				return
			}
			m.Draw(z)
			time.Sleep(10 * time.Millisecond)
		}
	}
}
