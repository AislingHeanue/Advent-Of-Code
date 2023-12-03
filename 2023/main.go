package main

import (
	"fmt"
	"os"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core/cmd"
)

//go:generate go run ./gen
func main() {
	if err := cmd.CreateCommand().Execute(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
