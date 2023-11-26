package util

import (
	"testing"
)

func init() {
	EbitenSetup()
	WindowBeingUsed = true
}

func awaitClosure() {
	if WindowBeingUsed {
		<-WindowClosureChan
	}
}

func TestMatrix(t *testing.T) {
	defer awaitClosure()

	m := NewMatrix[int](40, 40)
	m.SetByRule(func(y, x int) int {
		return x * x * y
	})
	m.PrintEvenlySpaced(", ")
	m.Draw()
}
