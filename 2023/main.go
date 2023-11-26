package main

import (
	"fmt"
	"os"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core/cmd"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
)

//go:generate go run ./gen
func main() {
	util.EbitenSetup()
	if err := cmd.CreateCommand().Execute(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	if util.WindowBeingUsed {
		<-util.WindowClosureChan
	}
}
