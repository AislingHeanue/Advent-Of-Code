package day19

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "19b",
		Short: "Day 19, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

// Oh god, they're really going to make me make a tree rip
// nodes encode max and min x,m,a,s, and the currentState
// leaf <=> currentState = A or R
// if state is A, add (maxX-minX)*(maxM-minM)*... also make sure you're not off by one (thankfully no edge cases this time)
type Node struct {
	mins   map[string]int
	maxes  map[string]int
	states []string
}

func partB(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	makeInstructionMap(challenge)
	root := Node{
		mins:   map[string]int{"x": 1, "m": 1, "a": 1, "s": 1},
		maxes:  map[string]int{"x": 4000, "m": 4000, "a": 4000, "s": 4000},
		states: []string{"in"},
	}
	return getChildren(root)
}

func getChildren(node Node) int {
	for _, letter := range []string{"x", "m", "a", "s"} {
		if node.maxes[letter] < node.mins[letter] {
			return 0 // unreachable state
		}
	}
	if node.states[len(node.states)-1] == "A" {
		return (node.maxes["x"] - node.mins["x"] + 1) * (node.maxes["m"] - node.mins["m"] + 1) * (node.maxes["a"] - node.mins["a"] + 1) * (node.maxes["s"] - node.mins["s"] + 1)
	} else if node.states[len(node.states)-1] == "R" {
		return 0
	}
	count := 0
	currentState := node.states[len(node.states)-1]
	instruction := instructionsMap[currentState]
	for _, rule := range instruction.rules {
		if rule.greater {
			// copy the Node struct and change the mins
			newNode := copyNode(node)
			newNode.states = append(node.states, rule.sendTo)
			newNode.mins[rule.letter] = rule.number + 1 // x > 5 means that minimum must be 6

			count += getChildren(newNode)
			node.maxes[rule.letter] = rule.number // we're in the world where we failed the rule check
		} else {
			// copy the Node struct and change the maxes
			newNode := copyNode(node)
			newNode.states = append(node.states, rule.sendTo)
			newNode.maxes[rule.letter] = rule.number - 1 // x < 5 means that maximum must be 4

			count += getChildren(newNode)
			node.mins[rule.letter] = rule.number
		}
	}
	node.states = append(node.states, instruction.final)
	count += getChildren(node)

	return count
}

func copyNode(node Node) Node {
	newStates := make([]string, len(node.states))
	copy(newStates, node.states)
	newMaxes := map[string]int{"x": node.maxes["x"], "m": node.maxes["m"], "a": node.maxes["a"], "s": node.maxes["s"]}
	newMins := map[string]int{"x": node.mins["x"], "m": node.mins["m"], "a": node.mins["a"], "s": node.mins["s"]}
	return Node{newMins, newMaxes, newStates}
}
