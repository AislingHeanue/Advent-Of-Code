package util

import (
	"testing"
	"time"
)

func init() {
	EbitenSetup()
	WindowBeingUsed = true
}

func TestMatrix(t *testing.T) {
	defer AwaitClosure()

	m := NewMatrix[int](40, 40)
	m.SetByRule(func(y, x int) int {
		return x * x * y
	})
	m.PrintEvenlySpaced(", ")
	m.Draw()
}

func TestThreeMatrix(t *testing.T) {
	go AwaitClosure()
	r := 10
	m := NewThreeMatrix[int](2*r+1, 2*r+1, 2*r+1)
	m.SetByRule(func(z, y, x int) int {
		return min((x-r)*(x-r)+(y-r)*(y-r)+(z-r)*(z-r), r*r)
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
