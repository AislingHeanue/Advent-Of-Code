package day5

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
)

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = partA(core.FromFile())
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = partB(core.FromFile())
	}
}
