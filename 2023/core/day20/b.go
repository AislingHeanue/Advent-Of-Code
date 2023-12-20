package day20

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "20b",
		Short: "Day 20, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

func partB(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	connections := core.InputMap[Connection](challenge, getConnection)
	connectionMap := getMap(connections)
	highCount := 0
	lowCount := 0
	veryVeryInterestingPointsToSaveTheWorld := []string{}
	for _, p := range connectionMap {
		if p.connectionType == Conjunction &&
			len(p.connectTo) == 1 && p.connectTo[0].connectionType == Conjunction &&
			len(p.connectFrom) == 1 && p.connectFrom[0].connectionType == Conjunction {
			veryVeryInterestingPointsToSaveTheWorld = append(veryVeryInterestingPointsToSaveTheWorld, p.id)
		}
	}

	pressesNeeded := make([]int, len(veryVeryInterestingPointsToSaveTheWorld))
	for i, id := range veryVeryInterestingPointsToSaveTheWorld {
		found := false
		for !found {
			found = pressTheButton(connectionMap, &highCount, &lowCount, id)
			pressesNeeded[i]++
		}
		resetTheButton(connectionMap)
	}
	total := pressesNeeded[0]
	for i := 1; i < len(pressesNeeded); i++ {
		total = util.Lcm(total, pressesNeeded[i])
	}

	return total
}

func resetTheButton(connectionMap map[string]*Connection) {
	for _, p := range connectionMap {
		p.on = false
		p.mostRecentHigh = false
	}
}
