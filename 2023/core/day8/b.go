package day8

import (
	"fmt"
	"regexp"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "8b",
		Short: "Day 8, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

var reB = regexp.MustCompile(`([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)`)
var nodeMap map[string]Node

func partB(challenge *core.Input) int {
	nodes := core.InputMap(challenge, ExtractNodeB)
	nodeMap = make(map[string]Node)
	instructions := ""
	starterLabels := []string{}
	for _, node := range nodes {
		if node.label != "" {
			nodeMap[node.label] = node
			if node.label[2] == 'A' {
				starterLabels = append(starterLabels, node.label)
			}
		} else {
			instructions = node.text
		}
	}

	cycles := make([]Cycle, len(starterLabels))
	for i, label := range starterLabels {
		cycles[i] = getCycle(label, instructions)
		cycles[i] = getZIndexes(cycles[i], instructions)
	}

	// For each cycle, they will have an associated set of possible lines their values can fall on and give a valid result.
	// The first step is to list out every possible line for each cycle.
	lines := [][]Line{}
	for i, cycle := range cycles {
		lines = append(lines, []Line{})
		for _, z := range cycle.zIndexesInCycle {
			lines[i] = append(lines[i], Line{start: cycle.start + z, slope: cycle.length})
		}
	}

	// next, preform a depth first search to enumerate every combination of paths that can be taken, using one
	// of the lines for each cycle at a time.
	// After this, take the n*lines you've obtained and turn them into one big line.
	paths := [][]Line{}
	currentPath := []Line{}
	minimum := -1
	makePaths(lines, currentPath, 0, &paths)
	for _, path := range paths {
		aggregatedLine := path[0]
		for i := 1; i < len(path); i++ {
			// fmt.Println(aggregatedLine.slope * path[i].slope / util.Gcd(aggregatedLine.slope, path[i].slope))
			aggregatedLine.product(path[i])
		}
		if (minimum == -1 || minimum > aggregatedLine.start) && aggregatedLine.slope > 0 {
			minimum = aggregatedLine.start
		}
	}

	return minimum

}

func makePaths(lines [][]Line, currentPath []Line, index int, paths *[][]Line) {
	if index == len(lines) {
		*paths = append(*paths, currentPath)
		return
	}
	for _, line := range lines[index] {
		newPath := append(currentPath, line)
		makePaths(lines, newPath, index+1, paths)
	}
}

// given lines a1+b1*x = n = a2+b2*y
// find the possible values of n in the form a3+b3*z
// a1-a2 = b1x - b2y
// the rest is taken from this textbook
// https://math.libretexts.org/Courses/Mount_Royal_University/MATH_2150%3A_Higher_Arithmetic/5%3A_Diophantine_Equations/5.1%3A_Linear_Diophantine_Equations#:~:text=A%20Linear%20Diophantine%20equation%20(LDE,and%20y%20are%20unknown%20variables.
func (l1 *Line) product(l2 Line) {
	d, _, y0 := util.Egcd(l1.slope, l2.slope)
	c := l1.start - l2.start
	if c%d != 0 {
		*l1 = Line{0, -1} // there are no solutions
		return
	}
	ySlope := l1.slope / d
	yStart := (c / d * y0) % ySlope
	slope := l2.slope * ySlope
	start := (l2.start + l2.slope*yStart) % slope
	if start == 0 {
		start = slope
	}
	*l1 = Line{start: start, slope: slope}
}

type Line struct {
	start int
	slope int
}

type Cycle struct {
	beginningLabel  string
	start           int
	length          int
	zIndexesInCycle []int
}

type CacheEntry struct {
	string
	int
}

// this answer assumes that
// 1. Z will be present in the cycle
// 2. None of the Z's that appear before the cycle begins will be the one in the actual answer
// 3. The instruction list itself is not repeating
// I can make assumptions 1 and 2 since I know the final answer is much much larger than the length of my input.
func getCycle(start string, instructions string) Cycle {
	nodeCache := make(map[CacheEntry]int)
	count := 0
	label := start
	instructionsIndex := 0
	found := false
	cycleStart := -1
	for {
		if found {
			return Cycle{start, cycleStart, count - cycleStart, []int{}}
		} else {
			nodeCache[CacheEntry{label, instructionsIndex}] = count
		}
		label = getNext(label, rune(instructions[instructionsIndex]))
		count++
		instructionsIndex++
		cycleStart, found = nodeCache[CacheEntry{label, instructionsIndex}]

		if instructionsIndex == len(instructions) {
			instructionsIndex = 0
		}
	}
}

func getZIndexes(cycle Cycle, instructions string) Cycle {
	indexesInCycle := []int{}
	label := cycle.beginningLabel
	instructionsIndex := 0
	count := 0
	for count < cycle.start {
		label = getNext(label, rune(instructions[instructionsIndex]))
		instructionsIndex++
		count++
		if instructionsIndex == len(instructions) {
			instructionsIndex = 0
		}
	}
	repeatedCacheEntry := CacheEntry{label, instructionsIndex}
	for {
		if label[2] == 'Z' {
			indexesInCycle = append(indexesInCycle, count-cycle.start)
		}
		label = getNext(label, rune(instructions[instructionsIndex]))
		count++
		instructionsIndex++
		if instructionsIndex == len(instructions) {
			instructionsIndex = 0
		}
		c := CacheEntry{label, instructionsIndex}
		if c == repeatedCacheEntry {
			cycle.zIndexesInCycle = indexesInCycle
			return cycle
		}
	}
}

func getNext(label string, instruction rune) string {
	switch instruction {
	case 'L':
		label = nodeMap[label].left
	case 'R':
		label = nodeMap[label].right
	default:
		panic("wrong direction")
	}

	return label
}

func ExtractNodeB(line string) Node {
	regexRes := reB.FindStringSubmatch(line)
	if len(regexRes) == 0 {
		return Node{text: line} // there will be an empty node in the list, ignore it
	}
	return Node{
		regexRes[1], regexRes[2], regexRes[3], line,
	}
}
