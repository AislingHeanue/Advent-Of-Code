package day22

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "22b",
		Short: "Day 22, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

func partB(challenge *core.Input) int {
	boxes, matrix := DoSetup(challenge)
	total := 0
	for i := range boxes {
		boxes[i].grounded = false
		// check for any newly-ungrounded boxes
		for z := 1; z < matrix.GetDepth(); z++ {
			for x := 0; x < matrix.GetWidth(); x++ {
				for y := 0; y < matrix.GetHeight(); y++ {
					value := matrix.Get(z, y, x)
					if value == -1 || value == i || !boxes[value].grounded {
						continue
					}
					boxes[value].grounded = checkGrounded(value, boxes)
				}
			}
		}
		for j, newBox := range boxes {
			if j != i && !newBox.grounded {
				total += 1
			}
			boxes[j].grounded = true
		}
	}

	return total
}

func checkGrounded(boxId int, boxes []Box) bool {
	restingMap := boxes[boxId].restingOn
	if len(restingMap) == 0 {
		return true
	}
	grounded := false
	for id := range restingMap {
		if boxes[id].grounded {
			grounded = true
		}
	}
	return grounded
}
