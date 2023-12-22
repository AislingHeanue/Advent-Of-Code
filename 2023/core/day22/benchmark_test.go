package day22

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
)

func BenchmarkA(b *testing.B) {
	util.ForceNoWindow = true
	for i := 0; i < b.N; i++ {
		_ = partA(core.FromFile())
	}
}

func BenchmarkB(b *testing.B) {
	util.ForceNoWindow = true
	for i := 0; i < b.N; i++ {
		_ = partB(core.FromFile())
	}
}
