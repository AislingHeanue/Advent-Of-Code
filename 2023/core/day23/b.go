package day23

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "23b",
		Short: "Day 23, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

type Node struct {
	Location util.Point2D
	EdgeKeys map[EdgeKey]bool
}

type EdgeKey struct {
	Node1Loc util.Point2D
	Node2Loc util.Point2D
}

type EdgeValue int

type Graph struct {
	Nodes map[util.Point2D]Node
	Edges map[EdgeKey]EdgeValue
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

// var leavesFoundB int = 0

func partB(challenge *core.Input) int {
	// solve by using my part A implementation
	// it does get the correct answer, it just takes 30 minutes to run
	// so i'll implement a graph instead
	// return solve(challenge, true)
	tiles := challenge.TileMap()
	g := Graph{
		Nodes: make(map[util.Point2D]Node),
		Edges: map[EdgeKey]EdgeValue{},
	}

	//DO A BFS TO GET ALL NODES AND EDGES
	startTile := util.Point2D{}
	endTile := util.Point2D{}
	for x := 0; x < tiles.GetWidth(); x++ {
		if tiles.MustGet(0, x) == '.' {
			startTile = util.Point2D{Y: 0, X: x}
		}
		if tiles.MustGet(tiles.GetHeight()-1, x) == '.' {
			endTile = util.Point2D{Y: tiles.GetHeight() - 1, X: x}
		}
	}
	g.Nodes[startTile] = Node{Location: startTile, EdgeKeys: make(map[EdgeKey]bool)}
	g.Nodes[endTile] = Node{Location: endTile, EdgeKeys: make(map[EdgeKey]bool)}
	findEdges(startTile, endTile, tiles, &g, 0, Down, startTile)
	maxD := 0
	// hooray we finally have the graph, now preform 1.2 million steps of depth first search and hope it's fast
	findBigWalk(g, startTile, endTile, make(map[util.Point2D]bool), 0, &maxD)

	return maxD
}

// this is left here so i remember
// var cache map[CacheEntry]int
// but wait, cacheEntry would need to be a frozen set of checked nodes (and the current node), how the hell would i deal with that?

func findBigWalk(g Graph, currentNode util.Point2D, endNode util.Point2D, checkedNodes map[util.Point2D]bool, currentD int, maxD *int) {
	checkedNodes[currentNode] = true
	defer func() { checkedNodes[currentNode] = false }()
	if currentNode == endNode {
		// leavesFoundB++
		if currentD > *maxD {
			// fmt.Printf("%d found after %d leaves checked\n", currentD, leavesFoundB)
			*maxD = currentD
			return
		}
	}
	for key := range g.Nodes[currentNode].EdgeKeys {
		if key.Node1Loc != currentNode {
			if !checkedNodes[key.Node1Loc] {
				findBigWalk(g, key.Node1Loc, endNode, checkedNodes, currentD+int(g.Edges[key]), maxD)
			}
		} else if key.Node2Loc != currentNode {
			if !checkedNodes[key.Node2Loc] {
				findBigWalk(g, key.Node2Loc, endNode, checkedNodes, currentD+int(g.Edges[key]), maxD)
			}
		} else {
			panic("what?")
		}
	}
}

func findEdges(currentTile util.Point2D, endTile util.Point2D, tiles util.Matrix[rune], g *Graph, stepsSinceNode int, facing Direction, lastNode util.Point2D) {
	allowedDirections := []Direction{}
	allowedNext := []util.Point2D{}
	// special case for the top edge (the start tile only)
	if currentTile.Y == 0 {
		allowedDirections = []Direction{Down}
		allowedNext = []util.Point2D{{Y: currentTile.Y + 1, X: currentTile.X}}
	} else if currentTile.Y != tiles.GetHeight()-1 {
		if point := (util.Point2D{Y: currentTile.Y - 1, X: currentTile.X}); tiles.MustGet(point.Y, point.X) != '#' && facing != Down {
			allowedDirections = append(allowedDirections, Up)
			allowedNext = append(allowedNext, point)
		}
		if point := (util.Point2D{Y: currentTile.Y + 1, X: currentTile.X}); tiles.MustGet(point.Y, point.X) != '#' && facing != Up {
			allowedDirections = append(allowedDirections, Down)
			allowedNext = append(allowedNext, point)
		}
		if point := (util.Point2D{Y: currentTile.Y, X: currentTile.X - 1}); tiles.MustGet(point.Y, point.X) != '#' && facing != Right {
			allowedDirections = append(allowedDirections, Left)
			allowedNext = append(allowedNext, point)
		}
		if point := (util.Point2D{Y: currentTile.Y, X: currentTile.X + 1}); tiles.MustGet(point.Y, point.X) != '#' && facing != Left {
			allowedDirections = append(allowedDirections, Right)
			allowedNext = append(allowedNext, point)
		}
	}
	if len(allowedDirections) == 1 {
		findEdges(allowedNext[0], endTile, tiles, g, stepsSinceNode+1, allowedDirections[0], lastNode)
	} else {
		// we're at a new node, everyone scream
		// if this node already exists, then we don't need to check all the directions around it again
		node, nodeAlreadyExists := g.Nodes[currentTile] // uninitialised if not present, i think that's okay?
		if !nodeAlreadyExists {
			node.Location = currentTile
			node.EdgeKeys = make(map[EdgeKey]bool)
			g.Nodes[currentTile] = node
		}
		// check if this edge already exists in the graph
		edgeVal, ok1 := g.Edges[EdgeKey{currentTile, lastNode}]
		if ok1 {
			g.Edges[EdgeKey{currentTile, lastNode}] = max(edgeVal, EdgeValue(stepsSinceNode))
			return // edge already exists, we're good
		}
		edgeVal, ok2 := g.Edges[EdgeKey{lastNode, currentTile}]
		if ok2 {
			g.Edges[EdgeKey{lastNode, currentTile}] = max(edgeVal, EdgeValue(stepsSinceNode))
			return // edge already exists, we're good
		}

		key := EdgeKey{lastNode, currentTile}

		g.Edges[key] = EdgeValue(stepsSinceNode)

		g.Nodes[currentTile].EdgeKeys[key] = true
		g.Nodes[lastNode].EdgeKeys[key] = true

		// remember recursion? That's still what this function is meant to be
		if !nodeAlreadyExists {
			for i := range allowedDirections {
				findEdges(allowedNext[i], endTile, tiles, g, 1, allowedDirections[i], currentTile)
			}
		}
	}

}
