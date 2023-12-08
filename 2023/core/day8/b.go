package day8

import (
	"fmt"
	"regexp"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
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

	parities := make([]Cycle, len(starterLabels))
	for i, label := range starterLabels {
		parities[i] = getParity(label, instructions)
		parities[i] = getZIndexes(parities[i], instructions)
	}

	// we need to find n such that for all i, n = offset_i + m*cycleLength_i + (one of the z values)
	// eg. 6 = 1 + 2*2 + 1 and 6 = 1 + 0*6 + 5
	// (n - offset)%cycleLength = (one of the z values)

	// find the longest cycle with only one z
	longCycle := parities[0]
	for _, cycle := range parities {
		if len(cycle.zIndexesInCycle) != 1 {
			continue
		}
		if cycle.length > longCycle.length {
			longCycle = cycle
		}
	}

	count := longCycle.start + longCycle.zIndexesInCycle[0]
	for {
		answerFound := true
		for _, cycle := range parities {
			if !isThisNGood(count, cycle) {
				answerFound = false
				break
			}
		}
		if answerFound {
			return count
		}
		count += longCycle.length
	}
}

func isThisNGood(n int, cycle Cycle) bool {
	for _, zIndex := range cycle.zIndexesInCycle {
		if (n-cycle.start)%cycle.length == zIndex {
			return true
		}
	}

	return false
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
func getParity(start string, instructions string) Cycle {
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
