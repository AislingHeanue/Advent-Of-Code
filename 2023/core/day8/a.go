package day8

import (
	"fmt"
	"regexp"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "8a",
		Short: "Day 8, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

type Node struct {
	label string
	left  string
	right string
	text  string
}

var re = regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)

func partA(challenge *core.Input) int {
	nodes := core.InputMap(challenge, ExtractNode)
	nodeMap := make(map[string]Node)
	instructions := ""
	for _, node := range nodes {
		if node.label != "" {
			nodeMap[node.label] = node
		} else {
			instructions = node.text
		}

	}
	currentLabel := "AAA"
	count := 0
	for currentLabel != "ZZZ" {
		for _, dir := range instructions {
			switch dir {
			case 'L':
				count += 1
				currentLabel = nodeMap[currentLabel].left
			case 'R':
				count += 1
				currentLabel = nodeMap[currentLabel].right
			default:
				panic("wrong direction")
			}
			if currentLabel == "ZZZ" {
				break
			}
		}
	}

	return count
}

func ExtractNode(line string) Node {
	regexRes := re.FindStringSubmatch(line)
	if len(regexRes) == 0 {
		return Node{text: line} // there will be an empty node in the list, ignore it
	}
	return Node{
		regexRes[1], regexRes[2], regexRes[3], line,
	}
}
