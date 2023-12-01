package day1

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	t.Parallel()
	input := core.FromLiteral(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)
	// fmt.Print(strings.Join(input.aqaLineSlice(), ","))
	result := partA(input)

	require.Equal(t, 142, result)
}
