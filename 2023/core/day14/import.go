package day14

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func AddCommandsTo(root *cobra.Command) {
	root.AddCommand(aCommand())
	root.AddCommand(bCommand())
}

func RunQuestions() {
	fmt.Printf("Part A: %d\n", partA(core.FromFile()))
	fmt.Printf("Part B: %d\n", partB(core.FromFile()))
}
