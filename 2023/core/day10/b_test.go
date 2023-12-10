package day10

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Parallel()
	input := core.FromLiteral(`............
..S-------7.
..|F-----7|.
..||.....||.
..||.....||.
..|L-7.F-J|.
..|..|.|..|.
..L--J.L--J.
............`)

	result := partB(input)

	require.Equal(t, 4, result)
}

func TestB2(t *testing.T) {
	t.Parallel()
	input := core.FromLiteral(`FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`)

	result := partB(input)

	require.Equal(t, 10, result)
}
