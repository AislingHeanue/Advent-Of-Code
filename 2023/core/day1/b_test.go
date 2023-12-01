package day1

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Parallel()
	input := core.FromLiteral(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)
	// fmt.Print(strings.Join(input.LineSlice(), ","))
	result := partB(input)

	require.Equal(t, 281, result)
}
