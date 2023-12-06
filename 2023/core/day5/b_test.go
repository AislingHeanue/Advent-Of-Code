package day5

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Parallel()
	input := core.FromLiteral(`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`)

	result := partB(input)

	require.Equal(t, 46, result)
}

func TestUnion(t *testing.T) {
	i1 := Interval{1, 4}
	i2 := Interval{6, 8}
	i3 := Interval{2, 7}

	result := IntervalSet{}
	result.Union(i1)
	result.Union(i2)
	result.Union(i3)
	require.Equal(t, IntervalSet{Interval{1, 8}}, result)
}

func TestRemove(t *testing.T) {
	i2 := Interval{51, 52}
	i4 := Interval{50, 98}

	set := IntervalSet{}
	set.Union(i4)
	removed, set := set.FilterAndRemoveByInterval(i2)

	require.Equal(t, IntervalSet{Interval{51, 52}}, removed)
	require.Equal(t, IntervalSet{Interval{50, 51}, Interval{52, 98}}, set)

}

func TestRemove2(t *testing.T) {
	i2 := Interval{53, 61}
	i4 := Interval{51, 55}

	set := IntervalSet{}
	set.Union(i4)
	removed, set := set.FilterAndRemoveByInterval(i2)

	require.Equal(t, IntervalSet{Interval{53, 55}}, removed)
	require.Equal(t, IntervalSet{Interval{51, 53}}, set)
}

func TestRemove3(t *testing.T) {
	i2 := Interval{53, 61}
	i4 := Interval{51, 55}

	set := IntervalSet{}
	set.Union(i2)
	removed, set := set.FilterAndRemoveByInterval(i4)

	require.Equal(t, IntervalSet{Interval{53, 55}}, removed)
	require.Equal(t, IntervalSet{Interval{55, 61}}, set)
}

func TestRemove4(t *testing.T) {
	i1 := Interval{25, 95}
	i2 := IntervalSet{Interval{96, 120}}

	removed, set := i2.FilterAndRemoveByInterval(i1)

	require.Equal(t, IntervalSet{}, removed)
	require.Equal(t, IntervalSet{Interval{96, 120}}, set)
}
