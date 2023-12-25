package day25

import (
	"fmt"
	"strings"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
	"github.com/yourbasic/graph"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "25a",
		Short: "Day 25, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "25b",
		Short: "Day 25, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

func partB(challenge *core.Input) int {
	return 0
}

type Graph struct {
	Nodes map[string]bool
	Edges []EdgeKey
}

type EdgeKey [2]string

func partA(challenge *core.Input) int {
	aislingGraph := Graph{Nodes: make(map[string]bool), Edges: []EdgeKey{}}
	firstNode := ""
	for _, line := range challenge.LineSlice() {
		parts := strings.Split(line, ":")
		from := parts[0]
		if !aislingGraph.Nodes[from] {
			aislingGraph.Nodes[from] = true
			if firstNode == "" {
				firstNode = from
			}
		}
		to := strings.Split(parts[1], " ")
		for _, toNode := range to {
			if len(toNode) == 3 {
				if !aislingGraph.Nodes[toNode] {
					aislingGraph.Nodes[toNode] = true
				}
				aislingGraph.Edges = append(aislingGraph.Edges, [2]string{from, toNode})
			}
		}
	}
	a := getNumNear(aislingGraph, firstNode)

	return a * (len(aislingGraph.Nodes) - a)
}

func getNumNear(aislingGraph Graph, firstNode string) int {
	g := graph.New(len(aislingGraph.Nodes))
	i := 0
	nodeMap := make(map[string]int)
	for node := range aislingGraph.Nodes {
		nodeMap[node] = i
		i++
	}
	for _, edge := range aislingGraph.Edges {
		g.AddBothCost(nodeMap[edge[0]], nodeMap[edge[1]], 1)
	}
	total := 0
	for name, index := range nodeMap {
		if name == firstNode {
			total++
			continue
		}
		if j, _ := graph.MaxFlow(g, nodeMap[firstNode], index); j > 3 {
			total++
		}
	}

	return total
}
