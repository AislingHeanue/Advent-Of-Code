// Code generated by 'go run ./gen'; DO NOT EDIT
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day1"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day2"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day3"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day4"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day5"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day6"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day7"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day8"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day9"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day10"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day11"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day12"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day13"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day14"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day15"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day16"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day17"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day18"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day19"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day20"
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day21"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
)
	

func addDays(root *cobra.Command) {
	day1.AddCommandsTo(root)
	day2.AddCommandsTo(root)
	day3.AddCommandsTo(root)
	day4.AddCommandsTo(root)
	day5.AddCommandsTo(root)
	day6.AddCommandsTo(root)
	day7.AddCommandsTo(root)
	day8.AddCommandsTo(root)
	day9.AddCommandsTo(root)
	day10.AddCommandsTo(root)
	day11.AddCommandsTo(root)
	day12.AddCommandsTo(root)
	day13.AddCommandsTo(root)
	day14.AddCommandsTo(root)
	day15.AddCommandsTo(root)
	day16.AddCommandsTo(root)
	day17.AddCommandsTo(root)
	day18.AddCommandsTo(root)
	day19.AddCommandsTo(root)
	day20.AddCommandsTo(root)
	day21.AddCommandsTo(root)
}


func CreateCommand() *cobra.Command {
	var startTime time.Time

	root := &cobra.Command{
		Use:     "2023",
		Short:   "Advent of Code 2023",
		Long:    "Go implementation of my solutions for the 2023 Advent of Code\nGeneration template provided by github.com/nlowe",
		Example: "go run . 1a",
		Args:    cobra.ExactArgs(1),
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			startTime = time.Now()
		},
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Time: %v\n", time.Since(startTime))
			util.AwaitClosure()
		},
	}

	addDays(root)
	root.AddCommand(&cobra.Command{
		Use: "all",
		Short: "Run all days",
		Long: "Gives the result from every day that has been completed",
		Run: func(_ *cobra.Command, _ []string) {
			util.ForceNoWindow = true
			fmt.Printf("Day 1\n")
			day1.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 2\n")
			day2.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 3\n")
			day3.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 4\n")
			day4.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 5\n")
			day5.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 6\n")
			day6.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 7\n")
			day7.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 8\n")
			day8.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 9\n")
			day9.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 10\n")
			day10.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 11\n")
			day11.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 12\n")
			day12.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 13\n")
			day13.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 14\n")
			day14.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 15\n")
			day15.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 16\n")
			day16.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 17\n")
			day17.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 18\n")
			day18.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 19\n")
			day19.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 20\n")
			day20.RunQuestions()
			fmt.Println()
			fmt.Printf("Day 21\n")
			day21.RunQuestions()
			fmt.Println()
		},
	})

	return root
}
