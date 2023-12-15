package day15

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "15b",
		Short: "Day 15, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

type Item struct {
	name  string
	value int
}

var re = regexp.MustCompile(`([a-z]+)([-=])(\d+)?`)

func partB(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	res := re.FindAllStringSubmatch(challenge.LineSlice()[0], -1)
	boxes := make([][]Item, 256)
	currentBoxIndex := 0
	for _, instruction := range res {
		currentBoxIndex = HASH(instruction[1])
		box := boxes[currentBoxIndex]
		if box == nil {
			box = []Item{}
		}
		switch instruction[2] {
		case "-":
			// itemFound := false
			for i, item := range box {
				if item.name == instruction[1] {
					// itemFound = true
					box = append(box[:i], box[i+1:]...)
					boxes[currentBoxIndex] = box
					break
				}
			}

		case "=":
			itemFound := false
			for i, item := range box {
				if item.name == instruction[1] {
					itemFound = true
					item.value, _ = strconv.Atoi(instruction[3])
					box[i] = item
					break
				}
			}
			if !itemFound {
				val, _ := strconv.Atoi(instruction[3])
				box = append(box, Item{name: instruction[1], value: val})
			}
			boxes[currentBoxIndex] = box
		}
	}
	total := 0
	for i, box := range boxes {
		for j, item := range box {
			total += (i + 1) * (j + 1) * item.value
			// fmt.Println((i + 1) * (j + 1) * item.value)
		}
	}
	return total
}
