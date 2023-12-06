package day5

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "5b",
		Short: "Day 5, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

// we have a list of non overlapping (I hope) intervals in layer 0, which will lead to another set of
// possible intervals in layer 1 after going through this transformation, its a surjective mapping so
// the total size of the interval will get smaller each step.
// Then we just need the start of the first interval and we win.

// tools needed: interval union, and interval difference (needed for the exception case).
// Actually, interval intersection would also be useful for checking
// a transition against a set of intervals and returning a set of (transformed) intervals

func partB(challenge *core.Input) int {
	lines := challenge.LineSlice()
	re := regexp.MustCompile(`\d+`)
	maps := make([][][]int, 7)
	i := 1
	mapsIndex := -1
	for i < len(lines) {
		strings := re.FindAllString(lines[i], 3)
		if len(strings) != 3 {
			i += 2
			mapsIndex++
			maps[mapsIndex] = [][]int{}
		} else {
			numbers := make([]int, 3)
			for i, str := range strings {
				numbers[i], _ = strconv.Atoi(str)
			}
			maps[mapsIndex] = append(maps[mapsIndex], numbers)
			i++
		}
	}

	seeds := re.FindAllString(lines[0], -1)
	sets := make([]IntervalSet, 8)
	for i := 0; i < len(seeds); i += 2 {
		start, _ := strconv.Atoi(seeds[i])
		length, _ := strconv.Atoi(seeds[i+1])
		sets[0].Union(Interval{start, start + length}) // end point is not in interval
	}

	// filter by each rule according to "input and length", and transform the result in a new interval according to "output"
	for layer := 0; layer < 7; layer++ {
		oldLayer := sets[layer]
		sets[layer+1] = IntervalSet{}
		for _, filter := range maps[layer] {
			var removed IntervalSet
			removed, oldLayer = oldLayer.FilterAndRemoveByInterval(Interval{filter[1], filter[1] + filter[2]})
			removed.Offset(filter[0] - filter[1])
			sets[layer+1].UnionSet(removed)
		}
		sets[layer+1].UnionSet(oldLayer)
	}

	// fmt.Println(sets[7])
	return sets[7][0].start
}

type Interval struct {
	start int
	end   int
}

type IntervalSet []Interval

func (i1 Interval) FilterAndRemoveByInterval(i2 Interval) (Interval, IntervalSet) { //removed,filteredInterval (set because it may split in two)
	// lets actually design something
	// s : start end start end start end
	// i : start end

	// cases
	// |----|
	//         |----|
	// removed: nothing, set: unchanged

	// |----|
	//   |----|
	// removed: [i.start,s.interval.end) s: remove s.interval, add [s.interval.start,i.start)

	//     |----|
	// |-----|
	// removed: [i.start, s.interval.end) s: remove s.interval, add [i.start,s.interval.end)

	//  |---|
	// |-----|
	// removed: [s.interval.start,s.interval.end) s: remove s.interval

	points := []int{}
	points = append(points, i1.start, i1.end, i2.start, i2.end)
	removed := Interval{}
	newS := IntervalSet{}
	removedStart := -1
	newSStart := -1
	inTopInterval := false
	inFilterInterval := false
	sort.Slice(points, func(i, j int) bool {
		return points[i] < points[j]
	})
	for _, point := range points {
		inTopInterval = point >= i1.start && point < i1.end
		inFilterInterval = point >= i2.start && point < i2.end
		if inTopInterval && !inFilterInterval {
			if newSStart == -1 {
				newSStart = point
			} else {
				newS.Union(Interval{newSStart, point})
				newSStart = -1
			}
			if removedStart != -1 {
				removed = Interval{removedStart, point}
				removedStart = -1
			}
		}
		if inTopInterval && inFilterInterval {
			removedStart = point
			if newSStart != -1 {
				newS.Union(Interval{newSStart, point})
				newSStart = -1
			}
		}
		if !inTopInterval {
			if newSStart != -1 {
				newS.Union(Interval{newSStart, point})
				newSStart = -1
			}
			if removedStart != -1 {
				removed = Interval{removedStart, point}
				removedStart = -1
			}
		}
	}

	return removed, newS
}

func (s IntervalSet) FilterAndRemoveByInterval(i Interval) (IntervalSet, IntervalSet) {
	newS := IntervalSet{}
	removedS := IntervalSet{}
	for _, interval := range s {
		removed, toAdd := interval.FilterAndRemoveByInterval(i)
		newS.UnionSet(toAdd)
		removedS.Union(removed)
	}

	return removedS, newS
}

func (s *IntervalSet) UnionSet(s2 IntervalSet) {
	for _, i := range s2 {
		(*s).Union(i)
	}
}

func (s *IntervalSet) Union(i Interval) {
	if i.start > i.end {
		fmt.Println("invalid interval")
		return
	}
	if i.start != 0 || i.end != 0 {
		*s = append(*s, i)
	}
	s.Flatten()
}

func (s *IntervalSet) Flatten() {
	points := s.GetPoints()
	newS := IntervalSet{}
	start := -1
	for k := 0; k < len(points); k++ {
		if start == -1 {
			start = points[k]
		}
		presentInInterval := false
		for _, i := range *s {
			if points[k] >= i.start && points[k] < i.end {
				presentInInterval = true
			}
		}
		if !presentInInterval {
			newS = append(newS, Interval{start, points[k]})
			start = -1
		}
	}
	*s = newS
}

func (s IntervalSet) GetPoints() []int {
	points := []int{}
	for _, i := range s {
		points = append(points, i.start, i.end)
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i] < points[j]
	})

	return points
}

func (s *IntervalSet) Offset(value int) {
	for i, interval := range *s {
		(*s)[i] = Interval{interval.start + value, interval.end + value}
	}
}
