package day15

import (
	"fmt"
	"strings"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "15a",
		Short: "Day 15, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

func partA(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	total := 0
	for _, line := range strings.Split(challenge.LineSlice()[0], ",") {
		total += HASH(line)
	}

	return total
}

func HASH(word string) int {
	current := 0
	for _, letter := range word {
		current = ((current + decode(letter)) * 17) % 256
	}
	return current
}

func decode(letter rune) int {
	// src := []byte{byte(letter)}
	// dst := make([]byte, 1)
	// ascii85.Decode(dst, src, true)
	return int(letter)
}
